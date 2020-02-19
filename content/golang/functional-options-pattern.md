+++
title = "インスタンスの生成と Functional Options パターン"
tags = ["golang", "programming", "functional-options"]
description = "今回も自分用の覚え書きとして書いておく。"
date = "2017-04-04T01:01:59+09:00"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[Go言語のFunctional Option Pattern - Qiita](http://qiita.com/weloan/items/56f1c7792088b5ede136)」を参考にして今回も自分用の覚え書きとして書いておく。


[Go 言語]には C++ や Java 等にある class 宣言がない。
つまりインスタンス（instance）生成時の構築子（constructor）もない。
ではどうやってインスタンスを生成するのか。

たとえば以下のような型を考える。

```go
package ui

import (
    "io"
)

// UI is Command line user interface
type UI struct {
    reader      io.Reader
    writer      io.Writer
    errorWriter io.Writer
}
```

`ui.UI` 型のインスタンスを生成するにはいくつかの方法がある。

まずは `new()` 関数を使う方法。

```go
u := new(ui.UI)
```

`new()` 関数でインスタンスを生成する場合は必ずゼロ値で初期化される[^lc1]。

[^lc1]: `new()` 関数は指定した型のインスタンスを生成し，そのポインタ値を返すが， C++ や Java の new 演算子と違い，生成されたインスタンスの場所がヒープ領域かスタックかについてはコードの文脈依存になっていて隠蔽されている。言い方を変えると，コード記述者はインスタンスの生存期間について気にする必要がない，ということでもある。これは `new()` 関数を使わないインスタンス生成でも同じ。

`ui.UI` の場合は `reader`, `writer`, `errorWriter` の各フィールドには `nil` がセットされる。
しかし [`io`]`.Reader` や [`io`]`.Writer` のような interface 型は `nil` のまま使用すると panic になるため何らかの初期値を与える必要がある。

2番目は複合リテラル（composite literals）で記述する方法。

```go
u := &UI{reader: os.Stdin, writer: os.Stdout, errorWriter: os.Stderr}
```

この方法であれば各フィールドに初期値を与えることができる。
ただしフィールドがパッケージの外からは不可視の場合は（普通そうするよね）この手は使えない。

そこで，3番目の方法として構築子に相当する関数を考える。

```go
package ui

import (
    "bytes"
    "io"
    "io/ioutil"
)

// UI is Command line user interface
type UI struct {
    reader      io.Reader
    writer      io.Writer
    errorWriter io.Writer
}

// New returns a new UI instance
func New(r io.Reader, w, e io.Writer) *UI {
    if r == nil {
        r = ioutil.NopCloser(bytes.NewReader(nil))
    }
    if w == nil {
        w = ioutil.Discard
    }
    if e == nil {
        e = ioutil.Discard
    }
    return &UI{reader: r, writer: w, errorWriter: e}
}
```

こうすれば

```go
u := ui.New(os.Stdin, os.Stdout, os.Stderr)
```

と記述することでパッケージ外でも初期化済みのインスタンスを生成できる。
また

```go
u := ui.New(nil, nil, nil)
```

と無効な値（`nil`）を引数に指定した場合でもフィールドには（`nil` ではなく）安全な値がセットされる。

この方法の問題点は引数に必ず何らかの値をセットしなければならないことだ（[Go 言語]にはデフォルト引数（default argument）のような仕組みはない）。
たとえば `errorWriter` フィールドを使わないことが分かっていてもインスタンス生成時には

```go
u := ui.New(os.Stdin, os.Stdout, nil)
```

などとしなければならない。
また

```go
// NewWithoutErr returns a new UI instance
func NewWithoutErr(r io.Reader, w io.Writer) *UI {
    return New(r, w, nil)
}
```

などと構築子を別途増やす手もあるが，それでは有効なフィールドの組み合わせが増えると関数の管理が煩雑になってしまう。

そこで4番目の方法。
構築子の引数に初期値をセットするのではなく，初期化関数をセットするのである。
この初期化関数の型を

```go
//Option is function value of functional options
type Option func(*UI)
```

と定義する[^srf]。
すると構築子は

[^srf]: これを自己参照関数（self-referential function）と呼ぶそうだ。 “[Self-referential functions and the design of options](https://commandcenter.blogspot.jp/2014/01/self-referential-functions-and-design.html "command center: Self-referential functions and the design of options")” には自己参照関数の様々なバリエーションが紹介されている。この記事ではもっとも簡単な構造のみ紹介している。

```go
//Option is function value of functional options
type Option func(*UI)

// New returns a new UI instance
func New(opts ...Option) *UI {
    u := &UI{reader: ioutil.NopCloser(bytes.NewReader(nil)), writer: ioutil.Discard, errorWriter: ioutil.Discard}
    for _, opt := range opts {
        opt(u)
    }
    return u
}
```

と記述することができる。

さらにフィールドごとに `Option` 関数を返す関数も定義する（これらの関数を用意することで `ui` パッケージを利用するユーザから関数閉包（closure）を隠蔽できる）。

```go
//WithReader returns closure as type Option
func WithReader(r io.Reader) Option {
    return func(u *UI) {
        if r != nil {
            u.reader = r
        }

    }
}

//WithWriter returns closure as type Option
func WithWriter(w io.Writer) Option {
    return func(u *UI) {
        if w != nil {
            u.writer = w
        }
    }
}

//WithErrorWriter returns closure as type Option
func WithErrorWriter(e io.Writer) Option {
    return func(u *UI) {
        if e != nil {
            u.errorWriter = e
        }
    }
}
```

こうしておけばインスタンス生成時の記述は

```go
u := ui.New(ui.WithReader(os.Stdin), ui.WithWriter(os.Stdout))
```

などと初期化の必要なフィールドのみ引数で指定でき，かつコードの見た目も分かりやすくできる。
このようなプログラミング・パターンを “Functional Options" と呼ぶようである。

## ブックマーク

- [command center: Self-referential functions and the design of options](https://commandcenter.blogspot.jp/2014/01/self-referential-functions-and-design.html)
- [Functional options for friendly APIs | Dave Cheney](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

- [Go 言語における「オブジェクト」]({{< relref "object-oriented-programming.md" >}})
- [spiegel-im-spiegel/gocli: Command line interface](https://github.com/spiegel-im-spiegel/gocli) : 本記事と全く同じではないが， Functional Options パターンの実装例を作ってみた

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
