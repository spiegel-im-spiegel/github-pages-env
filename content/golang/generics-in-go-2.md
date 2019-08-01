+++
title = "次期 Go 言語で導入される（かもしれない）総称型について予習する"
date = "2018-11-11T06:10:48+09:00"
description = "現時点ではまだドラフト案なのでフワっとした説明になるけど，あしからずご了承の程を。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "language", "programming", "generics", "type", "contract" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

2018年8月に次期 [Go 言語]で追加される（かもしれない）仕様についてアナウンスがあった。

- [Go 2 Draft Designs - The Go Blog](https://blog.golang.org/go2draft)
- [Go 2 Draft Designs](https://go.googlesource.com/proposal/+/master/design/go2draft.md)

{{< fig-youtube id="6wIP3rO6On8" title="Go 2 Drafts Announcement - YouTube" >}}

“Go 2” といってもメジャー・バージョンが変わるのではなく現行バージョンに対する追加仕様らしい。
したがって後方互換性は確保されているようだ。

紹介されているドラフト案は大きく2つある。

- 総称型（generics）
    - [Generics — Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)
    - [Contracts — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md)
- エラー・ハンドリングに関するもの
    - [Error Handling — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md)
    - [Error Inspection — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md)
    - [Error Printing — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-printing.md)

このうち今回は総称型について予習してみる。

なお “Go 2” の提案はまだドラフト段階なので大幅に変更になったり場合によっては立ち消えになる可能性もある。
なので，この記事では深いところまで踏み込まずフワっとした説明になるけど，あしからずご了承の程を。

## 総称型のメリット

ソフトウェア・エンジニアには自明だと思うが，まずは復習から。

具体例として2つの値のうち大きい方を返す関数を考えてみる。

{{< highlight go "hl_lines=5-10" >}}
package main

import "fmt"

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func main() {
    x := 1
    y := 2
    fmt.Printf("max(%v, %v) = %v\n", x, y, max(x, y))
    //Output
    //max(1, 2) = 2
}
{{< /highlight >}}

この関数 `max()` は int 型で記述しているが byte 型や float32/float64 型でも関数の中身は全く同じコードになる。

{{< highlight go "hl_lines=5-10" >}}
package main

import "fmt"

func max(x, y float64) float64 {
    if x < y {
        return y
    }
    return x
}

func main() {
    x := 1.1
    y := 1.2
    fmt.Printf("max(%v, %v) = %v\n", x, y, max(x, y))
    //Output
    //max(1.1, 1.2) = 1.2
}
{{< /highlight >}}

ならば，最初から汎化した型で単一のコード記述すれば型ごとに複数のコードを量産しなくてもいんじゃね？ という発想になる[^oaoo1]。
これが総称型の原点である。

[^oaoo1]: このように同じコードを重複させないように記述するコーディング指針を「OAOO (Once And Only Once) 原則」と呼ぶ。そういえばよくある勘違いで「DRY (Don't Repeat Yourself) 原則」と解説している記事が見受けられるが， DRY 原則は同じ意味を持つ情報やデータを複数に散らばせないというシステム設計や開発環境の指針を指すものである。

たとえば Java で総称型を使うと以下のような記述になる[^java1]。

[^java1]: 久しぶりに Java コード書いたら型を前置することに違和感がハンパない。型の前置なんて非合理的だよなあ（もちろん偏見`w`）

```java
public static <T> T max(T x, T y) {
    ...
}
```

`<T>` の部分が総称型の定義にあたる。
ちなみに総称型の名前（この場合は `T`）はスコープ内で被らなければ任意に指定できる。

現行版 [Go 言語]において総称型の不在で割りを食っている典型例のひとつが [`sort`] パッケージで，基本型の slice のソートだけで以下の型が用意されている[^sort1]。

[^sort1]: Slice のソートについてはバージョン 1.8 から任意の型に対応する [`sort`]`.Slice()` 関数が用意された。内部で [`reflect`] パッケージを使っているが，かなり巧妙に組まれているため，パフォーマンス低下は殆どないらしい。ただし slice 以外のインスタンスを指定すると（コンパイル時ではなく）実行時に panic を吐く。詳しくは「[ソートを使う]({{< relref "sort.md" >}})」を参照のこと。

- [`sort`]`.Float64Slice`
- [`sort`]`.IntSlice`
- [`sort`]`.StringSlice`

後方互換性を確保するため，これらの型がなくなることはないだろうけど，総称型が実現すれば内部実装の refactoring が進むかも知れない。

このように総称型はコンテナ（container; オブジェクトの集合を表現するデータ構造）操作で特に威力を発揮する。
また，現行版 [Go 言語]では総称型を用いずとも interface 型と [`reflect`] パッケージを使ってかなりの部分を代替できるが，コード量のコスト高を別にしても，記述の正しさは実行時での評価ではなくコンパイル時に評価して欲しいところである。

そういうわけで，今までずうっと後回しにされてきたが，総称型を導入できるのであれば是非とも期待したいものである。

## 型パラメータ（Type Parameter）と型引数（Type Argument）

次期 [Go 言語]で総称型を導入するために型パラメータおよび型引数の構文を追加するようだ。

例えば先程の `max()` 関数であれば以下のように記述できる。

{{< highlight go "hl_lines=5" >}}
package main

import "fmt"

func max(type T)(x, y T) T {
    if x < y {
        return y
    }
    return x
}

func main() {
    x := 1
    y := 2
    fmt.Printf("max(%v, %v) = %v\n", x, y, max(x, y))
    //Output
    //max(1, 2) = 2
}
{{< /highlight >}}


`(type T)` の部分が型パラメータで，これによって総称型を定義している。
既存の語彙だけで構成しているのがポイント（`<` や `>` は演算子だし `[...]` は 配列/slice や map の構文で使われるので避けたのだろう）。

関数を呼び出す側は型推論によって引数の型が一意に決定するので特別な記述は必要ない。
相変わらず refatoring に優しい言語だよな（笑）

明示的に型を指定するなら

```go
max(int)(x, y)
```

などと記述する。
`(int)` の部分が型引数にあたる。

型宣言とインスタンス生成でも同様に

```go
type List(type T) []T

var listInt = List(int){0, 1, 2, 3, 4 ,5}
```

などと記述できる。

## 型コントラクト（Type Contract）

先程の `max()` 関数だが

```go
func max(type T)(x, y T) T {
    if x < y {
        return y
    }
    return x
}
```

型 `T` のインスタンス同士で比較演算（具体的には `x < y`）が可能である必要がある。
たとえば complex64/complex128 は基本型だが `<` 演算子による比較ができない。

Java の場合は継承を構成して

```java
public static <T extends Comparable<? super T>> T max(T x, T y) {
    ...
}
```

などと記述することで総称型に対する制約（type constraint）を表現できる[^con1]。

[^con1]: C++ や C# でも where 句を用いて総称型に対する制約を表現できるが，基本は継承を利用したものである。

しかし [Go 言語]ではいわゆる「継承」の仕組みを持ってないため別のアプローチをとる必要がある。
それが型コントラクトである。
型コントラクトでは contract キーワードおよびそれを使った構文を追加する。
具体的には以下のようなコードになる。

{{< highlight go "hl_lines=1-3 5" >}}
contract comparable(t T) {
    t < t
}

func max(type T comparable)(x, y T) T {
    if x < y {
        return y
    }
    return x
}
{{< /highlight >}}

なお `comparable` の型引数を明示する場合は

```go
func max(type T comparable(T))(x, y T) T { ... }

```

と書く。

型コントラクトの記述はバイナリ・コードにコンパイルされない。
上の例では型 `T` に対して比較演算子 `<` が使えることを要求しているとコンパイラに知らせるものである。
これなら `T` を complex128 に展開しようとしてもコンパイル時に「契約違反」になるわけだ。

型コントラクトは型コントラクトに埋め込むことができる。
例えばこんな感じ。

{{< highlight go "hl_lines=5" >}}
contract equalable(t T) {
    t == t
}
contract comparable(t T) {
    equalable(T)
    t < t
}
{{< /highlight >}}

これで `comparable` は `==` 演算子と `<` 演算子が使えることを要求していることになる。

継承を利用した制約と異なり，型コントラクトの自由度は高く応用範囲も広い。
たとえばポインタ型を要求するなら

```go
contract pointer(t T) {
    *t
}
```

などと書くこともできるらしい。
他にも [`io`]`.Reader` と汎化・特化の関係があることを要求するなら

```go
contract readable(r T) {
    io.Reader(r)
}
```

などと書けばいいようだ。

もし今回のドラフト案の通りに総称型が実現するなら型コントラクトの整備が喫緊の作業となるだろう。

## ブックマーク

- [Go 2のgenerics/contract簡易まとめ](https://qiita.com/lufia/items/242d25e8c93d88e22a2e)
- [Why Generics? - The Go Blog](https://blog.golang.org/why-generics)

- [Go 言語における「オブジェクト」]({{< relref "object-oriented-programming.md" >}})
- [きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」]({{< ref "/remark/2017/03/generics-vs-duck-typing.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[`sort`]: https://golang.org/pkg/sort/ "sort - The Go Programming Language"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
