+++
title = "Go 1.10.3 および 1.9.7 がリリース"
date = "2018-06-24T16:55:53+09:00"
description = "TLS および X.509 関連の改修と vgo 以降のための前準備？ かな。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "golang", "engineering", "versioning" ]

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

3週間前の話でゴメンペコン。

[Go 言語]コンパイラの 1.10.3 と 1.9.7 がリリースされている。

- [Go 1.10.3 and Go 1.9.7 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/_S9YQriFKuU)
    - [Go 1.10.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.10.3)

{{< fig-quote title="Go 1.10.3 and Go 1.9.7 are released" link="https://groups.google.com/forum/#!topic/golang-announce/_S9YQriFKuU" lang="en" >}}
<q>These releases include fixes to the go command, and to the crypto/tls, crypto/x509, and strings packages. In particular, they add minimal support to the go command for the vgo transition.</q>
{{< /fig-quote >}}

ちうわけで Web 関連や暗号関連の製品を手がけている人はご注意を。
vgo サポートについての詳細は以下を参考にするといいだろう。

- [cmd/go: add minimal module-awareness for legacy operation](https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9)

なお vgo の考え方については以下を参考にどうぞ。

- [vgo (Versioned Go) に関する覚え書き]({{< ref "/golang/go-and-versioning.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"

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
