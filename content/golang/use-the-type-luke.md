+++
date = "2015-12-13T19:57:29+09:00"
description = "description"
draft = true
tags = ["golang", "programming", "type"]
title = "ルークよ， Type を使え！"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（あー，タイトルに同じネタを2回使ってしまった。
まぁ，いいや。
気を取り直して）

今回のソースコードは “[A Tour of Go](https://tour.golang.org/)” のものをかなり流用しているため取り扱いに注意。
[Go 言語]の公式ドキュメントは CC License の by 3.0，そのうちのソースコードは [BSD license](https://golang.org/LICENSE) で提供されている。

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

[^rn]: rune は Unicode 文字の符号位置（code point）を示す型で文字そのものを表現する。 string と rune の関係については「[String と Rune]({{< relref "golang/string-and-rune.md" >}})」を参照のこと。

さらにこれらの基本型を束ねた構造体 struct を定義できる。

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

ちなみに構造体のフィールド（field）には構造体を含めることもできる。

## 型に名前を付ける

全ての型には [type] キーワードを使って名前を付けることができる。
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

[type] キーワードが使えるのは構造体だけではない。
上述の基本型も [type] キーワードを使って型を再定義できる。

たとえば，2つの時点間の時間を表す [`time`].`Duration` は以下のように定義されている。

```go
type Duration int64
```

また，配列なども型として再定義できる。

```go
type Msgs []string
```

[type] キーワードを使って型に名前を付ける利点は3つある。

1. 名前を付けることでコードの可読性を上げる（オブジェクト指向設計では名前はとても重要）
2. 再利用性の向上（特に構造体の場合）
3. 型に対する関数（メソッド（method）と呼ばれる）を定義できる。

## 型にメソッドを付ける

型に対してメソッドを定義するには以下のように記述する。

```go
func (v *Vertex) String() string {
	return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}
```

`(v *Vertex)` の部分はメソッド・レシーバ（method receiver）と呼ばれ，型とメソッドを関連付ける役割を果たす。
メソッドの呼び出し側は

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v *Vertex) String() string {
	return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}

func main() {
	vertex := &Vertex{X: 1, Y: 2}
	fmt.Println(vertex.String())
}
```

のように記述する[^pr]。
内部処理としては

```go
func String(v *Vertex) string {
	return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}
```

と等価なのだが，型との関連があるかどうかの違いがある。

[^pr]: [`fmt`].`Print` などでは引数のデフォルト書式として `String()` メソッドを持っていることを期待する。したがって `fmt.Println(vertex.String())` と `fmt.Println(vertex)` は同じ結果になる。

ちなみに基本型[^mt] や構造体にはメソッドを付与できない。
たとえば

```go
package main

import "fmt"

func (v *struct{ X, Y int }) String() string {
	return fmt.Sprint("X =", v.X, ", Y =", v.Y)
}

func main() {
	var vertex = &struct {
		X int
		Y int
	}{X: 1, Y: 2}
	fmt.Println(vertex.String())
}
```

などと書いても，コンパイル時に

```
invalid receiver type *struct { X int; Y int } (struct { X int; Y int } is an unnamed type)
```

と怒られる。
[type] キーワードによって型に名前がついていることが重要なのだ。

[^mt]: 基本型だけでなく他パッケージで定義されている型にメソッドを追加することはできない。

[Go 言語]にはクラスがないと言われるが [type] キーワードを使うことで型をクラス・オブジェクトのように記述することができる。

## 汎化と委譲

オブジェクト指向設計においてオブジェクト同士の関係は大きく3つある。

1. 継承（is-a 関係）
2. 包含（has-a 関係）
3. 関連（is-a でも has-a でもない関係）

このうち包含と関連はこれまで説明した方法で実現できるが，継承は表現できない。
そこで以下の機能を使って継承を実現する。

### interface 型

[interface] 型では振る舞いのみを記述する。

[interface] 型で最もよく目にするのは [error] だろう。
[error] は以下のように定義できる[^er1]。

```go
type error interface {
	Error() string
}
```

つまり [error] は「string 型を返す `Error()` メソッド」のみが定義されている。
逆に言うと「string 型を返す `Error()` メソッド」が定義されている全ての型は [error] として扱うことができる[^dt]。

たとえば

```go
package os

import (
	"errors"
)

// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
```

と定義される [`os`].`PathError` は [error] の一種（つまり is-a 関係）であると見なすことができる。

[interface] 型では振る舞いのみで具体的な実装を含まないため，多態性（polymorphism）を持たせた記述が可能になる[^if]。
また， [interface] 型も [type] キーワードで名前を付けることができ，他の型と同じように扱うことができる。

[^er1]: [error] は組み込み型なので，実際にこのような定義が標準パッケージにあるわけではない。 [error] について詳しくは「[エラー・ハンドリングについて]({{< relref "golang/error-handling.md" >}})」を参照のこと。
[^dt]: [Go 言語]では Java の implement のようなものはない。このように振る舞いを記述してオブジェクトを決定する方法を「[ダック・タイピング（duck typing）](https://en.wikipedia.org/wiki/Duck_typing)」と呼ぶ。ダック・タイピングとは「[ダック・テスト（duck test）](https://en.wikipedia.org/wiki/Duck_test)」に由来する言葉だそうで，ダック・テストとは “If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.” と帰納法的に対象を推測する手法を指すらしい。
[^if]: たとえば `interface{}` と記述すればあらゆる型を含むものとなる。これを利用して [`fmt`].`Print` は `func Print(a ...interface{}) (n int, err error) { ... }` と定義されている。

### 型の埋め込み

上の例は継承のうちでも典型的な汎化・特化の振る舞いを記述するものだが，実装上の継承にはもうひとつのタイプがある。
それは機能拡張を行う場合だ。
[Go 言語]では埋め込み（embed）によって機能拡張を行う。

たとえば [`io`].`ReadWriter` は以下のように `Reader` および `Writer` を埋め込んでいる。
これによって `ReadWriter` は `Read()` および `Write()` メソッドを自身のメソッドのように扱うことができる。

```go
package io

import (
	"errors"
)

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

同様に [`bufio`].`ReadWriter` についても

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

と実装されていて， [`bufio`] の `Reader` および `Writer` を埋め込んでいる。
また [`bufio`].`ReadWriter` は [`io`].`ReadWriter` の特化型として機能している点にも注目してほしい。


## ブックマーク

- [オブジェクト指向言語としてGolangをやろうとするとハマること - Qiita](http://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1)
    - [オブジェクト指向言語としてGolangをやろうとするとハマる点を整理してみる - Qiita](http://qiita.com/sona-tar/items/2b4b70694fd680f6297c)
- [Go言語に継承は無いんですか【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/01/15/220136)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[interface]: https://golang.org/ref/spec#Interface_types "Interface types"
[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"
[`time`]: https://golang.org/pkg/time/ "time - The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "time - The Go Programming Language"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[`bufio`]: https://golang.org/pkg/bufio/ "io - The Go Programming Language"
