+++
title = "Go 1.10 がリリース"
date = "2018-02-17T13:29:57+09:00"
description = "古いバージョンには脆弱性を含むものがあるので，この機会にとっとと 1.10 にアップデートしましょう。"
image = "/images/attention/kitten.jpg"
tags        = [ "golang" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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
