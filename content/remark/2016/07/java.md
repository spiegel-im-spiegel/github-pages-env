+++
date = "2016-07-09T15:27:27+09:00"
update = "2016-11-08T20:52:07+09:00"
description = "現在 Java でバリバリ仕事しているおぢさんから（老婆心ながら）言っておくと「Java はやめておけ」である。"
draft = false
tags = ["programming", "language", "java"]
title = "Java はやめておけ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

- [ITエンジニアがこれから重要になると思うプログラミング言語TOP10｜ニフティニュース](https://news.nifty.com/article/item/neta/dime-269849/)

うわぁ。
まじすか，これ。
確かに [TIOBE](http://www.tiobe.com/tiobe_index?page=index) でも Java は常に1位か2位だけどね。

現在 Java でバリバリ仕事しているおぢさんから（老婆心ながら）言っておくと「Java はやめておけ」である。

- [Java EEの開発が停滞？--オラクルの不透明な態度にコミュニティーが不信感 - ZDNet Japan](http://japan.zdnet.com/article/35085410/)
- [オラクル、次期Java EEはマイクロサービスやコンテナをサポートするものになるとコメント。9月のJavaOneで計画を発表予定。ただし本気度には疑問符も － Publickey](http://www.publickey1.jp/blog/16/java_ee9javaone.html)

どういうわけか日本人は Java が大好きで，確かにあと数年は飯の種になるだろうが，将来を見据えれば Java は間違いなく COBOL と同じ道をたどる。
今 COBOL-er は季節労働者のごとく仕事を求めて全国行脚しているそうだ。
以前仕事で一緒になった COBOL プログラマ（私と同年代）は，その前は富山で仕事をしていたと言っていた。

今後10年20年をにらんで「飯の種」としてプログラミング言語をきちんと学びたいなら「関数型プログラミング言語」を選択すべき[^a]。
個人的には（純正の関数型言語ではないが） Python か ES6 以降の JavaScript をお薦めする。
あっ Swift も多分オススメ。

[^a]: 今回紹介した以外にも Erlang や Haskell や Scala といった言語も気になっているのだが，私が評価できてないので割愛する（Haskell については[この辺の記事](http://postd.cc/becoming-productive-in-haskell/ "Haskellで生産性を高める-Pythonからの移行 | プログラミング | POSTD")が参考になるかも）。 Scala は飯の種になりそうな気がして本も買ってるのだが... Python や Swift のような「いまどき」の言語は multiparadigm programming language と呼ぶそうで，古い資産との整合性を取りやすいのが利点である（Scala なんかモロにそうだよね）。見方を変えるなら，今はもっと大きなパラダイム・シフトの真っ最中であると言うこともできる。故に本当に10年後を見据えた場合に，どれが主流になるかなんて誰も予測できないのではないだろうか。たとえばコーディングを AI がやるようになれば，今ある言語は絶滅し， AI が記述しやすい（かつ人間も読みやすい）言語体系が主流になる可能性だってあるのだ。故に「どの言語か？」という議論はあまり意味がなくて，泡沫のように現れては消える技術トレンドをキチンと押さえておくことが重要である。人間死ぬまで勉強ですよ。

日本では何故か Python の人気が薄い気がするが，新しいアイデアは大抵の場合，真っ先に Python で実装される。
Python 自体は飯の種にならなくても技術トレンドを追いかけるなら外せない。

JavaScript はもともと手続き型言語だが， ES6 で関数型の言語仕様がいろいろと追加されている。
なにより Web システムにおいてサーバ側からクライアント側まで同じ言語で記述できるのは有利な点である。

- [関数型プログラミングはまず考え方から理解しよう - Qiita](http://qiita.com/stkdev/items/5c021d4e5d54d56b927c)
- [for文を使わないプログラミングって？（関数型プログラミング入門） - Qiita](http://qiita.com/srd7/items/fad2d0a94b99d1de2e48)

Swift はオープンソースになってサーバ用途でも使えるようになった。
個人的には色々試してみたい。

- [IBM、ヴイエムウェアとハイブリッドクラウドで提携--Swiftのクラウド対応も発表 - ZDNet Japan](http://japan.zdnet.com/article/35078299/)
- [Swiftで自然数を作ってみた（ペアノの公理） - Qiita](http://qiita.com/taketo1024/items/2ab856d21bf9b9f30357)
- [Swiftで代数学入門 〜 1. 数とは何か？ - Qiita](http://qiita.com/taketo1024/items/bd356c59dc0559ee9a0b)

ハードウェア寄りの仕事に興味があるなら C/C++ のような手続き型の言語もいまだに有効である。
たとえば Google の [TensorFlow](https://www.tensorflow.org/ "TensorFlow — an Open Source Software Library for Machine Intelligence") のバックエンドは C++ で実装されているそうな。

- [TensorFlow に関するブックマーク]({{< ref "/remark/2016/02/tensorflow.md" >}})

あと，ものすごく個人的な意見として，手続き型言語なら [Go 言語]をお薦めする（笑）

- [プログラミング言語 Go](/golang/)

基礎をきっちり固めてから「第2言語」や「第3言語」として Java や .NET を学ぶのは悪くない。
将来的にはともかく，現時点で Java は「飯の種」になっているので。
でも，上で挙げたような言語を学んだあとで Java をやったら，間違いなくモニタに向かって中指をおっ立てるハメになるであろう（笑）

最後に。

反論は自由ですが，私は関知しません。
言語論争は宗教論争と同じで出口がないからね。

[Go 言語]: https://golang.org/ "The Go Programming Language"

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
