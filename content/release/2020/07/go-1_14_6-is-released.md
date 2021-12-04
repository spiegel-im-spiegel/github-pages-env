+++
title = "Go 1.14.6 がリリースされた"
date =  "2020-07-18T08:57:01+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.14.6 がリリースされた。

- [Go 1.14.6 and Go 1.13.14 are released](https://groups.google.com/g/golang-announce/c/9KNZ5g9tv4o)

セキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.14.minor" lang="en" >}}
{{% quote %}}go1.14.6 (released 2020/07/16) includes fixes to the go command, the compiler, the linker, vet, and the `database/sql`, `encoding/json`, `net/http`, `reflect`, and `testing` packages. See the [Go 1.14.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.6+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.14.6.linux-amd64.tar.gz`](https://golang.org/dl/go1.14.6.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.14.6.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.14.6.linux-amd64.tar.gz
$ sudo mv go go1.14.6
$ sudo ln -s go1.14.6 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.14.6 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
