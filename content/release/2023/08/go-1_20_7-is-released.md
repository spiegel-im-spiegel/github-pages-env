+++
title = "Go 1.20.7 のリリース【セキュリティ・アップデート】"
date =  "2023-08-02T08:18:19+09:00"
description = "CVE ID ベースで合計3件の脆弱性の修正がある。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/7b0c3Z5Ko8g "[security] Go 1.20.7 and Go 1.19.12 pre-announcement")どおり [Go] 1.20.7 がリリースされた。

- [[security] Go 1.20.7 and Go 1.19.12 are released](https://groups.google.com/g/golang-announce/c/X0b6CsSAaYI)

今回は CVE ID ベースで1件の脆弱性修正を含んでいる。

## [CVE-2023-29409] crypto/tls: restrict RSA keys in certificates to <= 8192 bits

{{< fig-quote type="markdown" title="Go 1.20.7 and Go 1.19.12 are released" link="https://groups.google.com/g/golang-announce/c/X0b6CsSAaYI" lang="en" >}}
Extremely large RSA keys in certificate chains can cause a client/server to expend significant CPU time verifying signatures. Limit this by restricting the size of RSA keys transmitted during handshakes to <= 8192 bits.

Based on a survey of publicly trusted RSA keys, there are currently only three certificates in circulation with keys larger than this, and all three appear to be test certificates that are not actively deployed. It is possible there are larger keys in use in private PKIs, but we target the web PKI, so causing breakage here in the interests of increasing the default safety of users of `crypto/tls` seems reasonable.
{{< /fig-quote >}}

どうやら 8192 bit より大きいサイズの RSA 鍵は取り扱わないことにしたようだ。
まぁ，妥当かな。

（以下未稿）

## その他のパッケージの更新

[Go] 標準パッケージ以外にも以下のパッケージでセキュリティ・アップデートが行われている。

- [[security] Vulnerabilities in golang.org/x/net/html and golang.org/x/image/tiff](https://groups.google.com/g/golang-announce/c/B15Lav568-o)

### [CVE-2023-3978] Vulnerabilities in golang.org/x/net/html

{{< fig-quote type="markdown" title="Vulnerabilities in golang.org/x/net/html and golang.org/x/image/tiff" link="https://groups.google.com/g/golang-announce/c/B15Lav568-o" lang="en" >}}
Version v0.13.0 of [`golang.org/x/net`](http://golang.org/x/net) fixes a vulnerability in the [`golang.org/x/net/html`](http://golang.org/x/net/html) package which caused a mismatch between parsing and rendering.<br>
Text nodes not in the HTML namespace were being incorrectly literally rendered, causing text which should've been escaped to not be. This could lead to an XSS attack.{{< /fig-quote >}}

（以下未稿）

### [CVE-2023-29407] Vulnerabilities in golang.org/x/image/tiff

{{< fig-quote type="markdown" title="Vulnerabilities in golang.org/x/net/html and golang.org/x/image/tiff" link="https://groups.google.com/g/golang-announce/c/B15Lav568-o" lang="en" >}}
Version v0.10.0 of [`golang.org/x/image`](http://golang.org/x/image) fixes two vulnerabilities in the [`golang.org/x/image/tiff`](http://golang.org/x/image/tiff) package which can cause excessive CPU or memory consumption when parsing images.
{{< /fig-quote >}}

また TIFF か...

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.7.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.7.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.7.linux-amd64.tar.gz
$ sudo mv go go1.20.7
$ sudo ln -s go1.20.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.7 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.20.7@latest
$ go1.20.7 download
$ go1.20.7 version
go version go1.20.7 linux/amd64
```

てな感じで導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-29409]: https://nvd.nist.gov/vuln/detail/CVE-2023-29409
[CVE-2023-3978]: https://nvd.nist.gov/vuln/detail/CVE-2023-3978
[CVE-2023-29407]: https://nvd.nist.gov/vuln/detail/CVE-2023-29407

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
