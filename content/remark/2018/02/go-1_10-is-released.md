+++
title = "Go 1.10 がリリース"
date = "2018-02-17T13:29:57+09:00"
description = "古いバージョンには脆弱性を含むものがあるので，この機会にとっとと 1.10 にアップデートしましょう。"
image = "/images/attention/kitten.jpg"
tags        = [ "golang" ]

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
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週の予告通り [Go 言語]コンパイラの 1.10 がリリースされた。

- [Go 1.10 is released - The Go Blog](https://blog.golang.org/go1.10)
- [Go 1.10 Release Notes - The Go Programming Language](https://golang.org/doc/go1.10)

今回も修正・変更は多岐にわたるが，とりあえず [`strings`]`.Buffer` は気になるところである。
おそらく [`bytes`]`.Buffer` で実装している部分の幾つかを置き換えれるんじゃないかと思うのだが，私が利用している CI サービスの [Travis CI] が現時点（2018-02-17）で 1.10 に対応してないのよ。

まぁ，おいおい試していくか。

そうそう。
1.8.6 およびそれ以前， 1.9.3 およびそれ以前，および 1.10rc1 には [`#cgo` ディレクトリに絡むかなり大きめの脆弱性]({{< ref "/remark/2018/02/go-1_8_7-and-1_9_4-are-released.md" >}} "Go 1.8.7 および 1.9.4 がリリース（セキュリティ・アップデート）")が存在するので，この機会にとっとと 1.10 にアップデートしましょう。

## ブックマーク

- [go1.10beta1の標準パッケの大きな変更点を確認しておく。 - Qiita](https://qiita.com/A_Resas/items/59bf6cda976e29751890)
- [Go1.10で入るstrings.Builderを検証した #golang - Qiita](https://qiita.com/tenntenn/items/94923a0c527d499db5b9)
- [go1.10のtest2jsonを試してみた - Qiita](https://qiita.com/dproject21/items/c406b0044280508b41ff)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`bytes`]: https://golang.org/pkg/bytes/ "bytes - The Go Programming Language"
[`strings`]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
