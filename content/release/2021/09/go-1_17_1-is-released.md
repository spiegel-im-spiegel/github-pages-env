+++
title = "Go 1.17.1 のリリース【セキュリティ・アップデート】"
date =  "2021-09-10T19:58:44+09:00"
description = "Go 1.16.5 の修正が直りきってなかったようだ。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.1 がリリースされた。

- [[security] Go 1.17.1 and Go 1.16.8 are released](https://groups.google.com/g/golang-announce/c/dx9d7IOseHw)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.1 (released 2021-09-09) includes a security fix to the `archive/zip` package, as well as bug fixes to the compiler, linker, the go command, and to the `crypto/rand`, `embed`, `go/types`, `html/template`, and `net/http` packages. See the [Go 1.17.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.1+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-39293]

{{< fig-quote type="markdown" title="Go 1.17.1 and Go 1.16.8 are released" link="https://groups.google.com/g/golang-announce/c/dx9d7IOseHw" lang="en" >}}
{{% quote %}}The fix for CVE-2021-33196 can be bypassed by crafted inputs. As a result, the `NewReader` and `OpenReader` functions in [`archive/zip`](https://pkg.go.dev/archive/zip) can still cause a panic or an unrecoverable fatal error when reading an archive that claims to contain a large number of files, regardless of its actual size{{% /quote %}}.
{{< /fig-quote >}}

[Go 1.16.5 の修正]({{< ref "/release/2021/06/go-1_16_5-is-released.md" >}})が直りきってなかったようだ。
[CVE-2021-33196] は CVSSv3.1 基本評価値が 7.5 とヤバめの値だったので，今回も早めに対処するのがいいだろう。

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.17.1.linux-amd64.tar.gz`](https://golang.org/dl/go1.17.1.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.17.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.1.linux-amd64.tar.gz
$ sudo mv go go1.17.1
$ sudo ln -s go1.17.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.1 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-33196]: https://nvd.nist.gov/vuln/detail/CVE-2021-33196
[CVE-2021-39293]: https://nvd.nist.gov/vuln/detail/CVE-2021-39293

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
