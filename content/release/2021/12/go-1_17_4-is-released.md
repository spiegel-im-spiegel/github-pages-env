+++
title = "Go 1.17.4 のリリースと ssh パッケージのセキュリティ・アップデート"
date =  "2021-12-04T12:02:12+09:00"
description = "Go のサイトが golang.org から go.dev に完全移行したみたい"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability", "openssh" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.4 がリリースされた。

- [Go 1.17.4 and Go 1.16.11 are released](https://groups.google.com/g/golang-announce/c/Xcefh1Tfj-U)

セキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.4 (released 2021-12-02) includes fixes to the compiler, linker, runtime, and the `go/types`, `net/http`, and `time` packages. See the [Go 1.17.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.4+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.17.4.linux-amd64.tar.gz`](https://go.dev/dl/go1.17.4.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.17.4.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.4.linux-amd64.tar.gz
$ sudo mv go go1.17.4
$ sudo ln -s go1.17.4 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.4 linux/amd64
```

そうそう。
[Go] のサイトが `golang.org` から [`go.dev`][Go] に完全移行したみたい。
今のところは旧サイトから適切にリダイレクトしてくれるけど，ブログのリンクとかボチボチ直さんといかんかねぇ。

## ssh パッケージのセキュリティ・アップデート

[Go] コンパイラのアップデートに先立ち [`golang.org/x/crypto/ssh`][golang.org/x/crypto/ssh] パッケージについてセキュリティ・アップデートのアナウンスが出ている。

- [[security] Vulnerability in golang.org/x/crypto/ssh](https://groups.google.com/g/golang-announce/c/2AR1sKiM-Qs/m/9LAF9FxvBwAJ)

### [CVE-2021-43565]

あー。
panic で落ちちゃうのか。
そりゃアカンね。

{{< fig-quote type="markdown" title="Vulnerability in golang.org/x/crypto/ssh" link="https://groups.google.com/g/golang-announce/c/2AR1sKiM-Qs/m/9LAF9FxvBwAJ" lang="en" >}}
{{% quote %}}Version v0.0.0-20211202192323-5770296d904e of [`golang.org/x/crypto`](http://golang.org/x/crypto) fixes a vulnerability in the [`golang.org/x/crypto/ssh`](http://golang.org/x/crypto/ssh) package which allowed unauthenticated clients to cause a panic in SSH servers{{% /quote %}}.
{{< /fig-quote >}}

（以下未稿）

## アップデートは...

アップデートは計画的に。

[Go]: https://go.dev/
[golang.org/x/crypto/ssh]: https://pkg.go.dev/golang.org/x/crypto/ssh "ssh package - golang.org/x/crypto/ssh - pkg.go.dev"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-43565]: https://nvd.nist.gov/vuln/detail/CVE-2021-43565

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
