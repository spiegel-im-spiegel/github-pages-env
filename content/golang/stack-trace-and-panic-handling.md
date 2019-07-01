+++
date = "2016-02-13T14:48:05+09:00"
update = "2016-11-21T20:48:03+09:00"
description = "panic 時の出力をカスタマイズすることを考える。スタック情報を取得するには， panic を recover で捕まえた上で runtime.Caller() 関数を使う。"
draft = false
tags = ["golang", "stack", "panic"]
title = "スタック追跡とパニック・ハンドリング"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

今回は軽めの小ネタで。

「[エラー・ハンドリングについて]({{< relref "error-handling.md" >}})」でも少し説明したが， [Go 言語]では回復不能のエラー（ゼロ除算やメモリ不足など）が発生した場合には [panic] を投げる仕様になっている。
たとえば以下のコードでは

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    os.Exit(run())
}

func run() int {
    f()
    return 0
}

func f() {
    numbers := []int{0, 1, 2}
    fmt.Println(numbers[3]) //panic!
}
```

以下のスタック情報が標準エラー出力に表示される[^s]。
（[The Go Playground](https://play.golang.org/) での実行結果）

[^s]: ちなみにこの情報は `-s` のリンクオプション（ビルド時に `-ldflags "-s"` と指定する）でデバッグ用のシンボル情報を削除しても表示されるようだ。

```
panic: runtime error: index out of range

goroutine 1 [running]:
main.f()
    /tmp/sandbox269685094/main.go:19 +0x160
main.run(0x20300, 0x104000e0)
    /tmp/sandbox269685094/main.go:13 +0x20
main.main()
    /tmp/sandbox269685094/main.go:9 +0x20
```

まぁ必要な情報はあるのでこれでも構わないのだが，ファイル名がフルパスで表示されるのがアレな感じである。
また出力先が標準エラー出力で固定されているのも面白くない。

そこで [panic] 時の出力をカスタマイズすることを考える。
スタック情報を取得するには， [panic] を [recover] で捕まえた上で [`runtime`].`Caller()` 関数を使う。

```go
package main

import (
    "fmt"
    "io"
    "os"
    "runtime"
)

func main() {
    os.Exit(run(os.Stderr))
}

func run(log io.Writer) (exit int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Fprintf(log, "Panic: %v\n", r)
            for depth := 0; ; depth++ {
                pc, src, line, ok := runtime.Caller(depth)
                if !ok {
                    break
                }
                fmt.Fprintf(log, " -> %d: %s: %s(%d)\n", depth, runtime.FuncForPC(pc).Name(), src, line)
            }
            exit = 1
        }
    }()

    f()
    exit = 0
    return
}

func f() {
    numbers := []int{0, 1, 2}
    fmt.Println(numbers[3]) //panic!
}
```

これで出力は

```
Panic: runtime error: index out of range
 -> 0: main.run.func1: /tmp/sandbox562252505/main.go(19)
 -> 1: runtime.call16: /usr/local/go/src/runtime/asm_amd64p32.s(390)
 -> 2: runtime.gopanic: /usr/local/go/src/runtime/panic.go(423)
 -> 3: runtime.panicindex: /usr/local/go/src/runtime/panic.go(12)
 -> 4: main.f: /tmp/sandbox562252505/main.go(36)
 -> 5: main.run: /tmp/sandbox562252505/main.go(29)
 -> 6: main.main: /tmp/sandbox562252505/main.go(11)
 -> 7: runtime.main: /usr/local/go/src/runtime/proc.go(111)
 -> 8: runtime.goexit: /usr/local/go/src/runtime/asm_amd64p32.s(1133)
```

となる。
ファイル名を出力したくないなら for 文の中を

```go
for depth := 0; ; depth++ {
    pc, _, line, ok := runtime.Caller(depth)
    if !ok {
        break
    }
    fmt.Fprintf(log, " -> %d: %s: (%d)\n", depth, runtime.FuncForPC(pc).Name(), line)
}
```

とする手もある。
コードを書いてる人はスタック追跡情報とファイルの行番号があれば大体あたりをつけられるので，これだけでもありがたい。

## ブックマーク

- [Goでfunctionが実行された順番を追いかける - sgykfjsm.github.com](http://sgykfjsm.github.io/blog/2016/01/20/golang-function-tracing/)
- [Go のバイナリには -ldflags '-w -s' でコンパイルしてもたくさんパスが埋め込まれていた - Qiita](http://qiita.com/kitsuyui/items/d03a9de90330d8c275c8)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[recover]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[`runtime`]: https://golang.org/pkg/runtime/ "runtime - The Go Programming Language"
