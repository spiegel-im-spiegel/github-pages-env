+++
title = "saracen/walker で歩いてみる"
date =  "2019-10-26T22:25:35+09:00"
description = "saracen/walker を使うと filepath.Walk() 関数を置き換えることができる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "concurrency", "filepath" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

日頃から [mattn/jvgrep] には大変お世話になっているので常時 watch しているのだが，最近組み込まれたらしい [saracen/walker] が良さげである。

指定したパス以下のファイル・ディレクトリを探索する機能としては標準の [`filepath`]`.Walk()` 関数がある。
たとえば，こんな感じで使う。

```go
func WalkWithSingle(rootPath string) (int64, error) {
    count := int64(0)
    lastErr := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            count++
        }
        return nil
    })
    return count, lastErr
}
```

ちなみにこれは，指定したパス以下に存在する（ディレクトリ以外の）ファイルの数を数える関数である。

[saracen/walker] の [`walker`]`.Walk()` 関数を使って [`filepath`]`.Walk()` 関数を置き換えることができる。
[`walker`]`.Walk()` 関数の特徴は，ファイル・ディレクトリの探索を並行処理で行うことである。
したがってマルチコアの環境下（最近のパソコンや携帯端末は皆そうだが）ではコア数に応じた高速処理が期待できる。

[`walker`]`.Walk()` 関数を使う際にはひとつだけ注意点があって，それは [`walker`]`.Walk()` 関数の引数で指定する関数は goroutine-safe でなければならないということだ。

たとえば{{< ruby "closure" >}}関数閉包{{< /ruby >}}を使って，ついうっかり

{{< highlight go "hl_lines=5" >}}
func WalkWithMultiple(rootPath string) (int64, error) {
    count := int64(0)
    err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
        if !info.IsDir() {
            count++
        }
        return nil
    })
    return count, err
}
{{< /highlight >}}

なんてなコード書くと，返ってくるファイル数の値は不定になってしまう（どうしてそうなるかは自分で考えよう）。
まぁ，これに限っては [`sync`]`/`[`atomic`] パッケージを使って

{{< highlight go "hl_lines=5" >}}
func WalkWithMultiple(rootPath string) (int64, error) {
    count := int64(0)
    err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
        if !info.IsDir() {
            atomic.AddInt64(&count, 1)
        }
        return nil
    })
    return count, err
}
{{< /highlight >}}

とすれば{{< ruby "no problem" >}}無問題{{< /ruby >}}である。

[saracen/walker] の性能評価についてはリポジトリのドキュメントを見てもらうとして，ここではもっとふわっとしたコードで性能差を体感してみよう。
用意したコードは上述した関数をちょっと弄ってこんな感じにしてみた。

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "sync/atomic"
    "time"

    "github.com/saracen/walker"
)

func WalkWithSingle(rootPath string) (int64, time.Duration, error) {
    count := int64(0)
    start := time.Now()
    lastErr := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            count++
        }
        return nil
    })
    return count, time.Since(start), lastErr
}

func WalkWithMultiple(rootPath string) (int64, time.Duration, error) {
    count := int64(0)
    start := time.Now()
    err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
        if !info.IsDir() {
            atomic.AddInt64(&count, 1)
        }
        return nil
    })
    return count, time.Since(start), err
}

func main() {
    rootPath := "/usr/local/go/src"
    ct, dt, err := WalkWithSingle(rootPath)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println("WalkWithSingle")
    fmt.Println("\tDuration:", dt)
    fmt.Println("\t   Count:", ct)

    ct, dt, err = WalkWithMultiple(rootPath)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println("WalkWithMultiple")
    fmt.Println("\tDuration:", dt)
    fmt.Println("\t   Count:", ct)
}
```

これを実行するとこんな感じになった。

```text
$ go run sample1.go 
WalkWithSingle
    Duration: 38.305071ms
       Count: 6008
WalkWithMultiple
    Duration: 9.328229ms
       Count: 6008
```

私のマシンは10年前に買った4コアのパソコンなので，まぁ妥当な値だろう。

数を数えるだけでは芸がないのでファイルの一覧を取得してみようか。
たとえば，以下のような `walking.PathList` 型を用意する[^map1]。

[^map1]: 配列ではなく連想配列を使うなら標準の [`sync`]`.Map` 型を使うのがいいだろう。

```go
package walking

import (
    "sync"
)

type PathList struct {
    mutex *sync.Mutex
    list  []string
}

func New() *PathList {
    return &PathList{mutex: &sync.Mutex{}, list: make([]string, 0, 10240)}
}

func (p *PathList) Init() {
    p.mutex.Lock()
    p.list = p.list[:0]
    p.mutex.Unlock()
}

func (p *PathList) Append(path string) {
    p.mutex.Lock()
    p.list = append(p.list, path)
    p.mutex.Unlock()
}

func (p *PathList) List() []string {
    p.mutex.Lock()
    list := make([]string, len(p.list))
    copy(list, p.list)
    p.mutex.Unlock()
    return list
}
```

これを使って先程の `WalkWithMultiple()` 関数を以下のように書き直してみる。

```go
package main

import (
    "fmt"
    "os"
    "sort"

    "walking"

    "github.com/saracen/walker"
)

func WalkWithMultiple(rootPath string) ([]string, error) {
    list := walking.New()
    err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
        if !info.IsDir() {
            list.Append(path)
        }
        return nil
    })
    return list.List(), err
}

func main() {
    rootPath := "/usr/local/go/src"
    list, err := WalkWithMultiple(rootPath)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println("WalkWithMultiple ( Count =", len(list), "):")
    sort.Strings(list)
    for _, path := range list {
        fmt.Println("\t", path)
    }
}
```

これを実行すると以下のような結果になる。

```text
$ go run sample1/sample1.go 
WalkWithMultiple ( Count = 6008 ):
     /usr/local/go/src/Make.dist
     /usr/local/go/src/README.vendor
     /usr/local/go/src/all.bash
     /usr/local/go/src/all.bat
     /usr/local/go/src/all.rc
     /usr/local/go/src/archive/tar/common.go
     /usr/local/go/src/archive/tar/example_test.go
     /usr/local/go/src/archive/tar/format.go
     /usr/local/go/src/archive/tar/reader.go
     /usr/local/go/src/archive/tar/reader_test.go
     /usr/local/go/src/archive/tar/stat_actime1.go
     /usr/local/go/src/archive/tar/stat_actime2.go
     ...
```

よーし，うむうむ，よーし。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[mattn/jvgrep]: https://github.com/mattn/jvgrep "mattn/jvgrep: grep for japanese vimmer"
[saracen/walker]: https://github.com/saracen/walker "saracen/walker: walker is a faster, parallel version, of filepath.Walk"
[`walker`]: https://github.com/saracen/walker "saracen/walker: walker is a faster, parallel version, of filepath.Walk"
[`filepath`]: https://golang.org/pkg/path/filepath/ "filepath - The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"
[`atomic`]: https://golang.org/pkg/sync/atomic/ "atomic - The Go Programming Language"

## ブックマーク

- [Windows でも Grep したい]({{< ref "/remark/2018/03/grep.md" >}})

## 参考図書

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
