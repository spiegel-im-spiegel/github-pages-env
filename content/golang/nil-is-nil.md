+++
title = "nil は nil"
date =  "2019-08-18T08:46:56+09:00"
description = "nil は状態を表す「識別子」あるいは「表現」に過ぎず，それ自身は型も値も持たない。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "null", "interface", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Qiita を覗いてたら

- [Goのnilは(nil, nil)という(型, 値)ペアのインターフェースだと把握すれば混乱しない - Qiita](https://qiita.com/momotaro98/items/ee2aea840017266e659d)

という記事を見かけたが，おそらくは `nil` の理解のための方便として意図的に書かれていると思われ，それはそれで悪くないのだが，微妙に危険な香りがするので私なりの解説を記しておく。

## nil は nil

たとえば `fmt.Printf()` 関数などで `nil` の型と値を取ろうとすると

```go
fmt.Printf("Type: %T, Value: %v", nil, nil)
// Output:
// Type: <nil>, Value: <nil>
```

などと表示されるので，いかにも `nil` 型のようなものがあるように見えるが，実際にはこれは「型がない」ことを示している。
同様に値についても，厳密には `nil` という値ではなく「値がない」ことを示しているのだ。

「`nil` とは何か」をきちんと定義した文章は見かけないが， [Go 言語の仕様書](https://golang.org/ref/spec "The Go Programming Language Specification - The Go Programming Language")には，型 `T` の変数 `x` に対して

{{< fig-quote type="markdown" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" >}}
{{< quote >}}`x` is the predeclared identifier `nil` and `T` is a pointer, function, slice, map, channel, or interface type.{{< /quote >}}
{{< /fig-quote >}}

であると記されている。

[Go 言語]では，ある型の値が「宣言されていない」状態のことを「ゼロ値（zero value）」と呼んでいる。
たとえば int などの数値型では「ゼロ値」として数値の `0` を， bool では false を，文字列では空文字列をとる。
同じようにポインタ型や interface 型などでは `nil` を「ゼロ値」としましょう，ということなのである。
このように仕様として定義することで曖昧な状態を排除でき，私達ユーザは安心してその変数を使用することができるわけだ。

したがって `nil` は状態を表す「識別子」あるいは「表現」に過ぎず[^null1]，それ自身は型も値も持たない[^null2]。
強いて言うなら（プログラミング言語で最も悪名高いとされる[^null3]）「null 参照」の一種だとは言えるだろう。

[^null1]: もちろんこれは言語仕様上の話で実装上は何らかの値をとる。大昔のC言語なんかでは「`#define NULL ((void*)0)`」みたいな記述もあったが，さすがにそーゆーのはない（よね？）。
[^null2]: そういう意味では最初に紹介した記事で “Goのnilは(nil, nil)” という部分は間違いではないだろう。
[^null3]: 拙文「[「null 安全」について]({{< ref "/remark/2016/11/null-safety.md" >}})」を参照のこと。

だから本当は

```go
if err != nil {
    ...
}
```

なんかも

```go
if !(err is nil) {
    ...
}
```

みたいな感じに書ければ分かりやすかったのかもしれないが，シンプルを旨とする [Go 言語]でそんな迂遠な表現がとられるわけもなく，敢えて「`nil` との同値性（equality）」という表現をとっているわけだ（偏見）。

## interface 型は，型と値への参照を属性に持つ

まずは，以下の簡単なコードを考えてみる。

```go
package main

import (
	"fmt"
	"strconv"
)

type Binary uint64

func (i Binary) Get() uint64 {
	return uint64(i)
}

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func main() {
	b := Binary(200)
	fmt.Println(b.String())
	// Output:
	// 11001000
}
```

変数 `b` を図式化してみよう[^g1]。
こんな感じ。

[^g1]: 引用元の記事（“[Go Data Structures: Interfaces](https://research.swtch.com/interfaces)”）では 1 word = 32 bits のシステムとして解説されているのでご注意を。

{{< fig-quote type="markdown" title="Go Data Structures: Interfaces" link="https://research.swtch.com/interfaces" lang="en" >}}
{{< fig-img src="https://research.swtch.com/gointer1.png" link="https://research.swtch.com/interfaces" >}}
{{< /fig-quote >}}

これを覚えておいてね。

ここで `main()` 関数の中身を以下のように変えてみる。

{{< highlight go "hl_lines=3-4" >}}
func main() {
	b := Binary(200)
	s := fmt.Stringer(b)
	fmt.Println(s.String())
}
{{< /highlight >}}

ちなみに [`fmt`]`.Stringer` 型は以下に定義される interface 型である。

```go
type Stringer interface {
        String() string
}
```

ゆえに変数 `s` は以下のように図式化できる。

{{< fig-quote type="markdown" title="Go Data Structures: Interfaces" link="https://research.swtch.com/interfaces" lang="en" >}}
{{< fig-img src="https://research.swtch.com/gointer2.png" link="https://research.swtch.com/interfaces" >}}
{{< /fig-quote >}}

このように interface 型は，型と値への参照を属性に持つオブジェクトとして実装されている。

{{< div-box type="markdown" >}}
ただし要素が空の `interface{}` 型では

{{< fig-quote type="markdown" title="Go Data Structures: Interfaces" link="https://research.swtch.com/interfaces" lang="en" >}}
{{< fig-img src="https://research.swtch.com/gointer3.png" link="https://research.swtch.com/interfaces" >}}
{{< /fig-quote >}}

のように最適化されているらしい。
まぁ，ユーザレベルで両者を区別する必要はないけれど。
{{< /div-box >}}

interface 型では `nil` を「ゼロ値」とすると書いたが，そのためには「型と値」の2つの属性とも `nil` でなければならない。
値（への参照）だけが `nil` でも型全体としては `nil` にならないのである。

これでハマりやすいのがエラーハンドリングである。

## エラーハンドリングのハマりどころ

[Go 言語]の組み込み型である `error` は以下のように表すことができる。

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

これを踏まえて，こんなコードを考えてみよう。

```go
package main

import "fmt"

type ErrorObject struct{}

func (m *ErrorObject) Error() string {
	return "I'm error object."
}

func foo() *ErrorObject {
	return nil
}

func bar() error {
	return foo()
}

func main() {
	if err := bar(); err != nil {
		fmt.Printf("%#v is not nil\n", err)
	} else {
		fmt.Printf("%#v is nil\n", err)
	}
}
```

このコードはコンパイルエラーにならないし完全に動くが，実行結果は

```text
(*main.ErrorObject)(nil) is not nil
```

となる。
前節まで読んだならもうお分かりだろうが `bar()` 関数の返り値の `error` は `*ErrorObject` という型を持つため `nil` にならないのである。
したがって `err != nil` は真（`true`）となる。

`bar()` 関数の返り値を正しく評価するには [Conversion](https://golang.org/ref/spec#Conversions) 構文で型を括りだすか，いっそ `foo()` 関数を

```go
func foo() error {
	return nil
}
```

と書き換えるかだろう。
まぁ，後者だよね。
そうすれば実行結果は

```text
<nil> is nil
```

となる。

## ブックマーク

- [エラー・ハンドリングについて（追記あり）]({{< ref "/golang/error-handling.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`fmt`]: https://golang.org/pkg/fmt/ "fmt - The Go Programming Language"

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
