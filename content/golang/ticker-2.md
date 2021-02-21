+++
title = "time.Ticker で遊ぶ【Go 1.16 バージョン】"
date =  "2021-02-18T21:34:25+09:00"
description = "素敵なキャンセルライフを（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "time", "context", "goroutine", "concurrency", "signal" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ずいぶん前に「[time.Ticker で遊ぶ][前回]」と言う記事を書いたのだが，先日[リリース]({{< ref "/release/2021/02/go-1_16-is-released.md" >}} "Go 1.16 がリリースされた")された [Go] 1.16 で [`signal`]`.NotifyContext()` 関数が追加された記念に，これを使った改訂版の記事を書いてみたいと思う。

[前回]と同じくお題は以下の通り。

1. 一定周期ごとの処理を行う
2. Ctrl+C 等の割り込み処理を行う

## 一定周期ごとの処理を行う

これは[前回]の記事をほぼそのまま使いまわそう。

```go
// +build run

package main

import (
    "fmt"
    "time"
)

func ticker() {
    t := time.NewTicker(1 * time.Second) //1秒周期の ticker
    defer func() {
        fmt.Println("Stopping ticker...")
        t.Stop()
    }()

    for {
        select {
        case now := <-t.C:
            fmt.Println(now.Format(time.RFC3339))
        }
    }
}

func main() {
    ticker()
}
```

前回でも説明した通り， [defer] 構文を使って終了時に [`time`]`.Ticker.Stop()` 関数で周期イベントを止めようとしているが，実際には無限ループなので return まで到達しない（笑）

## NotifyContext 関数で SIGNAL を捕まえる

[Go] では SIGINT や SIGTERM といった OS から送信される [SIGNAL] をイベントとして [channel] に送り込む仕掛けがある（ちなみに Ctrl+C は SIGINT として送られる）。
さらに [Go] 1.16 では  [SIGNAL] イベントを [`context`]`.Context` のキャンセル・イベントとして実装できるようになった。

たとえば，こんな感じに書ける。

```go {hl_lines=[13, "24-26", "31-34", 37]}
// +build run

package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "time"
)

func ticker(ctx context.Context) {
    t := time.NewTicker(1 * time.Second) //1秒周期の ticker
    defer func() {
        fmt.Println("Stopping ticker...")
        t.Stop()
    }()

    for {
        select {
        case now := <-t.C:
            fmt.Println(now.Format(time.RFC3339))
        case <-ctx.Done():
            fmt.Println("cancellation from context:", ctx.Err())
            return
        }
    }
}

func run() {
    ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
    ticker(ctx)
}

func main() {
    run()
}
```

[`context`] パッケージは並行処理下で使うことが多いだろう。
たとえば `run()` 関数をこんな感じに書き換えてみるか。

```go
func run() {
    ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        n := i + 1
        wg.Add(1)
        go func() {
            defer wg.Done()
            ticker(ctx, n)
        }()
    }
    wg.Wait()
}
```

これで平行に動作している全ての `ticker()` に対してキャンセルを送り込むことができる。

上のコード例ではひとつの [`context`]`.Context` インスタンスを複数の [goroutine] で使いまわしているが，以下のように

```go
func run() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        n := i + 1
        wg.Add(1)
        go func() {
            defer wg.Done()
            ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
            ticker(ctx, n)
        }()
    }
    wg.Wait()
}
```

各 [goroutine] ごとに [`context`]`.Context` インスタンスを生成してセットしても全ての `ticker()` を `Ctrl+C` で問題なく止めることができた。

## キャンセル・イベントの伝搬

[`context`] パッケージは，名前の通り，異なるレイヤやドメイン間でコンテキスト情報を受け渡しするためのパッケージだが，親から子にキャンセルイベントが伝搬する性質がある（逆向きには伝搬しない）。
たとえば

```go
func run() {
    parent, _ := signal.NotifyContext(context.Background(), os.Interrupt)
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        n := i + 1
        wg.Add(1)
        go func() {
            defer wg.Done()
            child, _ := context.WithTimeout(parent, time.Duration(n)*5*time.Second)
            ticker(child, n)
        }()
    }
    wg.Wait()
}
```

などとすれば各  [goroutine] の `ticker()` 関数に [SIGNAL] イベントとタイムアウト・イベントの両方を仕込むことができる。

また

```go
ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
```

とした場合の返り値の `cancel` は関数値になっていて，これをキックすることでペアとなっている [`context`]`.Context` インスタンス（上のコードなら `ctx`）にキャンセル・イベントを発生させることができる。
実際の使い方として [`signal`]`.NotifyContext()` 関数は main [goroutine] に近いところで [`context`]`.WithCancel()` 関数と置き換えることが多いのではないだろうか。

[`context`] について詳しくは『[Go 言語による並行処理](https://www.amazon.co.jp/dp/4873118468?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の 4.12 章が参考になる。
素敵なキャンセル・ライフを（笑）

## ブックマーク

- [Big Sky :: os/signal に NotifyContext が入った。](https://mattn.kaoriya.net/software/lang/go/20200916090416.htm)

[Go]: https://golang.org/ "The Go Programming Language"
[`time`]: http://golang.org/pkg/time/ "time - The Go Programming Language"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`syscall`]: https://golang.org/pkg/syscall/ "syscall - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`signal`]: https://golang.org/pkg/os/signal/ "signal - The Go Programming Language"
[channel]: http://golang.org/ref/spec#Channel_types "The Go Programming Language Specification - The Go Programming Language"
[defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[goroutine]: http://golang.org/ref/spec#Go_statements "The Go Programming Language Specification - The Go Programming Language"
[SIGNAL]: https://linuxjm.osdn.jp/html/LDP_man-pages/man7/signal.7.html "Man page of SIGNAL"
[前回]: {{< relref "./ticker.md" >}} "time.Ticker で遊ぶ"

## 参考図書

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
