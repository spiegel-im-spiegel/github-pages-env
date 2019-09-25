+++
title = "継承を捨てて「オブジェクト指向」を始めよう"
date =  "2019-09-25T09:41:27+09:00"
description = "description"
image = "/images/attention/kitten.jpg"
tags = [ "object-oriented", "design", "engineering", "golang" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日，面白い記事を見つけた。

- [Go言語のInterfaceの考え方、Accept interfaces,return structs - Qiita](https://qiita.com/weloan/items/de3b1bcabd329ec61709)

これって，私くらいの「ちょっと [Go 言語]に慣れた」程度の人であれば無意識にやっていることだと思うが，こうして改めて言語化されるととても納得できる内容である。
そして今後は “accept interfaces, return structs” を [Go 言語]の主要なデザインパターンとして意識的に書けるようになるだろう。

まぁ，詳しい内容はこの記事を読んでもらうとして，この記事ではちょっとポエムに走ってみようか（笑）

## [Go] の interface 型は Duck Typing ではない（らしい）

確か [Go 言語]関連の公式ドキュメントでもあちこちに書かれている気がするし，私も割と最近までそう思っていたのだが，実は [Go] の interface 型は [duck typing] ではないらしい。

そもそも [duck typing] は Ruby などが採用している「動的型付け」システムにおける型推論の手法（のひとつ）である。
その名前は [duck test] から来ていて

{{< fig-quote type="md" title="Duck test - Wikipedia" link="https://en.wikipedia.org/wiki/Duck_test" lang="en" >}}
{{< quote >}}If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.{{< /quote >}}
{{< /fig-quote >}}

というフレーズに集約されている。
つまり，それが『にゃーん』と鳴くのなら，机器猫だろうが猫耳メイドだろうがサーバルキャットだろうが全部「猫」である，ということだ。

[Go] の interface 型もこうした帰納法的な機能を持つが，こちらは「構造的部分型（structural subtyping）と呼ぶらしい。
ちなみに構造的部分型に対して C++ や Java のテンプレート・クラスやインタフェース・クラスは「公称型（nominal subtyping）」とか呼ぶそうな。

[Go] の interface 型の特徴は，単なる型推論にとどまらず，オブジェクト間の「汎化」関係を記述する強力な手段である，ということだ。

## 「継承」ではなく「汎化」で考える

オブジェクト間の関係を表すものはいくつかあるが，特に「オブジェクト指向」っぽいのは「汎化・特化」あるいは “is-a” の関係だろう。
C++ や Java のようなレガシーなオブジェクト指向プログラミング言語は，この「汎化・特化」の関係を「継承」で表す。

継承の考え方は基本的にはまず親クラスがあって，親クラスの性質や機能を「継承」した子クラスがある。
そして更にその下に小クラスの小クラスがある，といった感じの木構造になっている。
逆に言うとこの形に当てはまらない構造は「継承」を使っては実装できない。

しかし実際にオブジェクト指向設計を要件定義のレベルから検討していくと分かるが，木構造の「継承」関係を構成するのは（特定のデザインパターンに嵌まれば簡単だが，大抵はそうならないため）結構難しいのだ。
それは人間がシステムを設計する際は 抽象→具象 と考えるのではなく 具象→抽象 で考えるからだ。

まずはシステムに登場する（アクターを含む）具象化されたモノや事象を全て列挙した上でそれらの関係を考察し，その過程でそれまで見えなかった抽象的なモノが定義可能になる。

実は [Go] の interface 型は  具象→抽象 という思考の進め方と相性がいい。

例えば，最初に机器猫とか猫耳メイドとかサーバルキャットとかいった型をつくり，「これってみんな『にゃーん』って鳴くぢゃん」と気付き，それらをまとめて扱いたいと考えれば「『にゃーん』と鳴く」機能を持つ interface 型を作って「猫」とラベリングすればいいわけだ。
同様に機能・操作に対して「ロボット」とか「フレンズ」とか「メイドさん」とか「ダーウィンが来た！」とかラベリングする。

定義した interface 型の間には関係があってもいいしなくてもいい。
何故なら interface 型の構造を決めるのは具象型を定義する側ではなく利用する側だからだ。

## 依存の注入

最近「[Go の疑似乱数生成器は Goroutine-Safe ではないらしい]({{< ref "/golang/pseudo-random-number-generator.md" >}})」という記事を書いたので，これを例に話をしてみよう。

標準パッケージの [math/rand] は疑似乱数生成器を生成・操作するためのパッケージで，以下のように定義されている（コメントは端折ってるけど，ご容赦を）。

```go
type Source interface {
	Int63() int64
	Seed(seed int64)
}

type Source64 interface {
	Source
	Uint64() uint64
}

type Rand struct {
	src Source
	s64 Source64
	readVal int64
	readPos int8
}
```

[`rand`]`.Source` および [`rand`]`.Source64` が interface 型である。
[`rand`]`.Source64` は [`rand`]`.Source` を包含しているのがポイントである。

そして乱数生成器の生成は以下の [`rand`]`.New()` 関数で行う。

```go
func New(src Source) *Rand {
	s64, _ := src.(Source64)
	return &Rand{src: src, s64: s64}
}
```

引数の `src` には `Int63() int64` および `Seed(int64)` のメソッドを持つ型であればどんなインスタンスでも受け入れる。
つまり [`rand`]`.New()` 関数で依存の注入（dependency injection）を行うわけだ。

また，この中の `src.(Source64)` の部分は型変換の構文で，引数の `src` を [`rand`]`.Source64` 型に変換して格納している。
如何にも後付け臭い記述だが，これで `Uint64() uint64` メソッドがある型でもない型でも

```go
func (r *Rand) Uint64() uint64 {
	if r.s64 != nil {
		return r.s64.Uint64()
	}
	return uint64(r.Int63())>>31 | uint64(r.Int63())<<32
}
```

のように対応できる。

ポイントは「どのような型を受け入れるかは（依存を）受け入れる側が決める」ことだ。
これが “accept interfaces” の意味するところである。

これを実装型を提供する側から見ると，受け入れ側が提示するインタフェースさえ満たしていれば自由に実装できる。
例えば疑似乱数生成器として [`mt`]`/mt19937` パッケージを自作する場合でも






[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"
[duck typing]: https://en.wikipedia.org/wiki/Duck_typing "Duck typing - Wikipedia"
[duck test]: https://en.wikipedia.org/wiki/Duck_test "Duck test - Wikipedia"
[math/rand]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`mt`]: https://github.com/spiegel-im-spiegel/mt "spiegel-im-spiegel/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented"
<!-- eof -->
