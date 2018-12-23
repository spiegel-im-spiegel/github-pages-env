+++
title = "「もう参照渡しとは言わせない」とかメンドくせーんだよ"
date =  "2017-11-05T17:39:04+09:00"
description = "description"
tags        = [ "communication", "golang" ]
draft = true

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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

最近，微妙に Qiita 界隈が騒がしいみたいで，「もう参照渡しとは言わせない」とか面倒くさいことを言ってる連中がいるらしい。

最初に言っておくと，議論というのは予めドメインとレイヤをきっちり決めておくことが重要である。
それをしないでふわっとした主張をするから発散してしまうのである。
今回の件のドメインを Java 界隈に限るとしても，概念上の話をしているのか，言語仕様の話をしているのか，更に下位の実装上の話をしているのか，などで全く変わってくる。

殆どの人は前提無しで出てくる「参照渡し」を概念的な意味でしか捉えていない。
つまり一般的な関数スコープの話であろうと考えるわけだ。
そこで Java の言語仕様はポインタや参照が云々とか言っても話が噛み合うわけがない。

こういう議論以下の言い合いに巻き込まれるのはご免なので，[手持ちの記事]({{< ref "/golang/function-and-pointer.md" >}})は書き換えた。
やれやれである。
こういう実りの少ないリファクタリングは徒労感が半端ない。

折角なので少し回収しておく。

## C 言語は「値渡し」

ポインタの話が出るなら Java より前に C 言語で [K&R] だろ，常考。

[K&R] では C 言語は call by value だと明記されているが，引数にポインタを渡すことを call by reference だとは一言も書かれていない。
Call by reference は Fortran などの既存の言語で使われるシーケンスであり， C 言語は{{< ruby "ザク" >}}奴等{{< /ruby >}}とは違うのだよ，と謳っている。

{{< fig-quote title="プログラミング言語C" link="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/" >}}
<q>この “call by value” は資産であり負債ではない。通常は余分な変数はほとんど使わずにすみ，プログラムはよりコンパクトになる。というのは，呼ばれたサブルーチン内では，パラメータは，都合よく初期化された局所的変数として扱えるからである。</q>
{{< /fig-quote >}}

その上で「呼ばれた関数のほうで，呼び出した関数内の変数を直接変更することはできない」ことへの補償として引数にポインタを渡すことが示されているに過ぎない。
つまり「値渡し」こそが C 言語を特徴づけるもののひとつだ，ということなのである。


## [Go 言語]も「値渡し」 ...なのだが

[Go 言語]もまた「値渡し」であることが明記されている。

{{< fig-quote title="The Go Programming Language Specification" link="https://golang.org/ref/spec#Calls" lang="en" >}}
<q>In a function call, the function value and arguments are evaluated in <a href="https://golang.org/ref/spec#Order_of_evaluation">the usual order</a>. After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution. The return parameters of the function are passed by value back to the calling function when the function returns.</q>
{{< /fig-quote >}}

そして関数スコープ外のインスタンスへアクセスするために引数にポインタを渡せる点も同じである。
C++ 言語の「参照」型は，言語仕様上は存在しない。

のだが，しかし。
[Go 言語]では引数のインスタンスが何処に置かれるかはコンパイル時に決定され明示されない。
通常は [string] を除く基本型[^bt1] は値が直接スタックに積まれ，それ以外はヒープ上にインスタンスを生成し，そのポインタがスタックに積まれる。

[^bt1]: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128 が [Go 言語]に於ける基本型。























## ブックマーク

- [なぜ、Java等で「参照渡し」の言葉を使うべきではないのか？ - Qiita](https://qiita.com/raccy/items/59a6ac6c818918dd9651)

[K&R]: http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/ "プログラミング言語C 第2版 ANSI規格準拠 | B.W. カーニハン, D.M. リッチー, 石田 晴久 |本 | 通販 | Amazon"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[string]: http://golang.org/ref/spec#String_types
