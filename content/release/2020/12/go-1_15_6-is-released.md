+++
title = "Go 1.15.6 がリリースされた"
date =  "2020-12-05T10:01:58+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.15.6 がリリースされた。

- [Go 1.15.6 and Go 1.14.13 are released](https://groups.google.com/g/golang-announce/c/0kN8l6MprTo)

セキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.15.minor" lang="en" >}}
{{% quote %}}go1.15.6 (released 2020/12/03) includes fixes to the compiler, linker, runtime, the go command, and the `io` package. See the [Go 1.15.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.6+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.15.6.linux-amd64.tar.gz`](https://golang.org/dl/go1.15.6.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.15.6.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.15.6.linux-amd64.tar.gz
$ sudo mv go go1.15.6
$ sudo ln -s go1.15.6 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.15.6 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
