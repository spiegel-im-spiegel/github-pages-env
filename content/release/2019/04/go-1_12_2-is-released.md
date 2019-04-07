+++
title = "Go 1.12.2 がリリースされた"
date =  "2019-04-07T13:47:19+09:00"
description = "コンパイラ本体は通常のアップデート。外部パッケージで脆弱性の情報あり。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "security", "vulnerability" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 1.12.2 がリリースされた。
セキュリティ・アップデートはなし。

- [Go 1.12.2 and Go 1.11.7 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/z9eTD34GEIs)

{{< fig-quote title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.12.minor" lang="en" >}}
<q>go1.12.2 (released 2019/04/05) includes fixes to the compiler, the go command, the runtime, and the <code>doc</code>, <code>net</code>, <code>net/http/httputil</code>, and <code>os</code> packages. See the <a href="https://github.com/golang/go/issues?q=milestone%3AGo1.12.2">Go 1.12.2 milestone</a> on our issue tracker for details.</q>
{{< /fig-quote >}}

また，このリリースより少し前に [`golang.org/x/crypto/salsa20`](https://golang.org/x/crypto/salsa20) パッケージに関する脆弱性情報がアナウンスされている。

- [[security] Vulnerability in golang.org/x/crypto/salsa20 - Google Group](https://groups.google.com/forum/#!topic/golang-announce/tjyNcJxb2vQ)

{{< fig-quote title="Vulnerability in golang.org/x/crypto/salsa20" link="https://groups.google.com/forum/#!topic/golang-announce/tjyNcJxb2vQ" lang="en" >}}
<q>If more than 256 GiB of keystream is generated, or if the counter otherwise grows greater than 32 bits, the amd64 implementation will first generate incorrect output, and then cycle back to previously generated keystream. Repeated keystream bytes can lead to loss of confidentiality in encryption applications, or to predictability in CSPRNG applications.</q>
{{< /fig-quote >}}

けっこうヤバい脆弱性なので，ご利用の方はアップデートを。
それ以外にも（メール本文にも目を通して）必要な措置があれば行うこと。

アップデートは計画的に。

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
