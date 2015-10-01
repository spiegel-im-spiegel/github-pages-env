+++
date = "2015-09-30T00:27:48+09:00"
update = "2015-09-30T20:25:00+09:00"
description = "C++ や Java のような言語圏から来た（私のような）人間にとって Go 言語の「オブジェクト指向」はかなり異質なのだが，慣れてみると逆にとても合理的に見えてくる。この最たるものが error 型である。"
draft = false
tags = ["golang", "error", "exception"]
title = "エラー・ハンドリングについて"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

C++ や Java のような言語圏から来た（私のような）人間にとって [Go 言語]の「オブジェクト指向」はかなり異質なのだが，慣れてみると逆にとても合理的に見えてくる。
なんで C++ や Java はこのようなアプローチをとらなかったのか不思議なほどである。

この最たるものがエラー・ハンドリングだ。

## Go 言語には「例外」がない

「例外（exception）」は本来の処理の流れをぶった切って「大域脱出」するための仕組みである。
でも考えてみれば例外というのはかなり微妙な言語仕様だ。

例外が抱える問題というのは本質的に `goto` 文の問題と同じ[^aa]。
オブジェクトの状態ごと「大域脱出」するため，（脱出前ではなく）脱出後にオブジェクトの後始末を記述せざるを得ないし，記述するためには脱出前の状態（の可能性）を「知識」として知っていなければならない[^a]。
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

[^d]: [error] は組み込み型なので，実際にこのような定義が標準パッケージにあるわけではない。

```go
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

## ブックマーク

- [または私は如何にして例外するのを止めて golang を愛するようになったか — KaoriYa](http://www.kaoriya.net/blog/2014/04/17/)
- [Big Sky :: golang で複数のエラーをハンドリングする方法](http://mattn.kaoriya.net/software/lang/go/20140416212413.htm)
- [DSAS開発者の部屋:Go ではエラーを文字列比較する？という話について](http://dsas.blog.klab.org/archives/go-errors.html)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[Defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[Conversion]: https://golang.org/ref/spec#Conversions "The Go Programming Language Specification - The Go Programming Language"
