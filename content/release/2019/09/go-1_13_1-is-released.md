+++
title = "Go 1.13.1 のリリース【セキュリティ・アップデート】"
date =  "2019-09-26T20:58:03+09:00"
description = "Go 言語で Web サービスまたはクライアント・ツールを構築している人は要注意だ。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.13.1 がリリースされた。

- [[security] Go 1.13.1 and Go 1.12.10 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/cszieYyuL9Q)
- [CVE-2019-16276](https://nvd.nist.gov/vuln/detail/CVE-2019-16276) : `CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:N` (Score:7.5, Severity:High)
    - [CVE-2019-16276 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2019-16276) : `CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:L/I:L/A:L` (Score:6.3, Severity:Medium)

今回はセキュリティ・アップデートを含んでいる。
[Go 言語]で Web サービスまたはクライアント・ツールを構築している人は要注意だ。

{{< fig-quote type="markdown" title="Go 1.13.1 and Go 1.12.10 are released" link="https://groups.google.com/forum/#!topic/golang-announce/cszieYyuL9Q" lang="en" >}}
{{< quote >}}`net/http` (through `net/textproto`) used to accept and normalize invalid HTTP/1.1 headers with a space before the colon, in violation of RFC 7230. If a Go server is used behind an uncommon reverse proxy that accepts and forwards but doesn't normalize such invalid headers, the reverse proxy and the server can interpret the headers differently. This can lead to filter bypasses or [request smuggling](https://portswigger.net/research/http-desync-attacks-request-smuggling-reborn), the latter if requests from separate clients are multiplexed onto the same upstream connection by the proxy. Such invalid headers are now rejected by Go servers, and passed without normalization to Go client applications.{{< /quote >}}
{{< /fig-quote >}}

[Ubuntu] の APT は相変わらずサポートから外れた 1.10.x しか対応していないので[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.13.1.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.1.linux-amd64.tar.gz
$ sudo mv go go1.13.1
$ sudo ln -s go1.13.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.1 linux/amd64
```

アップデートは計画的に。

## ブックマーク

- [HTTP Desync Attacks: Request Smuggling Reborn | PortSwigger Research](https://portswigger.net/research/http-desync-attacks-request-smuggling-reborn)

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
