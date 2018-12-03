+++
tags = ["golang", "programming", "language"]
description = "Go 言語コンパイラのバージョン 1.8 がリリースされた。"
date = "2017-02-19T15:45:53+09:00"
update = "2017-05-02T12:35:52+09:00"
title = "Go 言語 1.8 がリリース"
draft = false

[author]
  flickr = "spiegel"
  avatar = "/images/avatar.jpg"
  url = "http://www.baldanders.info/spiegel/profile/"
  name = "Spiegel"
  flattr = "spiegel"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  tumblr = ""
  facebook = "spiegel.im.spiegel"
  instagram = "spiegel_2007"
+++

[Go 言語]コンパイラのバージョン 1.8 がリリースされた。

- [Go 1.8 is released - The Go Blog](https://blog.golang.org/go1.8)

主な改善点を上げると

- コンパイル時間の短縮
- ガベージコレクションの改善（通常で 100μsec 未満，しばしば 10μsec 未満）
- HTTP/2 Push の追加
- 標準ライブラリの [`context`] パッケージについてキャンセルとタイムアウトの仕組みを追加
    - HTTP サーバのシャットダウンの改善など
- [`sort`].[`Slice`](https://golang.org/pkg/sort/#Slice) 関数の追加。 [slice] のソートが簡単になった

その他，詳しくは以下のリリースノートを参照のこと。

- [Go 1.8 Release Notes - The Go Programming Language](https://golang.org/doc/go1.8)

## 戯れ言

2015年頃から（仕事以外でだが） [Go 言語]で遊ぶようになって2年半近くが過ぎた。

仕事においては，業務システムでは相変わらず Java への replace 仕事ばっかりだし（私のようなロートルエンジニアは過去の[技術的負債](https://ja.wikipedia.org/wiki/%E6%8A%80%E8%A1%93%E7%9A%84%E8%B2%A0%E5%82%B5)の後始末をするのがお役目），組み込みでは C/C++ がメインなので， [Go 言語]を使う機会がないのだが，恐ろしいことに私の中で [Go 言語]が「[母国語]」になりつつある。
つまり，あるロジックをプログラム・コードに「翻訳」する際に，まず [Go 言語]のコードが思い浮かぶようになってきた。

この業界に四半世紀以上も足を突っ込んでるが脳内の[母国語]が変わるという経験は初めてで，まるで転生物のラノベ作品を読むがごとく，年甲斐もなく「[わーい！ たのしー！](https://nijipi.com/it-news/kemono-lang_ruby-brainfuck/)」な気分でコードを眺める日々である。

もっとも，有り余る計算資源を持つクラウド環境ならともかくリソースの限られた RTOS (Real-Time Operating System) 環境下では息を吸うようにヒープを使いまくる [Go 言語]での実装は向いてない気がするので，「これは言語のチョイスを間違えたかなぁ」とも思わないでもない。
まぁでもそれならそれで C/C++ を使えばいいので困ることでもないんだけどね。

でも  [Go 言語]が[母国語]になると（アセンブラに近い C 言語はともかく） C++ って本当に面倒くさい言語だったんだなぁ，と涙が出ちゃう。
だってエンジニアだもん。

## ブックマーク

- [Go言語がダメな理由 | プログラミング | POSTD](http://postd.cc/why-go-is-not-good/)
- [Go言語のリアルタイムGC　理論と実践 | プログラミング | POSTD](http://postd.cc/golangs-real-time-gc-in-theory-and-practice/)

- [プログラミング言語 Go](/golang/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[母国語]: {{< ref "/remark/2015/programming-language.md" >}} "プログラミング言語との付き合い方"
[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"
[`sort`]: https://golang.org/pkg/sort/ "sort - The Go Programming Language"
[slice]: http://golang.org/ref/spec#Slice_types

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41W69WGATNL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4320026926/baldandersinf-22/">プログラミング言語C 第2版 ANSI規格準拠</a></dt><dd>B.W. カーニハン D.M. リッチー 石田 晴久 </dd><dd>共立出版 1989-06-15</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320027485/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320027485.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Cアンサー・ブック 第2版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4874084141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4874084141.09._SCTHUMBZZZ_.jpg"  alt="C言語による最新アルゴリズム事典 (ソフトウェアテクノロジー)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774111422/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774111422.09._SCTHUMBZZZ_.jpg"  alt="C言語ポインタ完全制覇 (標準プログラマーズライブラリ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797304952/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797304952.09._SCTHUMBZZZ_.jpg"  alt="定本 Cプログラマのためのアルゴリズムとデータ構造 (SOFTBANK BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4900900648/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4900900648.09._SCTHUMBZZZ_.jpg"  alt="C実践プログラミング 第3版"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4781908535/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4781908535.09._SCTHUMBZZZ_.jpg"  alt="工科系の数学 (5)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4781908896/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4781908896.09._SCTHUMBZZZ_.jpg"  alt="工科系の数学〈6〉関数論"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4756136494/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4756136494.09._SCTHUMBZZZ_.jpg"  alt="プログラミング作法"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798030147/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798030147.09._SCTHUMBZZZ_.jpg"  alt="苦しんで覚えるC言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798101036/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798101036.09._SCTHUMBZZZ_.jpg"  alt="Cの絵本―C言語が好きになる9つの扉"  /></a> </p>
<p class="description">通称 “K&amp;R”。その筋の人々には「バイブル」と呼ばれる名著。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-02-19">2017-02-19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
