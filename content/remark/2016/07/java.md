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

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4774182427/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4774182427.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.93"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> </p>
<p class="description">腹を括って発注かけました。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-09">2016-07-09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
