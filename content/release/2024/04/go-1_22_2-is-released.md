+++
title = "Go 1.22.2 のリリース【セキュリティ・アップデート】"
date =  "2024-04-06T09:24:31+09:00"
description = "CVE ID ベースで1件の脆弱性の修正がある。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/nUtcDA04HLM "[security] Go 1.22.2 and Go 1.21.9 pre-announcement")どおり [Go] 1.22.2 がリリースされた。

- [[security] Go 1.22.2 and Go 1.21.9 are released](https://groups.google.com/g/golang-announce/c/YgW0sx8mN3M)

今回は CVE ID ベースで1件の脆弱性修正を含んでいる。

## [CVE-2023-45288] http2: close connections when receiving too many headers

{{< fig-quote type="markdown" title="Go 1.22.2 and Go 1.21.9 are released" link="https://groups.google.com/g/golang-announce/c/YgW0sx8mN3M" lang="en" >}}
Maintaining HPACK state requires that we parse and process all HEADERS and CONTINUATION frames on a connection. When a request's headers exceed `MaxHeaderBytes`, we don't allocate memory to store the excess headers but we do parse them. This permits an attacker to cause an HTTP/2 endpoint to read arbitrary amounts of header data, all associated with a request which is going to be rejected. These headers can include Huffman-encoded data which is significantly more expensive for the receiver to decode than for an attacker to send.
{{< /fig-quote >}}

以下未稿。

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.22.2.linux-amd64.tar.gz`](https://go.dev/dl/go1.22.2.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.22.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.22.2.linux-amd64.tar.gz
$ sudo mv go go1.22.2
$ sudo ln -s go1.22.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.22.2 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.22.2@latest
$ go1.22.2 download
$ go1.22.2 version
go version go1.22.2 linux/amd64
```

てな感じに導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-45288]: https://nvd.nist.gov/vuln/detail/CVE-2023-45288

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Effective Go -->
