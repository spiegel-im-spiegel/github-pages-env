+++
title = "『プログラミング言語 Go』を眺める"
date = "2016-07-13T04:19:59+09:00"
description = "Go 言語をある程度使える（私程度のレベル）という方は第7章から第9章までを重点的に読むといい。"
tags = ["book", "programming", "language", "golang"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]のよいところのひとつは[ドキュメント](http://golang-jp.org/doc/ "")が分かりやすい形で提供されている点である。
はじめて [Go 言語]に接する人なら “[A Tour of Go]” から気楽に始められるし，手元にコンパイラがなくてもとりあえず “[Playground](https://play.golang.org/ "The Go Playground")” で遊ぶことはできる。
更に[言語仕様](https://go.dev/ref/spec)や[標準パッケージのドキュメント](https://golang.org/pkg/)や [FAQ](https://golang.org/doc/faq) といったものはもちろん， “[Effective Go](https://golang.org/doc/effective_go.html)” みたいなドキュメントも用意されている。
なので実際のところ，本を買わなくてもオンラインで充分学ぶことができる。

にも関わらず今回この本を買った理由は以下の2つ。

1. オフラインで参照できる完全なリファレンス本が欲しかった
2. 知識体系として [Go 言語]を学べる教科書が欲しかった

実は現在，職場が某セキュリティ・エリアの中にあってネットから物理的に切り離された環境にいる。
以前は分からないことは Google 先生に訊けたのに，それが出来なくなってしまったのだ。
スマホやタブレットといった電子機器も持ち込めないので「Kindle で」というわけにもいかない。
今ほど紙の本のありがたみを実感できたことはないよ。

まぁそういう経験をしてしまったので（仕事に絡みそうなものであれば）紙の本はちゃんと買っておくか，という気になったのだった。
逆に紙の本を買ってしまったので（今までみたいな余暇の遊びじゃなく）元を取らないとダメだなぁ，などと思ったり。

さて本題。

まず， [Go 言語]を全くはじめて習うという方は，『[プログラミング言語 Go]』からはじめるのではなく， “[A Tour of Go]” （[日本語版](https://go-tour-jp.appspot.com/)もある）からはじめることを強くおすすめする。
“[A Tour of Go]” では手を動かしながら学べるので「感触」を掴むのにちょうどよい教材と言える。
そうした後に『[プログラミング言語 Go]』を読み進めれば理解しやすいと思う。

[Go 言語]をある程度使える（私程度のレベル）という方は第7章から第9章までを重点的に読むといい。
Interface, goroutine, channel の概念や実装方法は [Go 言語]の中核技術といえるもので，ここを押さえておけばかなり使いこなせるようになるはず。
また第12章以降に登場する refrection や [`unsafe`](https://golang.org/pkg/unsafe/) パッケージの説明は個人的にかなり分かりやすかった。

第8章および第9章を読んでいて気がついたのだが， goroutine/channel を使った CSP (Communicating Sequential Processes) の真価は data driven な設計で真価を発揮するのではないだろうか。
「並行プログラミング」を意識するとどうしてもスレッドを連想してしまうけど，スレッドよりも遥かに軽量な goroutine はもっと無茶ができるはず。
たとえば多数の goroutine をネットワーク化した非ノイマン型っぽい「何か」とか。

## 参考

- [プログラミング言語 Go](/golang/) : 不定期に書いてます
- [なぜGo言語は設計が悪いのか – Go愛好者の見地から | 未分類 | POSTD](http://postd.cc/why-go-is-a-poorly-designed-language/)
- [6年間におけるGoのベストプラクティス | プログラミング | POSTD](http://postd.cc/go-best-practices-2016/)
- [Big Sky :: golang の channel を使ったテクニックあれこれ](http://mattn.kaoriya.net/software/lang/go/20160706165757.htm)
- [ukai's blog: 『プログラミング言語Go』刊行記念イベント「Goの設計思想を読み解く～実際の開発に活かすために」](http://blogger.ukai.org/2016/07/gogo.html)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[A Tour of Go]: https://tour.golang.org/
[プログラミング言語 Go]: https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/ "プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES) : Alan A.A. Donovan, Brian W. Kernighan, 柴田 芳樹 : 本 : Amazon"

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
