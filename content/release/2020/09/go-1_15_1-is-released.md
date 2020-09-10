+++
title = "Go 1.15.1 のリリース【セキュリティ・アップデート】"
date =  "2020-09-02T12:58:47+09:00"
description = "CVE-2020-24553 の改修を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週の[予告](https://groups.google.com/g/golang-announce/c/JvvJpgiIfFI "[security] Go 1.15.1 and Go 1.14.8 pre-announcement")通り， [Go] 1.15.1 がリリースされた。

- [[security] Go 1.15.1 and Go 1.14.8 are released](https://groups.google.com/g/golang-announce/c/8wqlSbkLdPs)

以下の脆弱性の改修を含んでいる。

## [CVE-2020-24553]

{{< fig-quote type="markdown" title="[security] Go 1.15.1 and Go 1.14.8 are released" link="https://groups.google.com/g/golang-announce/c/8wqlSbkLdPs" lang="en" >}}
{{% quote %}}When a Handler does not explicitly set the Content-Type header, the [`net/http/cgi`](https://pkg.go.dev/net/http/cgi?tab=doc) and [`net/http/fcgi`](https://pkg.go.dev/net/http/fcgi?tab=doc) packages would default to “`text/html`”, which could cause a Cross-Site Scripting vulnerability if an attacker can control any part of the contents of a response{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N`
- 深刻度: 警告 (6.1)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更あり     |
| 機密性への影響   | 低           |
| 完全性への影響   | 低           |
| 可用性への影響   | なし         |

## 例によって

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.15.1.linux-amd64.tar.gz`](https://golang.org/dl/go1.15.1.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する（もしくは自力でコンパイルするか）。

https://golang.org/dl/go1.15.1.linux-amd64.tar.gz

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.15.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.15.1.linux-amd64.tar.gz
$ sudo mv go go1.15.1
$ sudo ln -s go1.15.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.15.1 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[CVE-2020-24553]: https://nvd.nist.gov/vuln/detail/CVE-2020-24553
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
