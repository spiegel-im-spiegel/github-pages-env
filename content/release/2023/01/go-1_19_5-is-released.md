+++
title = "Go 1.19.5 のリリース"
date =  "2023-01-13T09:07:56+09:00"
description = "セキュリティ・アップデートはなし。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.19.5 がリリースされた。

- [Go 1.19.5 and Go 1.18.10 are released](https://groups.google.com/g/golang-announce/c/wrmRT7JlevE)

今回はセキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://go.dev/doc/devel/release#go1.19" lang="en" >}}
go1.19.5 (released 2023-01-10) includes fixes to the compiler, the linker, and the `crypto/x509`, `net/http`, `sync/atomic`, and `syscall` packages. See the [Go 1.19.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.5+label%3ACherryPickApproved) on our issue tracker for details.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.19.5.linux-amd64.tar.gz`](https://go.dev/dl/go1.19.5.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.19.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.19.5.linux-amd64.tar.gz
$ sudo mv go go1.19.5
$ sudo ln -s go1.19.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.19.5 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-41720]: https://nvd.nist.gov/vuln/detail/CVE-2022-41720
[CVE-2022-41717]: https://nvd.nist.gov/vuln/detail/CVE-2022-41717

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
