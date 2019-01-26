+++
title = "Go 1.9 と Type Alias"
date =  "2017-09-14T10:01:49+09:00"
update = "2018-01-23T16:07:00+09:00"
description = "なんで type alias なんて妙ちきりんな言語仕様が追加されたかというと，実はこれ，リファクタリングの為に設けられたのである。"
tags        = [ "programming", "language", "golang", "type", "refactoring" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

遅ればせながらの記事で申し訳ないが Go 1.9 がリリースされた。

- [Go 1.9 is released - The Go Blog](https://blog.golang.org/go1.9)
- [Go 1.9 Release Notes - The Go Programming Language](https://golang.org/doc/go1.9)

詳しい内容は[リリースノート]を見ていただくとして，今回の目玉は [type] alias 機能だろう。

まず [type] キーワードを使った簡単な足し算を書いてみる。

```go
package main

import (
    "fmt"
)

type Num1 int

func (n1 Num1) Add(n2 Num1) Num1 {
    return n1 + n2
}

func main() {
    n1 := Num1(1)
    n2 := Num1(2)
    fmt.Println(n1.Add(n2))
}
```

実行結果は 3 と出力されるはずである。
ここで “`type Num2 Num1`" と記述を追加し，この型を使って足し算を行ってみる。

```go
package main

import (
    "fmt"
)

type Num1 int

func (n1 Num1) Add(n2 Num1) Num1 {
    return n1 + n2
}

type Num2 Num1

func main() {
    n1 := Num2(1)
    n2 := Num2(2)
    fmt.Println(n1.Add(n2))
}
```

これを実行しようとすると

```text
n1.Add undefined (type Num2 has no field or method Add)
```

とコンパイルエラーになる。
何故か。
`Num1` と `Num2` は異なる型だからだ。
型 `Num1` に紐付いている関数 `Add()` は，型 `Num2` には紐付かない。
継承されないわけだ（`Num1` へキャストはできる）。

では今度は “`type Num2 = Num1`" と記述を変更してみる。

```go
package main

import (
    "fmt"
)

type Num1 int

func (n1 Num1) Add(n2 Num1) Num1 {
    return n1 + n2
}

type Num2 = Num1

func main() {
    n1 := Num2(1)
    n2 := Num2(2)
    fmt.Println(n1.Add(n2))
}
```

今度はコンパイルエラーにならず 3 と出力される。
この “`type Num2 = Num1`" という構文が [type] alias を指し，この記述によって `Num1` と `Num2` は **全く同じ型** として扱われる[^a1]。

[^a1]: 全く同じ型なので継承関係はなく，別名定義した型に独自に関数を紐付けることはできない。ちなみに別パッケージの型に対しても別名定義が可能である： `type Time = time.Time`

なんでこんな妙ちきりんな言語仕様が追加されたかというと，実はこれ，リファクタリングの為に設けられたのである。

{{< fig-quote title="Go 1.9 Release Notes" link="https://golang.org/doc/go1.9" lang="en" >}}
<q>Go now supports type aliases to support gradual code repair while moving a type between packages.
The type alias design document and an article on refactoring cover the problem in detail.</q>
{{< /fig-quote >}}

もともと [Go 言語]はリファクタリングを厚遇する言語と言える。
たとえば「構造的部分型（structural subtyping）」などはその最たる例だろう。

- [きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」]({{< ref "/remark/2017/03/generics-vs-duck-typing.md" >}})

まぁ，あまり積極的に使う機能ではないかもしれないが，こういうこともできると覚えておくといいだろう。

## ブックマーク

- [go言語1.9で追加予定の新機能 型エイリアス - Qiita](http://qiita.com/weloan/items/8abbb4003cfa1031a9e9)

- [Go 言語における「オブジェクト」]({{< relref "object-oriented-programming.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[リリースノート]: https://golang.org/doc/go1.9 "Go 1.9 Release Notes - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%AA%E3%83%95%E3%82%A1%E3%82%AF%E3%82%BF%E3%83%AA%E3%83%B3%E3%82%B0%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1EU?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1EU"><img src="https://images-fe.ssl-images-amazon.com/images/I/41GPVATQiZL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%AA%E3%83%95%E3%82%A1%E3%82%AF%E3%82%BF%E3%83%AA%E3%83%B3%E3%82%B0%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1EU?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1EU">Java言語で学ぶリファクタリング入門</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2007-01-26 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1EU</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">結城浩さんによる「リファクタリング本」の Kindle 版。意外にも Java 以外でも応用できる優れもの。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-11">2018-12-11</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
