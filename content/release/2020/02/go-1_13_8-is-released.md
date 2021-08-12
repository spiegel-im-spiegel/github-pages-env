+++
title = "Go 1.13.8 がリリースされた"
date =  "2020-02-13T22:42:59+09:00"
description = "今回はセキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.13.8 がリリースされた。

- [Go 1.13.8 and Go 1.12.17 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/pYiaiBU6kvs)

今回はセキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.13.minor" lang="en" >}}
{{% quote %}}go1.13.8 (released 2020/02/12) includes fixes to the runtime, the `crypto/x509`, and `net/http` packages. See the [Go 1.13.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.8+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.13.8.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.8.linux-amd64.tar.gz
$ sudo mv go go1.13.8
$ sudo ln -s go1.13.8 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.8 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
