+++
title = "エラー・ハンドリング再考"
date = "2019-02-24T15:57:44+09:00"
description = " golang.org/x/xerrors パッケージの特徴を活かしたエラー・ハンドリングを考えてみる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "error", "programming", "engineering", "design" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]は [golang.org/x/xerrors] パッケージを紹介かたがた簡単なコードをお試しで書いてみたのだが，大雑把に言って [golang.org/x/xerrors] パッケージの特徴は以下のようなものだと言えるだろう。

- エラーの発生位置を取得・表示できる
- エラーの等値性（equality）を検証できる
- エラー・チェイン内の指定した型のインスタンスを抽出できる

であるならエラー・ハンドリングの設計もこれらの特徴を活かしたいものである。

## errors 標準パッケージからの置き換え

たとえば以下のような

```go
func foo() error {
    ...
}

func bar() error {
    ...
}
```

関数 `foo()` と `bar()` から返却される error が等値であれば処理をまとめられる筈である。
今までは [`errors`]`.New()` 関数を使って

```go
var (
    ErrInstance1 = errors.New("error instance 1")
    ErrInstance2 = errors.New("error instance 2")
    ErrInstance3 = errors.New("error instance 3")
)
```

のように，あらかじめ必要な error インスタンスを作成しておき

```go
func errorHandling(err error) {
    switch err {
    case ErrInstance1:
        fmt.Fprintln(os.Stderr, "Error number 1 was triggered.")
        return
    case ErrInstance2:
        fmt.Fprintln(os.Stderr, "Error number 2 was triggered.")
        return
    case ErrInstance3:
        fmt.Fprintln(os.Stderr, "Error number 3 was triggered.")
        return
    default:
        fmt.Fprintln(os.Stderr, "Unknown error")
        return
    }
}
```

みたいな感じにインスタンごとに処理を書いていけばよかった。
しかし [`errors`] を [`xerrors`] に置き換えた場合，このままでは上手くいかないことが分かる。

実際にコードを書いて確かめてみよう。
`errorHandling()` 関数ではエラーの発生位置も出力するようにしてみる。

```go
package main

import (
    "fmt"
    "os"

    "golang.org/x/xerrors"
)

var (
    ErrInstance1 = xerrors.New("error instance 1")
    ErrInstance2 = xerrors.New("error instance 2")
    ErrInstance3 = xerrors.New("error instance 3")
)

func errorHandling(err error) {
    switch err {
    case ErrInstance1:
        fmt.Fprintf(os.Stderr, "Error number 1 was triggered.\n%+v\n", err)
        return
    case ErrInstance2:
        fmt.Fprintf(os.Stderr, "Error number 2 was triggered.\n%+v\n", err)
        return
    case ErrInstance3:
        fmt.Fprintf(os.Stderr, "Error number 3 was triggered.\n%+v\n", err)
        return
    default:
        fmt.Fprintf(os.Stderr, "Unknown error.\n%+v\n", err)
        return
    }
}

func foo() error {
    return ErrInstance1
}
func bar() error {
    return ErrInstance1
}

func main() {
    err1 := foo()
    err2 := bar()
    fmt.Println("foo() error == bar() error ?", err1 == err2)
    errorHandling(err1)
    errorHandling(err2)
}
```

これを実行すると以下のようになる。

```text
$ go run handling1/handling1.go
foo() error == bar() error ? true
Error number 1 was triggered.
error instance 1:
    main.init
        /tmp/xerrors/handling1/handling1.go:11
Error number 1 was triggered.
error instance 1:
    main.init
        /tmp/xerrors/handling1/handling1.go:11
```

これは [`xerrors`]`.New()` 関数を起動した行がエラー発生位置になるためで，これでは本当のエラー発生位置が分からない。
それじゃあ，というので

```go
func foo() error {
    return xerrors.New("error instance 1")
}
func bar() error {
    return xerrors.New("error instance 1")
}
```

と書き換えても駄目である。
[`xerrors`]`.New()` 関数はインスタンスへのポインタ値を返すが `foo()` や `bar()` 関数で返される error と `ErrInstance1` は異なるインスタンスなので [Go 言語]の等値演算子（equality operator）では（ポインタ値を比較するだけで）等値であることを示せない。
したがって実行結果も

```text
$ go run handling1b/handling1b.go
foo() error == bar() error ? false
Unknown error.
error instance 1:
    main.foo
        /tmp/xerrors/handling1b/handling1b.go:34
Unknown error.
error instance 1:
    main.bar
        /tmp/xerrors/handling1b/handling1b.go:37
```

と “`Unknown error`” になる。
もちろん `Error()` 関数で出力される文字列を比較するなんてのは論外である。

## Is 関数をカスタマイズする

上述の問題を解決する方法は色々あると思うが，今回はエラーメッセージのメンテンスのしやすさも考慮して，以下の方法を紹介してみる。

まず基本となる error インスタンスを以下のように定義する。

```go
package werror

type Num int

const (
    ErrInstance1 Num = iota + 1
    ErrInstance2
    ErrInstance3
)

var errMessage = map[Num]string{
    ErrInstance1: "error instance 1",
    ErrInstance2: "error instance 2",
    ErrInstance3: "error instance 3",
}

func (n Num) Error() string {
    if s, ok := errMessage[n]; ok {
        return s
    }
    return "unknown error"
}
```

インスタンスの比較を簡単にするために数値型の `Num` を定義し，値と値に紐づくエラーメッセージを定義する。

次に `Num` 型を埋め込む形で [`xerrors`] 互換の `wrapError` 型を定義する。

```go
type wrapError struct {
    Num
    frame xerrors.Frame
}

func New(n Num) error {
    return &wrapError{Num: n, frame: xerrors.Caller(1)}
}

func (we *wrapError) Format(s fmt.State, v rune) {
    xerrors.FormatError(we, s, v)
}

func (we *wrapError) FormatError(p xerrors.Printer) error {
    p.Print(we.Error())
    we.frame.Format(p)
    return nil
}
```

`Format()` および `FormatError()` 各関数については[前回]を参照のこと。
これで

```go
return werror.New(werror.ErrInstance1)
```

などとすれば定義された error 値を使って `werror.New()` 関数の起動位置で error インスタンスを生成することが出来る。
ただし，このままでは [`xerrors`]`.Is()` 関数によるインスタンスの等値性を正しく検証できない。

そこで `Num` および `wrapError` に以下の関数を追加する。

```go
func (n Num) Is(target error) bool {
	var t1 *wrapError
	if xerrors.As(target, &t1) {
		return n == t1.Num
	}
	var t2 Num
	if xerrors.As(target, &t2) {
		return n == t2
	}
	return false
}

func (we *wrapError) Is(target error) bool {
    var t1 *wrapError
    if xerrors.As(target, &t1) {
        return we.Num == t1.Num
    }
    var t2 Num
    if xerrors.As(target, &t2) {
        return we.Num == t2
    }
    return false
}
```

これで等値演算子 `==` の代わりに `Num.Is()` および `wrapError.Is()` 関数がインスタンスの等値性をチェックしてくれる。

この `werror` パッケージのクラス図はこんな感じかな。

{{< fig-img src="./werror.png" title="werror.png" link="./werror.puml" width="1731" >}}

この `werror` パッケージを使って最初のコードを書き換えてみよう。

```go
package main

import (
    "fmt"
    "os"

    "demo-xerrors/handling2/werror"

    "golang.org/x/xerrors"
)

func errorHandling(err error) {
    switch true {
    case xerrors.Is(err, werror.ErrInstance1):
        fmt.Fprintf(os.Stderr, "Error number 1 was triggered.\n%+v\n", err)
        return
    case xerrors.Is(err, werror.ErrInstance2):
        fmt.Fprintf(os.Stderr, "Error number 2 was triggered.\n%+v\n", err)
        return
    case xerrors.Is(err, werror.ErrInstance3):
        fmt.Fprintf(os.Stderr, "Error number 3 was triggered.\n%+v\n", err)
        return
    default:
        fmt.Fprintf(os.Stderr, "Unknown error.\n%+v\n", err)
        return
    }
}

func foo() error {
    return werror.New(werror.ErrInstance1)
}
func bar() error {
    return werror.New(werror.ErrInstance1)
}

func main() {
    err1 := foo()
    err2 := bar()
    fmt.Println("foo() error == bar() error ?", xerrors.Is(err1, err2))
    errorHandling(err1)
    errorHandling(err2)
}
```

これを実行すると以下のような結果になる。

```text
$ go run handling2/handling2.go
foo() error == bar() error ? true
Error number 1 was triggered.
error instance 1:
    main.foo
        /tmp/xerrors/handling2/handling2.go:30
Error number 1 was triggered.
error instance 1:
    main.bar
        /tmp/xerrors/handling2/handling2.go:33
```

よーし，うむうむ，よーし。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "./xerrors.md" >}} "階層化 Error パッケージ “xerrors” を試してみる"
[golang.org/x/xerrors]: https://godoc.org/golang.org/x/xerrors "xerrors - GoDoc"
[`xerrors`]: https://godoc.org/golang.org/x/xerrors "xerrors - GoDoc"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"

## ブックマーク

- [エラー・ハンドリングについて]({{< relref "./error-handling.md" >}})
- [Error の構造化]({{< relref "./error-handling2.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
