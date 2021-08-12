+++
title = "Go による Token Bucket 実装"
date =  "2021-06-03T20:44:50+09:00"
description = "golang.org/x/time/rate パッケージを使う。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "concurrency" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

Qiita で

- [golang.org/x/time/rate でレイトリミット - Qiita](https://qiita.com/msh5/items/f203f85452c5b814ba36)

という記事を見かける。

[`golang.org/x/time/rate`][`rate`] は，いわゆる「トークンバケット（token bucket）」アルゴリズムを実装するためのパッケージのようだ。
トークンバケット・アルゴリズムとは

{{< fig-quote type="markdown" title="Token bucket - Wikipedia" link="https://en.wikipedia.org/wiki/Token_bucket" lang="en" >}}
- A token is added to the bucket every $1/r$ seconds.
- The bucket can hold at the most $b$ tokens. If a token arrives when the bucket is full, it is discarded.
- When a packet (network layer PDU) of $n$ bytes arrives,
  - if at least $n$ tokens are in the bucket, $n$ tokens are removed from the bucket, and the packet is sent to the network.
  - if fewer than $n$ tokens are available, no tokens are removed from the bucket, and the packet is considered to be non-conformant.
{{< /fig-quote >}}

といったものらしい。
具体的には

{{< fig-quote type="markdown" class="nobox" title="rate · pkg.go.dev" link="https://pkg.go.dev/golang.org/x/time/rate" lang="en" >}}
```go
func NewLimiter(r Limit, b int) *Limiter
```
{{< /fig-quote >}}

で生成される [`rate`]`.Limiter` 型のインスタンスが上の説明の「バケット」に相当するようだ。
引数の `r` と `b` も同じ意味かな。

これを並行処理のジェネレータ・パターンと組み合わせると面白そうである。

というわけで，まずは以下のコードを起点としてみよう。

```go:sample1.go
// +build run

package main

import (
    "fmt"
    "sync"
)

func generater(wg *sync.WaitGroup, ch chan<- int) {
    defer func() {
        close(ch)
        wg.Done()
    }()
    for i := 0; i < 10; i++ {
        ch <- i + 1
    }
}

func output(wg *sync.WaitGroup, num int, ch <-chan int) {
    for n := range ch {
        fmt.Printf("Worker %d: %v\n", num, n)
    }
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan int, 1)
    wg.Add(1)
    go generater(&wg, ch)
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go output(&wg, i+1, ch)
    }
    wg.Wait()
}
```

`generater()` 関数はチャネルに対して 1 から 10 までの値を吐き出す。
`output()` 関数はチャネルから値を取り出して表示するという簡単なお仕事である。
ただし `output()` 関数は2つの goroutine で起動している。

これを実行すると

```text
$ go run sample1.go 
Worker 2: 1
Worker 2: 2
Worker 2: 3
Worker 1: 5
Worker 1: 6
Worker 1: 7
Worker 1: 8
Worker 1: 9
Worker 1: 10
Worker 2: 4
```

という出力になった。
2つの goroutine は並行に走ってるので出力順は不定となる。

このコードに [`golang.org/x/time/rate`][`rate`] パッケージを加えて流量の制御を行う。
こんな感じでどうだろう。

```go {hl_lines=[6, 8, 11, 19, "21-24"]}
// +build run

package main

import (
    "context"
    "fmt"
    "sync"
    "time"

    "golang.org/x/time/rate"
)

func generater(wg *sync.WaitGroup, ch chan<- int) {
    defer func() {
        close(ch)
        wg.Done()
    }()
    l := rate.NewLimiter(rate.Every(time.Second*2), 1)
    for i := 0; i < 10; i++ {
        if err := l.Wait(context.Background()); err != nil {
            fmt.Printf("generater: %v\n", err)
            return
        }
        ch <- i + 1
    }
}

func output(wg *sync.WaitGroup, num int, ch <-chan int) {
    for n := range ch {
        fmt.Printf("Worker %d: %v\n", num, n)
    }
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan int, 1)
    wg.Add(1)
    go generater(&wg, ch)
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go output(&wg, i+1, ch)
    }
    wg.Wait()
}
```

これでサイズ 1 のバケットに２秒毎にトークンが挿入される。
トークを取り出すタイミングでチャネルに値を入れるわけだ。
これを実行すると

```text
$ go run sample2.go 
Worker 2: 1
Worker 1: 2
Worker 2: 3
Worker 1: 4
Worker 2: 5
Worker 1: 6
Worker 2: 7
Worker 1: 8
Worker 2: 9
Worker 1: 10
```

という感じに２秒毎に結果が出力される。

バケットサイズをプロセッサ数と同数にしてみよう。
こんな感じかな。

```go {hl_lines=[15, 20, 38, "42-43"]}
// +build run

package main

import (
    "context"
    "fmt"
    "runtime"
    "sync"
    "time"

    "golang.org/x/time/rate"
)

func generater(wg *sync.WaitGroup, ch chan<- int, max int) {
    defer func() {
        close(ch)
        wg.Done()
    }()
    l := rate.NewLimiter(rate.Every(time.Second*2), max)
    for i := 0; i < 10; i++ {
        if err := l.Wait(context.Background()); err != nil {
            fmt.Printf("generater: %v\n", err)
            return
        }
        ch <- i + 1
    }
}

func output(wg *sync.WaitGroup, num int, ch <-chan int) {
    for n := range ch {
        fmt.Printf("Worker %d: %v\n", num, n)
    }
    wg.Done()
}

func main() {
    max := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    ch := make(chan int, max)
    wg.Add(1)
    go generater(&wg, ch, max)
    for i := 0; i < max; i++ {
        wg.Add(1)
        go output(&wg, i+1, ch)
    }
    wg.Wait()
}
```

これを実行すると（プロセッサ数：4）以下のようになった。

```text
$ go run sample3.go 
Worker 3: 3
Worker 3: 4
Worker 2: 2
Worker 4: 1
Worker 1: 5
Worker 3: 6
Worker 2: 7
Worker 4: 8
Worker 1: 9
Worker 3: 10
```

最初の4つは一気に出力されて，以降は2秒ずつの出力。

といった感じで，処理の制限をワーカ側ではなくジェネレータ側で行うのが特徴と言えるだろうか。
Web クローラとか使い道は色々あるかもしれない。

[Go]: https://golang.org/ "The Go Programming Language"
[`rate`]: https://pkg.go.dev/golang.org/x/time/rate "rate · pkg.go.dev"

## 参考図書

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
