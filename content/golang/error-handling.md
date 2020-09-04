+++
date = "2015-09-30T00:27:48+09:00"
description = "C++ や Java のような言語圏から来た（私のような）人間にとって Go 言語の「オブジェクト指向」はかなり異質なのだが，慣れてみると逆にとても合理的に見えてくる。この最たるものが error 型である。（追記あり）"
tags = ["golang", "error", "exception", "panic", "interface", "defer", "design"]
title = "エラー・ハンドリングについて（追記あり）"

[scripts]
  mathjax = false
  mermaidjs = false
+++

C++ や Java のような言語圏から来た（私のような）人間にとって [Go 言語]の「オブジェクト指向」はかなり異質なのだが，慣れてみると逆にとても合理的に見えてくる。
なんで C++ や Java はこのようなアプローチをとらなかったのか不思議なほどである。

この最たるものがエラー・ハンドリングだ。

## Go 言語には「例外」がない

「例外（exception）」は本来の処理の流れをぶった切って「大域脱出」するための仕組みである。
でも考えてみれば例外というのはかなり微妙な言語仕様だ。

例外が抱える問題というのは本質的に `goto` 文の問題と同じ[^aa]。
オブジェクトの状態ごと脱出するため，（脱出前ではなく）脱出後にオブジェクトの後始末を記述せざるを得ないし，記述するためには脱出前の状態（の可能性）を「知識」として知っていなければならない[^a]。
もし後始末をきちんとしないと，それがバグやリークやその他の脆弱性のもとになる。

[^a]: 例外を備える言語でこれを緩和する仕様はいくつかある。例えば Java は 1.7 から [try-with-resources 構文](http://docs.oracle.com/javase/tutorial/essential/exceptions/tryResourceClose.html)を導入した。もちろんこの構文を有効にするためには対象となるオブジェクトがこの構文に対応した作りになっていなければならない。

[^aa]: ちなみに [Go 言語]の `goto` や ラベル付きの `break`, `continue` は[飛び先に制約](https://golang.org/test/goto.go)があり，どこにでもジャンプできるわけではない。

[Go 言語]はそんな面倒くさいことは考えない。
どうするかというと，普通に返り値に [error] を返す。

```go
file, err := os.Open(filename)
```

[error] は無視することもできる[^0]。

[^0]: いや，ファイル・オープンのエラーを無視したらダメです（笑）

```go
file, _ := os.Open(filename)
```

検出した [error] はその場で処理して抜けてしまえばよい。

```go
file, err := os.Open(filename)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return false
}
```

開始処理と終了（後始末）処理が対になっている場合（`Open`/`Close` とは限らない）は [Defer] 構文で終了処理を保証する[^c]。

[^c]: [Defer] 構文で指定された処理は `return` 時に起動することが保証される。したがって，エラー発生時にはその都度 `return` で抜けて問題ない。むしろ `goto` や `break` で強制的に処理を抜けるよりは処理を分割して `return` で安全に処理を抜ける方法がないか検討すべきである。なお `os.Exit()` などで強制終了した場合は， [Defer] 構文で指定した処理は起動しないので注意。

```go
file, err := os.Open(filename)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return false
}
defer file.Close()
```

これが [Go 言語]の基本的な書き方。
特徴的なのは，ある処理に纏わる処理をセットで記述できる点である。
`try` と `catch` と `finally` の間で目線を行ったり来たりさせなくても，「そこ」だけを見れば把握できる。
ある意味でとても文芸的なコード[^b] であると言える[^bb]。

[^b]: いわゆる「文芸的プログラミング」とは異なるけど。紛らわしくてゴメン。

[^bb]: これからのコードは「文芸的」であることが必要条件だと思う。何故ならエンジニアにとって最も信頼できる「設計書」は（動いている）コードだからだ。コードをひとりで考えてひとりで書いてひとりで使ってひとりでメンテナンスするなら（本人さえ理解していれば）文芸的である必要はないかもしれない。が，実用的なコードでそんな状況はもはやありえない。コードにおいても暗黙知をできるだけ排除していくことが重要。

## error

ここで [error] について改めて説明しておく。
[error] は以下の形式で表現できる [interface] 型のひとつである[^d]。

[^d]: [error] は組み込み型。組み込み型や組み込み関数（`append()` とか）は [`builtin`] パッケージに定義が記述されているが，実際にこのパッケージを import して使うわけではない。

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}
```

つまり `Error()` 関数を持つオブジェクトなら [error] として使える。
これのいちばん簡単な実装が [`errors`] パッケージである。
[`errors`] パッケージの中身は以下のようになっている。

```go
package errors

// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

これは実体として [string] 型の property がひとつだけあって `Error()` 関数で property をそのまま返すというものだ。
[Go 言語]の標準パッケージの多くはこの [`errors`] パッケージを使って [error] を定義している。
たとえば [`os`] パッケージの最下位の [error] は以下のように定義されている。

```go
package os

import (
    "errors"
)

// Portable analogs of some common system call errors.
var (
    ErrInvalid    = errors.New("invalid argument")
    ErrPermission = errors.New("permission denied")
    ErrExist      = errors.New("file already exists")
    ErrNotExist   = errors.New("file does not exist")
)
```

もう少し複雑な [error] では，以下のように詳細情報を持つものもある。

```go
// PathError records an error and the operation and file path that caused it.
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
```

## 3つのエラー・ハンドリング

[error] を判別する方法としては以下の3つがある。

1. [error] インスタンスを比較する
2. [error] の型を判別する
3. `Error()` 関数で出力される文字列を解釈する

### インスタンスを比較する

あらかじめ定義済みの [error] インスタンスがあれば，インスタンスの比較で判別できる。

```go
if err != nil {
    switch err {
    case os.ErrInvalid:
        fmt.Fprintln(os.Stderr, "引数が不正")
    case os.ErrPermission:
        fmt.Fprintln(os.Stderr, "アクセスは許可できない")
    case os.ErrExist:
        fmt.Fprintln(os.Stderr, "そのファイルは既にある")
    case os.ErrNotExist:
        fmt.Fprintln(os.Stderr, "ファイルが存在しない")
    default:
        fmt.Fprintln(os.Stderr, err)
    }
    return
}
```

単純にエラーの種類が分かればいいのであれば，この方法が最もシンプル。

### 型を判別する

[error] は [interface] 型なので [Conversion] 構文で判別できる。

```go
if err != nil {
    switch e := err.(type) {
    case *os.PathError:
        if errno, ok := e.Err.(syscall.Errno); ok {
            switch errno {
            case syscall.ENOENT:
                fmt.Fprintln(os.Stderr, "ファイルが存在しない")
            case syscall.ENOTDIR:
                fmt.Fprintln(os.Stderr, "ディレクトリが存在しない")
            default:
                fmt.Fprintln(os.Stderr, "Errno =", errno)
            }
        } else {
            fmt.Fprintln(os.Stderr, "その他の PathError")
        }
    default:
        fmt.Fprintln(os.Stderr, "その他のエラー")
    }
    return
}
```

[error] に状態（status）を持たせる必要がある場合は，この方法を使うべき。

### 文字列を比較する

上述の方法で判別できない場合は `Error()` 関数で出力される文字列を解釈して処理するしかない。
[バッドノウハウ](http://0xcc.net/misc/bad-knowhow.html)。

## エラー・ハンドリングの設計

エラー・ハンドリングの方針としては，以下の2つのうちのどちらかだろう。

1. [error] を下位ロジックから上位ロジックまで持ち回し，最上位ロジックで最終的な判定と処理を行う
2. 下位ロジックの [error] をカプセル化した新たなインスタンスを生成し上位ロジックに渡す。上位ロジックは直近のロジックの [error] のみが見える

最初のやり方は一見よさげだが，この方針では上位ロジックが下位ロジックの全ての [error] を把握している必要があり現実的でない。
またオブジェクト指向設計では “Don't talk to strangers” の原則があり，いわゆる「友達の友達」のことは知らないふりをするのがよい設計と言われている。

こう考えると文字列での比較は最も下策であると言える。
また，型を判別する場合でも，下位レイヤの状態を生のまま見せるのではなく，必要な情報のみを返す関数を実装するほうが上策と言えるだろう。

もうひとつ考慮すべき点としてエラー・メッセージの設計が挙げられるだろう。
[error] に対するメッセージをどのように設計するかは（大規模アプリケーションでは特に）重要である。

（「[Error の構造化]({{< relref "error-handling2.md" >}})」および「[エラー・ハンドリング再考]({{< relref "./error-handling-again.md" >}})」へ続く）

## 【追記】 Panic と Recover

たとえばゼロ除算を行った場合や配列などで領域外を参照・設定しようとした場合，あるいは allocation に失敗した場合など，致命的なエラーが発生する場合がある。

```go
package main

import "fmt"

func main() {
    f()
}

func f() {
    numbers := []int{0, 1, 2}

    fmt.Println(numbers[3])
}
```

これを実行すると

```bash
C:\workspace\go-practice\src\panic01>go run panic01.go
panic: runtime error: index out of range

goroutine 1 [running]:
main.f()
        C:/workspace/go-practice/src/panic01/panic01.go:12 +0x14a
main.main()
        C:/workspace/go-practice/src/panic01/panic01.go:6 +0x1b
exit status 2
```

となり，大域脱出させてアプリケーションを強制終了させているのが分かる。
この仕組みを [Panic] と呼ぶ。

[Panic] は意図的に発生させることもできる。

```go
package main

func main() {
    f()
}

func f() {
    panic("Panic!")
}
```

これを実行すると

```bash
C:\workspace\go-practice\src\panic02>go run panic02.go
panic: Panic!

goroutine 1 [running]:
main.f()
        C:/workspace/go-practice/src/panic02/panic02.go:8 +0x6c
main.main()
        C:/workspace/go-practice/src/panic02/panic02.go:4 +0x1b
exit status 2
```

となる。

一方で， [Panic] を [Recover] することもできる[^e]。

[^e]: [Recover] は [Defer] 構文とともに使用する。つまり [Panic] 発生時でも [Defer] 構文で予約された処理は実行される。

```go
package main

import "fmt"

func main() {
    err := r()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Normal End.")
    }
}

func r() (err error) {
    defer func() {
        if rec := recover(); rec != nil {
            err = fmt.Errorf("Recovered from: %v", rec)
        }
    }()

    f()
    err = nil
    return
}

func f() {
    panic("Panic!")
}
```

これを実行すると

```bash
C:\workspace\go-practice\src\panic03>go run panic03.go
Recovered from: Panic!
```

となる。
[Panic] を [Recover] で捕まえて通常の [error] として返しているのがお分かりだろうか。

一般的に [Panic] はアプリケーション内で続行不可能な致命的エラーが発生した場合に投げられる。
例えばゼロ除算や領域外アクセスのようなエラーは [Panic] が発生する前に回避するコードにすべきだ。
Allocation エラーのような回避不能かつアプリケーション続行不可能なエラーの場合は [Panic] が投げられるのもやむを得ないが， [Recover] することにほとんど意味はない。

例外的な使い方として [`bytes`]`.Buffer` では，メモリ確保で panic が発生した際に [Recover] で捕まえ， [error] インスタンスを入れ替えて [Panic] を投げ直している。
このような用途で [Recover] を使うことはあり得る。

```go
// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []byte {
    // If the make fails, give a known error.
    defer func() {
        if recover() != nil {
            panic(ErrTooLarge)
        }
    }()
    return make([]byte, n)
}
```

また再帰処理中に続行不能なエラーが発生した場合に [Panic] を投げてトップレベルの関数に一気に復帰するような使い方をすることもある。
この場合，トップレベルの関数は（続行不可能なら）改めて [Panic] を投げるか（処理続行できる根拠があるのなら）通常の [error] を返すことになる[^f]。

[^f]: これ以外にサーバ用途などでプロセスを落とせない場合に [Recover] で回避することもあるが，既に続行不可能な状態で無理やりプロセスを続行するのが正しい動きなのかどうかは疑問が残る。

いずれにしろ，いわゆる「例外処理」的なハンドリングを [Panic]/[Recover] で行うべきではない。

## 【追記】 error の nil が nil にならない場合

- [Why Go is a poorly designed language — Medium](https://medium.com/@tucnak/why-go-is-a-poorly-designed-language-1cc04e5daf2#.ucutrogyz) （[日本語訳](http://postd.cc/why-go-is-a-poorly-designed-language/)）

{{< fig-gist "https://gist.github.com/tucnak/eccdb53e7884084f5674" >}}

このコードの実行結果は “Hello, Mr. Pike!” を出力する。
このコードのポイントは `Generate()` 関数が [error] ではなく `*MagicError` 型を返している点にある。

実は [error] を含む [interface] 型のインスタンスは値（への参照）と型情報をセットで保持っているため，上述のような形で nil を返しても受け取った側は「nil 状態を持つなにか」という評価になり，ただの `<nil>` にはならないのだ。

- [research!rsc: Go Data Structures: Interfaces](https://research.swtch.com/interfaces)

たとえば上のコードを以下のように書き換えると分かりやすいかもしれない。

```go
package main

import "fmt"

type MagicError struct{}

func (m *MagicError) Error() string {
    return fmt.Sprintf("%#v", m)
}

func Generate() *MagicError {
    return nil
}

func Test() error {
    return Generate()
}

func main() {
    if err := Test(); err != nil {
        fmt.Println(err)
    }
}
```

このコードを実行すると “`(*main.MagicError)(nil)`” と出力する。
`Generate()` 関数が返す nil がどのように機能しているか分かると思う。
「[Go 言語における「オブジェクト」]({{< relref "object-oriented-programming.md" >}})」で解説するが， [Go 言語]の型（[type]）は C++ や Java で言うところの class のように機能するため，このような動きになると思われる[^nil]。

[^nil]: それでもやっぱり nil は nil として扱ってほしいのだが。

エラーハンドリングを行う際は結構ここがハマりどころになる。
ご注意を。

## ブックマーク

- [または私は如何にして例外するのを止めて golang を愛するようになったか — KaoriYa](http://www.kaoriya.net/blog/2014/04/17/)
- [Big Sky :: golang で複数のエラーをハンドリングする方法](http://mattn.kaoriya.net/software/lang/go/20140416212413.htm)
- [DSAS開発者の部屋:Go ではエラーを文字列比較する？という話について](http://dsas.blog.klab.org/archives/go-errors.html)
- [panicはともかくrecoverに使いどころはほとんどない - Qiita](http://qiita.com/ruiu/items/ff98ded599d97cf6646e)
- [Golang: nil Pointer Receiverの話 - Qiita](http://qiita.com/stsn/items/73714caf8458b1d973f2)
- [echoのAPIサーバ実装とエラーハンドリングの落とし穴 - Qiita](http://qiita.com/tienlen/items/5f2bcfe06eb83830ee55)
- [Golangのエラー処理とpkg/errors | SOTA](http://deeeet.com/writing/2016/04/25/go-pkg-errors/) : 必見！
- [Go: Multiple Errors Management. Error management in Go is always prone… | by Vincent Blanchon | A Journey With Go | Sep, 2020 | Medium](https://medium.com/a-journey-with-go/go-multiple-errors-management-a67477628cf1)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"
[`builtin`]: https://golang.org/pkg/builtin/ "builtin - The Go Programming Language"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[Defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[Panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[Recover]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[Conversion]: https://golang.org/ref/spec#Conversions "The Go Programming Language Specification - The Go Programming Language"
[`bytes`]: https://golang.org/pkg/bytes/ "bytes - The Go Programming Language"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
