+++
date = "2015-12-15T21:19:58+09:00"
update = "2018-11-11T15:28:29+09:00"
description = "Go 言語がいわゆる「オブジェクト指向言語」と言えるかどうかについては色々とあるようだが，オブジェクト指向プログラミングを助けるための仕掛けはいくつか存在する。今回はその中の type キーワードを中心に解説していく。"
draft = false
tags = ["golang", "object-oriented", "programming", "type", "interface"]
title = "Go 言語における「オブジェクト」"

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
  url = "http://www.baldanders.info/spiegel/profile/"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]がいわゆる「オブジェクト指向言語」と言えるかどうかについては色々とあるようだが，オブジェクト指向プログラミングを助けるための仕掛けはいくつか存在する。
今回はその中の [type] キーワードを中心に解説していく。

なお，今回のソースコードは “[A Tour of Go](https://tour.golang.org/)” のものをかなり流用しているため取り扱いに注意。
[Go 言語]の公式ドキュメントは CC License の by 3.0，ソースコードは [BSD license](https://golang.org/LICENSE) で提供されている。

## Go 言語の基本型

今さらだけど， [Go 言語]の基本型（basic type）は以下の通り。

- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- uintptr
- byte
- rune
- float32, float64
- complex64, complex128

このうち byte は uint8 の別名で rune[^rn] は int32 の別名である。
また int, uint, uintptr のサイズはプラットフォーム依存になっている。
string は不変（immutable）な値で，その実体は byte 配列である。
基本型は組み込み型であり，振る舞いを追加・変更することはできない。

[^rn]: rune は Unicode 文字の符号点（code point）を示す型で文字そのものを表現する。 string と rune の関係については「[String と Rune]({{< relref "string-and-rune.md" >}})」を参照のこと。

さらにこれらの基本型を集約した構造体 [struct] を定義できる。

```go
package main

import "fmt"

func main() {
    vertex := struct {
        X int
        Y int
    }{X: 1, Y: 2}
    fmt.Println(vertex)
}
```

ちなみに構造体のフィールド（field）には構造体を含めることができ，入れ子構造にすることもできる。

この他に配列（slice[^slc] を含む）や連想配列（map）あるいは関数値（function value）といったものもあるが，今回は踏み込まない。

[^slc]: 配列と slice の関係については「[配列と Slice]({{< relref "array-and-slice.md" >}})」で紹介している。

## 型に名前を付ける

全ての型には [type] キーワードを使った宣言（type declaration）により名前を付けることができる。
例えば先ほどのコードは

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    vertex := Vertex{X: 1, Y: 2}
    fmt.Println(vertex)
}
```

と書き直すことができる。

[type] 宣言が使えるのは構造体だけではない。
上述の基本型も [type] 宣言を使って型を再定義できる。

たとえば，2つの時点間の時間を表す [`time`]`.Duration` は以下のように定義されている。

```go
type Duration int64
```

また，配列なども型として再定義できる。

```go
type Msgs []string
```

[type] 宣言を使って型に名前を付ける利点は3つある。

1. 名前を付けることでコードの可読性を上げる（オブジェクト指向設計では名前がとても重要）
2. 再利用性の向上（特に構造体の場合）
3. 型に関数を関連付けることができる。

[type] 宣言による名付けは単なる別名定義ではないということだ[^ta1]。

[^ta1]: 型の別名定義はバージョン 1.9 から導入された。詳しくは「[Go 1.9 と Type Alias]({{< relref "go-1_9-and-type-alias.md" >}})」を参照のこと

## 型に関数を関連付ける

型に関数を関連付けるには以下のように記述する。

```go
func (v Vertex) String() string {
    return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}
```

`(v Vertex)` の部分はメソッド・レシーバ（method receiver）と呼ばれ，これが型と関数を関連付ける役割を果たす。
内部処理としては

```go
func Vertex.String(v Vertex) string {
    return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}
```

と等価である[^call]。
関数の呼び出し側は

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func (v Vertex) String() string {
    return fmt.Sprint("X = ", v.X, ", Y = ", v.Y)
}

func main() {
    vertex := Vertex{X: 1, Y: 2}
    fmt.Println(vertex.String())
}
```

のようにインスタンスと関数をピリオドで連結して記述する[^pr]。

[^call]: [Go 言語]の関数呼び出しにおいて引数の渡し方は基本的に「値渡し」である。詳しくは「[関数とポインタ]({{< relref "function-and-pointer.md" >}})」を参照のこと。
[^pr]: ちなみに [`fmt`]`.Print` などでは引数の型が `String()` を持っていることを期待し，この関数の出力結果をデフォルト書式にしている。したがって `fmt.Println(vertex)` と書いても同じ結果になる。

構造体そのものには関数を付与できない[^mt]。
たとえば

[^mt]: 他にも基本型や他パッケージで定義されている型に関数を追加することはできない。当たり前だけど。

```go
package main

import "fmt"

func (v struct{ X, Y int }) String() string {
    return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}

func main() {
    var vertex = struct {
        X int
        Y int
    }{X: 1, Y: 2}
    fmt.Println(vertex.String())
}
```

などと書いても，コンパイル時に

```
invalid receiver type struct { X int; Y int } (struct { X int; Y int } is an unnamed type)
```

と怒られる。
[type] 宣言によって型に名前が付けられていることが重要なのだ。

[Go 言語]には class キーワードはないが， [type] 宣言を使うことで，名前と属性と操作を持つクラスを記述することができる[^cls]。

[^cls]: クラスは名前と属性と操作の3つの要素で構成されている。名前は他クラスと識別できるものを1個。属性と操作は0個以上存在する。 [Go 言語]では空のフィールドの [struct] を定義することにより0個の属性を持つクラスを構成できる。

## 汎化・特化と処理の委譲

オブジェクト指向設計においてクラス間の関係は大きく2つある。

1. 汎化・特化[^c1]（継承または is-a 関係）
2. 関連[^c2]（包含または has-a 関係）

[^c1]: 言わずもがなだが，サブクラスから見たスーパークラスが「汎化」でその逆が「特化」である。
[^c2]: 関連は更に集約と複合に分類できるが今回は踏み込まない。

このうち関連についてはこれまで説明した方法で実現できるが，汎化・特化は表現できない。
そこで以下の機能を使って汎化・特化を実現する。

### 振る舞いのみを定義した型

[interface] を使うと振る舞いのみを定義した汎化クラスを表現することができる。

[interface] で定義された型で最もよく目にするのは [error] だろう。
[error] は以下のように定義できる[^er1]。

[^er1]: [error] は組み込み型なので，実際にこのような定義が標準パッケージにあるわけではない。 [error] について詳しくは「[エラー・ハンドリングについて]({{< relref "error-handling.md" >}})」を参照のこと。

```go
type error interface {
    Error() string
}
```

定義が示すように [error] は「string 型を返す `Error()` 関数」のみが定義されている。
逆に言うと「string 型を返す `Error()` 関数」を持つ全ての型は [error] の一種（つまり is-a 関係）であると見なすことができる[^dt]。

[^dt]: [Go 言語]では Java の implement のような interface クラスからの継承を明示するキーワードはない。記述された振る舞いからクラス関係を決定する方法を「構造的部分型（structural subtyping）」と呼ぶ。構造的部分型のメリットのひとつは多重継承で発生する様々な問題（名前の衝突や菱形継承など）を気にする必要がない点である。

たとえば

```go
package os

// PathError records an error and the operation and file path that caused it.
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
```

と定義される [`os`]`.PathError` は [error] の一種である。

[interface] も [type] 宣言で名前を付けることができ，他の型と同じように扱うことができる。
さらに [interface] で定義した型は振る舞いのみで具体的な実装を含まないため，多態性を持たせた記述が可能になる[^if]。

[^if]: たとえば `interface{}` と記述すればあらゆる型を含むことになる。これを利用して [`fmt`]`.Print` は `func Print(a ...interface{}) (n int, err error) { ... }` と定義されている。

### 将来バージョンで総称型（generics）がサポートされるかも。

現行版の [Go 言語]ではいわゆる[「総称型」はサポートされていない](https://golang.org/doc/faq#generics)が，将来バージョンで導入される可能性がある。

- [次期 Go 言語で導入される（かもしれない）総称型について予習する]({{< relref "generics-in-go-2.md" >}})

### 型の埋め込み

もうひとつの汎化・特化の機能が型の埋め込み（embedding）である。
構造体や [interface] には別の型を埋め込むことができる。

たとえば [`io`]`.ReadWriter` は以下のように `Reader` および `Writer` を埋め込んでいる。
（このときの `Reader` および `Writer` を「埋め込みインタフェース（embedding interface）」と呼ぶ）

```go
package io

// Implementations must not retain p.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Implementations must not retain p.
type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter interface {
    Reader
    Writer
}
```

これによって `ReadWriter` は `Read()` および `Write()` を自身の振る舞いのように扱うことができる。
この場合も `ReadWriter` は `Reader` および `Writer` の一種であると見なすことができる。

同様に [`bufio`]`.ReadWriter` についても

```go
package bufio

// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader
    *Writer
}

// NewReadWriter allocates a new ReadWriter that dispatches to r and w.
func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
    return &ReadWriter{r, w}
}
```

と実装されていて， [`bufio`] の `Reader` および `Writer` を埋め込み，これらの型の一種として実装されている。
このときの `Reader` および `Writer` を「埋め込みフィールド（embedded field）」または「匿名フィールド（anonymous field）」と呼ぶ。
[`bufio`]`.ReadWriter` は [`io`]`.ReadWriter` の一種として機能している点にも注目してほしい。

### 関数のオーバーライドと処理の委譲

では，今まで述べたことを使って以下のコードを書いてみる。

```go
package main

import "fmt"

type ErrorInfo interface {
    error
    Errno() int
}

type ErrorInfo1 struct{}

func (err *ErrorInfo1) Error() string {
    return fmt.Sprint("Error Information: ", err.Errno())
}

func (err *ErrorInfo1) Errno() int {
    return 1
}

func Action() error {
    err := &ErrorInfo1{}
    return err
}

func main() {
    if err := Action(); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Normal End")
}
```

[error] の拡張として `ErrorInfo` [interface] を定義する。
`ErrorInfo` [interface] では [error] を埋め込み，さらに `Errno()` 関数を追加定義している。
更に `ErrorInfo` [interface] に適合するクラスとして `ErrorInfo1` 型を作った。
このコードの実行すると “`Error Information: 1`” が出力される。

{{< fig-img src="./interface.svg" title="ErrorInfo interface" link="./interface.svg" >}}

次に `ErrorInfo1` のバリエーションとして `ErrorInfo2` を追加してみよう。

```go
package main

import "fmt"

type ErrorInfo interface {
    error
    Errno() int
}

type ErrorInfo1 struct{}

func (err *ErrorInfo1) Error() string {
    return fmt.Sprint("Error Information: ", err.Errno())
}

func (err *ErrorInfo1) Errno() int {
    return 1
}

type ErrorInfo2 struct {
    ErrorInfo1
}

func (err *ErrorInfo2) Errno() int {
    return 2
}

func Action() error {
    err := &ErrorInfo2{}
    return err
}

func main() {
    if err := Action(); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Normal End")
}
```

`ErrorInfo2` では `Error()` 関数は `ErrorInfo1` のものをそのまま使い回したいが `Errno()` 関数では異なる値を出力したい，と考えた。
そこで `Errno()` 関数をオーバライドしようと考えた。
図で書くと以下のような構成を期待したわけだ。

{{< fig-img src="./override.svg" title="Override ?" link="./override.svg" width="518" >}}

しかし実際には，期待した “`Error Information: 2`”  ではなく，前回と同じ “`Error Information: 1`” が出力される。

上のコードでは `ErrorInfo2` と直接関連付けられた `Error()` がないため `ErrorInfo1` の `Error()` が呼ばれるが，その関数の中で呼ばれる `Errno()` は `ErrorInfo2` と関連付けられた関数ではなく `ErrorInfo1` と関連付けられた関数になる。

{{< fig-img src="./delegation.svg" title="delegation" link="./delegation.svg" width="518" >}}

これは [Go 言語]では埋め込みフィールドの関数呼び出しが「委譲」として機能しているためである。
たとえば C++ 言語では virtual 修飾子を付与して仮想関数化することで意図的にオーバーライドできるが[^dlg]， [Go 言語]ではこのような仕掛けがないため[^ef]，呼ばれた関数は常に委譲として機能する[^ovr]。

[^dlg]: 逆に Java では関数は常に仮想関数として機能しオーバーライドされる可能性がある。これを抑止するためには final 修飾子を付加する。
[^ef]: [Go 言語]的には埋め込みフィールドはフィールドのバリエーションのひとつにすぎないため，動作も通常のフィールドが持つ関数を呼び出した場合と変わらない。そういう意味では構造体への埋め込みは，見かけ上は「is-a 関係」でも，実質的には「has-a 関係」に近いと言えるかもしれない。
[^ovr]: 複数の型を埋め込んでいる場合，埋め込みフィールド間で名前が衝突しているフィールドや関数を使おうとするとコンパイルエラーになる。この場合は `err.ErrorInfo1.Error()` のように型を明示して回避できる。

上の例はクラス構成からして明らかにダメダメなのだが，今回のポイントはサブクラスである `ErrorInfo2` から `Errno()` 関数を上書きすることでスーパークラス `ErrorInfo1` の `Error()` 関数の処理を書き換えようとした点にある。
継承[^ih] の実装で一番よくあるミスがこの「カプセル化の破れ」で， [Go 言語]の場合は敢えて移譲を強制することでこの手の不具合が発生するのを回避しようとしているように見える。

他の言語では明示的に委譲を実装しようとすると冗長な記述になることが多いが， [Go 言語]の場合は埋め込みを使うことでシンプルな記述で委譲を実装できる点がメリットと言える。

[^ih]: ここで言う継承は設計時の「汎化・特化」のことではなく，言語機能などを使った実装上の継承のこと。

## ブックマーク

- [Big Sky :: Go言語でインタフェースの変更がそれ程問題にならない理由](http://mattn.kaoriya.net/software/lang/go/20130919023425.htm)
- [オブジェクト指向言語としてGolangをやろうとするとハマること - Qiita](http://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1)
    - [オブジェクト指向言語としてGolangをやろうとするとハマる点を整理してみる - Qiita](http://qiita.com/sona-tar/items/2b4b70694fd680f6297c)
- [Go言語に継承は無いんですか【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/01/15/220136)
- [Go言語でジェネリクスっぽいことがしたいでござる【generics】【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/06/16/224109)
- [Go 言語の値レシーバとポインタレシーバ | Step by Step](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)
- [埋込みとインタフェース #golang - Qiita](http://qiita.com/tenntenn/items/bba69d84a1e0b67c4db8)
- [Goにatexitやグローバルなデストラクタがない理由 - Qiita](http://qiita.com/ruiu/items/22910a4bae6cb716a391)
- [Python と Ruby と typing - methaneのブログ](http://methane.hatenablog.jp/entry/ruby-python-go-typing) : [interface] を使った汎化・特化の関係は duck typing ではなく「構造的部分型（structural subtyping）」と言うのが正しいらしい

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[struct]: https://golang.org/ref/spec#Struct_types "Struct types"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[interface]: https://golang.org/ref/spec#Interface_types "Interface types"
[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"
[`time`]: https://golang.org/pkg/time/ "time - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "time - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`bufio`]: https://golang.org/pkg/bufio/ "io - The Go Programming Language"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" height="160" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"></a></div>
    <dl class="fn">
      <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
      <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
      <dd>丸善出版</dd>
      <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
      </span></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.19</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
