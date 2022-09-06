+++
title = "Go 1.19.1 のリリース【セキュリティ・アップデート】"
date =  "2022-09-07T08:08:38+09:00"
description = "今回は2件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/rlvRpp6WVVY "[security] Go 1.19.1 and Go 1.18.6 pre-announcement")通り [Go] 1.19.1 がリリースされた。

- [[security] Go 1.19.1 and Go 1.18.6 are released](https://groups.google.com/g/golang-announce/c/x49AQzIVX-s)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://go.dev/doc/devel/release#go1.19.minor" lang="en" >}}
{{% quote %}}go1.19.1 (released 2022-09-06) includes security fixes to the `net/http` and `net/url` packages, as well as bug fixes to the compiler, the go command, the pprof command, the linker, the runtime, and the `crypto/tls` and `crypto/x509` packages. See the [Go 1.19.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.1+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2022-27664] net/http: handle server errors after sending GOAWAY

{{< fig-quote type="markdown" title="Go 1.19.1 and Go 1.18.6 are released" link="https://groups.google.com/g/golang-announce/c/x49AQzIVX-s" lang="en" >}}
{{% quote %}}A closing HTTP/2 server connection could hang forever waiting for a clean shutdown that was preempted by a subsequent fatal error. This failure mode could be exploited to cause a denial of service{{% /quote %}}.
{{< /fig-quote >}}

（以下未稿）

## [CVE-2022-32190] net/url: JoinPath does not strip relative path components in all circumstances

{{< fig-quote type="markdown" title="Go 1.19.1 and Go 1.18.6 are released" link="https://groups.google.com/g/golang-announce/c/x49AQzIVX-s" lang="en" >}}
{{% quote %}}`JoinPath` and `URL.JoinPath` would not remove `../` path components appended to a relative path. For example, `JoinPath("https://go.dev", "../go")` returned the URL `https://go.dev/../go`, despite the `JoinPath` documentation stating that `../` path elements are cleaned from the result{{% /quote %}}.
{{< /fig-quote >}}

えっ，これヤバくね？

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.19.1.linux-amd64.tar.gz`](https://go.dev/dl/go1.19.1.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.19.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.19.1.linux-amd64.tar.gz
$ sudo mv go go1.19.1
$ sudo ln -s go1.19.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.19.1 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-27664]: https://nvd.nist.gov/vuln/detail/CVE-2022-27664
[CVE-2022-32190]: https://nvd.nist.gov/vuln/detail/CVE-2022-32190

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
