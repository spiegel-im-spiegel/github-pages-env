+++
title = "Go 1.14.2 がリリースされた"
date =  "2020-04-09T10:34:18+09:00"
description = "セキュリティ・アップデートはなし，でいいかな。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.14.2 がリリースされた。

- [Go 1.14.2 and Go 1.13.10 are released - Google group](https://groups.google.com/forum/#!topic/golang-announce/9UJN3gwMzhY)

セキュリティ・アップデートはなし，でいいかな。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.14.minor" lang="en" >}}
{{% quote %}}go1.14.2 (released 2020/04/08) includes fixes to cgo, the go command, the runtime, `os/exec`, and testing packages. See the [Go 1.14.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.2+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.14.2.linux-amd64.tar.gz`](https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.14.2.linux-amd64.tar.gz
$ sudo mv go go1.14.2
$ sudo ln -s go1.14.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.14.2 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
