+++
title = "Go 言語で Factory Method Pattern を構成できるか"
date = "2018-12-30T15:24:48+09:00"
update = "2018-12-31T09:01:54+09:00"
description = "継承できないなら注入すればいいじゃない。"
image = "/images/attention/go-logo_blue.png"
tags = ["golang", "object-oriented", "programming", "type", "interface"]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

少し前に面白い記事を見かけたのだが

- [Go言語で factory method パターンを書いてみる - Qiita](https://qiita.com/daiching/items/e986ccd8a3f18c56c406)

「これって factory method pattern じゃないよね」と思いつつ「[Go 言語]には「継承」がない」ことを説明する好例だと気付いたので今回の記事を書いてみた。

## Factory Method Pattern とは

まずは factory method pattern の解説から。
以下にクラス図を示す。

{{< fig-img src="./factory-method-pattern.png" link="./factory-method-pattern.puml" width="1165" >}}

`Product` は抽象クラスで `ConcreteProduct` が `Product` の実装クラスである。
`Creator` クラスの `factoryMethod()` メソッドは `Product` インスタンスを返す抽象メソッドで，サブクラスの `ConcreteCreator` で実装されている。

`Creator` クラスを利用する側はサブクラスの `ConcreteCreator` を `Creator` クラスとして生成し `anOperation()` メソッドをキックする。

ポイントは `anOperation()` メソッドが実装メソッドで，内部では `factoryMethod()` メソッドが呼ばれている点である。
ただし `Creator` インスタンスの実体は `ConcreteCreator` クラスなので， `anOperation()` メソッドから呼ばれるのはオーバーライドされた `ConcreteCreator` のほうの `factoryMethod()` であることが期待される，というわけだ。

本来なら `Creator` インスタント生成時にどの `ConcreteProduct` インスタンスを（`Product` クラスとして）使うのか判断すればいいのだが（これが factory pattern），それだと `ConcreteProduct` クラスが増えるたびに `Creator` クラスも修正しなければならず管理が煩雑になる。

そこで `ConcreteProduct` クラスと一対一に対応する `ConcreteCreator` クラスをつくり，その中でどの `ConcreteProduct` インスタンスを（`Product` クラスとして）使うのか判断している，つまり「[インスタンス作成をサブクラスにまかせる](https://www.amazon.co.jp/exec/obidos/ASIN/B00I8ATHGW/baldandersinf-22 "増補改訂版 Java言語で学ぶデザインパターン入門 | 結城 浩 | コンピュータ・IT | Kindleストア | Amazon")」のである。

## [Go 言語]には「継承（Inheritance）」がない

しかし，このデザイン・パターンは [Go 言語]では使えない。
何故なら「[Go 言語]には「継承」がない」からである。

### 汎化と構造的部分型（Structural Subtyping）

[Go 言語]ではクラス間の「汎化」関係を実装する手段として interface 型を持っているが，これは「振る舞い」のみを定義する。
先程の `Creator` クラスのように抽象メソッドと実装メソッドをひとつのクラスの中に混ぜ込むことはできない。
Interface のような型を一般的には「構造的部分型」と呼ぶらしい[^dt1]。

[^dt1]: Interface 型の性質を示すものとして [duck typing] がよく挙げられる。私も以前はそう思っていたが [duck typing] は動的型付き言語で型を決定する手法（のひとつ）を指すのだそうだ。もちろん [Go 言語]は静的型付き言語なのでこれに当てはまらない。じゃあ [Go 言語]の interface 型は何かというと[「構造的部分型」と呼ぶのが正しい](https://methane.hatenablog.jp/entry/ruby-python-go-typing "Python と Ruby と typing - methaneのブログ")ようだ。

### 埋め込み（Embedding）と委譲（Delegation）

もうひとつ [Go 言語]には型を構造体のフィールドとして埋め込む機能がある。
たとえば

```go
type parent struct{}
```

という型があるとして

```go
type child struct {
    parent
}
```

とすれば `child` 型は `parent` 型の属性や操作をそのまま使うことができる。
こんな感じ。

```go
package main

import "fmt"

type parent struct{}

func (p parent) Say() string {
    return "I'm parent type."
}

type child struct {
    parent
}

func main() {
    p := parent{}
    c := child{}
    fmt.Println(p.Say())
    fmt.Println(c.Say())
    //Output:
    //I'm parent type.
    //I'm parent type.
}
```

これで `child` 型は `parent` 型を（擬似的に）継承しているように見える。
更に `child` 型に対して

```go
func (c child) Say() string {
    return "I'm child type."
}
```

という関数を追加すれば

{{< highlight go "hl_lines=15-17" >}}
package main

import "fmt"

type parent struct{}

func (p parent) Say() string {
    return "I'm parent type."
}

type child struct {
    parent
}

func (c child) Say() string {
    return "I'm child type."
}

func main() {
    p := parent{}
    c := child{}
    fmt.Println(p.Say())
    fmt.Println(c.Say())
    //Output:
    //I'm parent type.
    //I'm child type.
}
{{< /highlight >}}

という感じになり `child` 型の `Say()` 関数が `parent` 型の `Say()` 関数をオーバーライドしているように見える。
ところが `parent` 型に

```go
func (p parent) Speek() {
    fmt.Println(p.Say())
}
```

という関数を追加し，これを使って喋らせてみると


{{< highlight go "hl_lines=11-13 26" >}}
package main

import "fmt"

type parent struct{}

func (p parent) Say() string {
    return "I'm parent type."
}

func (p parent) Speek() {
    fmt.Println(p.Say())
}

type child struct {
    parent
}

func (c child) Say() string {
    return "I'm child type."
}

func main() {
    c := child{}
    fmt.Println(c.Say())
    c.Speek()
    //Output:
    //I'm child type.
    //I'm parent type.
}
{{< /highlight >}}

となり `Speek()` 関数はあくまでも `parent` 型の `Say()` 関数を呼び出していることが分かる。

実は [Go 言語]の埋め込みフィールドは「継承」ではなく「委譲」として機能する。
したがって C++ や Java の仮想関数のようにメソッドがオーバーライドされることもない。
型の中で関数名が被った際の暗黙の優先順位は決められているためオーバーライドしているように見えるだけなのである。

これを先程の factory method pattern のクラス図に当てはめて考えると [Go 言語]で無理やり実装しようとしても

{{< fig-img src="./delegation.png" link="./delegation.puml" width="1306" >}}

という感じになり，そもそも `Creator` を抽象クラスに（polymorphic に）できないし，サブクラスの `ConcreteCreator` から `anOperation()` メソッドを呼び出しても内部で呼び出される `factoryMethod()` メソッドは `Creator` クラスのものということになる。

## 「継承」できないなら「注入」すればいいじゃない

さて，どうしようか。
とりあえずコードを書きながら考えてみよう。

最初に紹介した記事を参考にしつつ，まずは `Product` に相当する interface 型を以下のように定義するところから始めようか。

```go
//Human is Product class of factory method pattern
type Human interface {
    SelfIntroduction() string
}
```

続いて `ConcreteProduct` に相当する型も定義する。

```go
//NamedHuman is parent class with name property
type NamedHuman struct {
    MyName string
}

func (nh NamedHuman) Name() string {
    if len(nh.MyName) == 0 {
        return "〈名無し〉"
    }
    return nh.MyName
}

//Man is Concrete Product class of factory method pattern
type Man struct {
    NamedHuman
}

func (hm Man) SelfIntroduction() string {
    return hm.Name() + "だぜ！"
}

//woman is Concrete Product class of factory method pattern
type Woman struct {
    NamedHuman
}

func (hm Woman) SelfIntroduction() string {
    return hm.Name() + "よ！"
}
```

ちょっと分かりにくいかも知れないが，クラス図にすると

{{< fig-img src="./human.png" link="./human.puml" width="1134" >}}

となっていて `Man` 型および `Woman` 型が（`Human` 型に対する） `ConcreteProduct` クラスに相当するのが分かると思う。
ちょろんと動かしてみよう。

```go
func NewMan(name string) Human {
    m := &Man{}
    m.MyName = name
    return m
}
func NewWoman(name string) Human {
    w := &Woman{}
    w.MyName = name
    return w
}
func main() {
    h1 := NewMan("太郎")
    h2 := NewWoman("花子")
    fmt.Println(h1.SelfIntroduction())
    fmt.Println(h2.SelfIntroduction())
    //Output:
    //太郎だぜ！
    //花子よ！
}
```

[問題なく動作する](https://play.golang.org/p/KYQV48YV2pz "The Go Playground")ことが確認できた。

ここまで書いたコードを眺めてみると `NewMan()` および `NewWoman()` 関数が `factoryMethod()` メソッドと似た機能を有していることが分かる。
じゃあ，これらの関数を実行時に `Creator` クラスに埋め込むように構成すればいんじゃね？

さっそく書いてみよう。

```go
//Speaker is Creator class of factory method pattern
type Speaker struct {
    createHuman func(string) Human
}

func (s Speaker) Speech(name string) {
    hm := s.createHuman(name)
    fmt.Println(hm.SelfIntroduction())
}
```

`createHuman` は関数型のメンバ変数で `func(string) Human` のインタフェースを持つ関数を表す。
これはすなわち `NewMan()` および `NewWoman()` 関数のことである。

それでは実際に動かしてみよう。

```go
func NewSpeeker(f func(string) Human) *Speaker {
    return &Speaker{createHuman: f}
}

func main() {
    s1 := NewSpeeker(NewMan)
    s2 := NewSpeeker(NewWoman)
    s1.Speech("太郎")
    s2.Speech("花子")
    //Output:
    //太郎だぜ！
    //花子よ！
}
```

[よーし，うむうむ，よーし](https://play.golang.org/p/_QFnKteL_c6 "The Go Playground")。
クラス図にするとこんな感じだろうか（factory method pattern のクラス図と比較してみよう）。

{{< fig-img src="./injection.png" link="./injection.puml" width="1168" >}}

これってだいぶ不格好だけど，まさしく依存（オブジェクト）の注入（dependency injection pattern）だよね[^di1]。

[^di1]: 私は dependency injection を「依存性の注入」と訳すことに懐疑的だという[主張]に同意します。

ていうか [Go 言語]プログラマは error 型や標準パッケージの [`io`].`Reader` などを通じて息をするように「依存の注入」を行っているわけで，「継承できないなら注入すればいいじゃない」という感じで，自国語に得意なパターンへ持ち込むのが得策だと思う。

## ブックマーク

- [Factory Method パターン - Wikipedia](https://ja.wikipedia.org/wiki/Factory_Method_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3) : クラス図はこちらを参考にした
- [Dependency injection - Wikipedia](https://en.wikipedia.org/wiki/Dependency_injection) : 英語版と日本語版でかなり内容が異なるのだが大丈夫か？

- [Go 言語における「オブジェクト」]({{< ref "./object-oriented-programming.md" >}})

<!--
- [The Go Playground](https://play.golang.org/p/lQlNziXLf9F)
-->

[Go 言語]: https://golang.org/ "The Go Programming Language"
[duck typing]: https://en.wikipedia.org/wiki/Duck_typing "Duck typing - Wikipedia"
[主張]: http://blog.a-way-out.net/blog/2015/08/31/your-dependency-injection-is-wrong-as-I-expected/ "やはりあなた方のDependency Injectionはまちがっている。 — A Day in Serenity (Reloaded) — PHP, FuelPHP, Linux or something"
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E5%A2%97%E8%A3%9C%E6%94%B9%E8%A8%82%E7%89%88-Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8ATHGW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8ATHGW"><img src="https://images-fe.ssl-images-amazon.com/images/I/41mh5r0NwLL._SL160_.jpg" width="126" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E5%A2%97%E8%A3%9C%E6%94%B9%E8%A8%82%E7%89%88-Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8ATHGW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8ATHGW">増補改訂版 Java言語で学ぶデザインパターン入門</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2004-06-18 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8ATHGW</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">結城浩さんによる通称「デザパタ本」の Kindle 版。 Java 以外でも使える優れもの。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-01-05">2016-01-05</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
