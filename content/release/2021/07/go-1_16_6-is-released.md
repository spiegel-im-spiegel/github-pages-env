+++
title = "Go 1.16.6 のリリース【セキュリティ・アップデート】"
date =  "2021-07-13T20:15:01+09:00"
description = "今回は HTTPS クライアントに影響するため Web API クライアントを構成しているツール等はご注意を。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/JvWG9FUUYT0 "[security] Go 1.16.6 and Go 1.15.14 pre-announcement")通り [Go] 1.16.6 がリリースされた。

- [[security] Go 1.16.6 and Go 1.15.14 are released](https://groups.google.com/g/golang-announce/c/n9FxMelZGAQ/m/4ZhvTx0dAQAJ)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.16.minor" lang="en" >}}
{{% quote %}}go1.16.6 (released 2021-07-12) includes a security fix to the `crypto/tls` package, as well as bug fixes to the compiler, and the `net` and `net/http` packages. See the [Go 1.16.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.6+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-34558]

{{< fig-quote type="markdown" title="Go 1.16.6 and Go 1.15.14 are released" link="https://groups.google.com/g/golang-announce/c/n9FxMelZGAQ/m/4ZhvTx0dAQAJ" lang="en" >}}
{{% quote %}}`crypto/tls` clients can panic when provided a certificate of the wrong type for the negotiated parameters.  `net/http` clients performing HTTPS requests are also affected. The panic can be triggered by an attacker in a privileged network position without access to the server certificate's private key, as long as a trusted ECDSA or Ed25519 certificate for the server exists (or can be issued), or the client is configured with `Config.InsecureSkipVerify`. Clients that disable all TLS_RSA cipher suites (that is, TLS 1.0–1.2 cipher suites without ECDHE), as well as TLS 1.3-only clients, are unaffected{{% /quote %}}.
{{< /fig-quote >}}

今回は HTTPS クライアントに影響するため Web API クライアントを構成しているツール等はご注意を。

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.16.6.linux-amd64.tar.gz`](https://golang.org/dl/go1.16.6.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.16.6.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.16.6.linux-amd64.tar.gz
$ sudo mv go go1.16.6
$ sudo ln -s go1.16.6 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.16.6 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-34558]: https://nvd.nist.gov/vuln/detail/CVE-2021-34558

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
