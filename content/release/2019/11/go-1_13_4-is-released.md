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

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.13.minor" lang="en" >}}
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

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
