+++
title = "Go 1.11.3 のリリース 【セキュリティ・アップデート】"
date = "2018-12-14T22:10:04+09:00"
update = "2018-12-15T16:00:33+09:00"
description = "3つのインシデントに対して修正が行われている。深刻度が高そう。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "golang", "security", "vulnerability" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]コンパイラの 1.11.3 および 1.10.6 がリリースされた。
今回はセキュリティ・アップデートを含んでいるので必ずアップデートすること。

- [[security] Go 1.11.3 and Go 1.10.6 are released - Google group](https://groups.google.com/forum/#%21topic/golang-announce/Kw31K8G7Fi0)
- [Go 1.11.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.11.3)

以下の3つのインシデントに対して修正が行われている。

{{% fig-quote title="Go 1.11.3 and Go 1.10.6 are released" link="https://groups.google.com/forum/#%21topic/golang-announce/Kw31K8G7Fi0" %}}
- cmd/go: remote command execution during "`go get -u`" : The issue is CVE-2018-16873 and Go issue [golang.org/issue/29230](https://golang.org/issue/29230). See the Go issue for details. Thanks to Etienne Stalmans from the Heroku platform security team for discovering and reporting this issue.
- cmd/go: directory traversal in "`go get`" via curly braces in import paths : The issue is CVE-2018-16874 and Go issue [golang.org/issue/29231](https://golang.org/issue/29231). See the Go issue for details. Thanks to ztz of Tencent Security Platform for discovering and reporting this issue.
- crypto/x509: CPU denial of service in chain validation : The issue is CVE-2018-16875 and Go issue [golang.org/issue/29233](https://golang.org/issue/29233). See the Go issue for details. Thanks to Netflix for discovering and reporting this issue.
{{% /fig-quote %}}

- [CVE-2018-16873](https://nvd.nist.gov/vuln/detail/CVE-2018-16873)
- [CVE-2018-16874](https://nvd.nist.gov/vuln/detail/CVE-2018-16874)
- [CVE-2018-16875](https://nvd.nist.gov/vuln/detail/CVE-2018-16875)

またこれとは別に “`go get github.com/golang/pkg/...`” でコードを再帰的にダウンロードしてしまう問題があるそうで，これは次回の 1.11.4 および Go 1.10.7 で解消されるとのこと。

さらに詳しい情報が分かればこの記事に加筆・修正を行う。

アップデートは計画的に。

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" height="160" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"></a></div>
	<dl class="fn">
      <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
      <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
      <dd>丸善出版</dd>
	  <dd>評価&nbsp;<abbr class="rating fa-sm" title="5">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
      </abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.19</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
