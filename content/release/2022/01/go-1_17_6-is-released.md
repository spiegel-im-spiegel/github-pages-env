+++
title = "Go 1.17.6 のリリース"
date =  "2022-01-07T20:09:23+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.6 がリリースされた。

- [Go 1.17.6 and Go 1.16.13 are released](https://groups.google.com/g/golang-announce/c/95ZD3rKn4DI)

セキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.6 (released 2022-01-06) includes fixes to the compiler, linker, runtime, and the `crypto/x509`, `net/http`, and `reflect` packages. See the [Go 1.17.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.6+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.17.6.linux-amd64.tar.gz`](https://go.dev/dl/go1.17.6.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.17.6.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.6.linux-amd64.tar.gz
$ sudo mv go go1.17.6
$ sudo ln -s go1.17.6 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.6 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
