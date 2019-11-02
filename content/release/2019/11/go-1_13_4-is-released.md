+++
title = "Go 1.13.4 がリリースされた"
date =  "2019-11-02T11:50:28+09:00"
description = "今回はセキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.13.4 がリリースされた。

- [Go 1.13.4 and Go 1.12.13 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/YVXawNxmEBw)

今回はセキュリティ・アップデートはなし。

{{< fig-quote type="md" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.13.minor" lang="en" >}}
{{< quote >}}go1.13.4 (released 2019/10/31) includes fixes to the `net/http` and `syscall` packages. It also fixes an issue on macOS 10.15 Catalina where the non-notarized installer and binaries were being [rejected by Gatekeeper](https://golang.org/issue/34986). See the [Go 1.13.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.4) on our issue tracker for details{{< /quote >}}.
{{< /fig-quote >}}

[Ubuntu] 19.10 からは 1.12 をサポートし始めたようだが（セキュリティ・アップデートを含め）どこまで追従しているか分からない（なんであんな分かりにくいバージョニングをするのかねぇ， Linux は）。
なので[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.13.4.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.4.linux-amd64.tar.gz
$ sudo mv go go1.13.4
$ sudo ln -s go1.13.4 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.4 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

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
