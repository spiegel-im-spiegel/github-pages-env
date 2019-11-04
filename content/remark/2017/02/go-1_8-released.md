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
  url = "https://baldanders.info/profile/"
  name = "Spiegel"
  flattr = ""
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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EC-%E7%AC%AC2%E7%89%88-ANSI%E8%A6%8F%E6%A0%BC%E6%BA%96%E6%8B%A0-B-W-%E3%82%AB%E3%83%BC%E3%83%8B%E3%83%8F%E3%83%B3/dp/4320026926?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4320026926"><img src="https://images-fe.ssl-images-amazon.com/images/I/41W69WGATNL._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EC-%E7%AC%AC2%E7%89%88-ANSI%E8%A6%8F%E6%A0%BC%E6%BA%96%E6%8B%A0-B-W-%E3%82%AB%E3%83%BC%E3%83%8B%E3%83%8F%E3%83%B3/dp/4320026926?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4320026926">プログラミング言語C 第2版 ANSI規格準拠</a></dt>
	<dd>B.W. カーニハン, D.M. リッチー</dd>
	<dd>石田 晴久 (翻訳)</dd>
    <dd>共立出版 1989-06-15</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4320026926, EAN: 9784320026926</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">通称 “K&amp;R”。その筋の人々には「バイブル」と呼ばれる名著（当時は）。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-07">2018-12-07</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
