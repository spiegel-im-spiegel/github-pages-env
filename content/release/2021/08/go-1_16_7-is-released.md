+++
title = "Go 1.16.7 のリリース【セキュリティ・アップデート】"
date =  "2021-08-07T09:59:53+09:00"
description = "競合状態でちゃんと panic になるだけマシ（笑）"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.16.7 がリリースされた。

- [Go 1.16.7 and Go 1.15.15 are released](https://groups.google.com/g/golang-announce/c/uHACNfXAZqk)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.16.minor" lang="en" >}}
{{% quote %}}go1.16.7 (released 2021-08-05) includes a security fix to the `net/http/httputil` package, as well as bug fixes to the compiler, the linker, the runtime, the go command, and the `net/http` package. See the [Go 1.16.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.7+label%3ACherryPickApproved) on our issue tracker for details. {{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-36221]

{{< fig-quote type="markdown" title="Go 1.16.7 and Go 1.15.15 are released" link="https://groups.google.com/g/golang-announce/c/uHACNfXAZqk" lang="en" >}}
{{% quote %}}A `net/http/httputil` `ReverseProxy` can panic due to a race condition if its Handler aborts with `ErrAbortHandler`, for example due to an error in copying the response body. An attacker might be able to force the conditions leading to the race condition{{% /quote %}}.
{{< /fig-quote >}}

競合状態でちゃんと panic になるだけマシ，と思うことにしよう。

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.16.7.linux-amd64.tar.gz`](https://golang.org/dl/go1.16.7.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.16.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.16.7.linux-amd64.tar.gz
$ sudo mv go go1.16.7
$ sudo ln -s go1.16.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.16.7 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-36221]: https://nvd.nist.gov/vuln/detail/CVE-2021-36221

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
