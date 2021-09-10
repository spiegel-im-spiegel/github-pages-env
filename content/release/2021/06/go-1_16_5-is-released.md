+++
title = "Go 1.16.5 のリリース【セキュリティ・アップデート】"
date =  "2021-06-04T19:46:27+09:00"
description = "4件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.16.5 がリリースされた。

- [Go 1.16.5 and Go 1.15.13 are released](https://groups.google.com/g/golang-announce/c/RgCMkAEQjSI/)

今回は4件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.16.minor" lang="en" >}}
{{% quote %}}go1.16.5 (released 2021-06-03) includes security fixes to the `archive/zip`, `math/big`, net, and `net/http/httputil` packages, as well as bug fixes to the linker, the go command, and the `net/http` package. See the [Go 1.16.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.5+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

（以下未稿）

## [CVE-2021-33198]

{{< fig-quote type="markdown" title="Go 1.16.5 and Go 1.15.13 are released" link="https://groups.google.com/g/golang-announce/c/RgCMkAEQjSI/" lang="en" >}}
{{% quote %}}The `SetString` and `UnmarshalText` methods of [`math/big.Rat`](https://pkg.go.dev/math/big#Rat) may cause a panic or an unrecoverable fatal error if passed inputs with very large exponents{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | なし |
| 可用性への影響 | 高 |

## [CVE-2021-33197]

{{< fig-quote type="markdown" title="Go 1.16.5 and Go 1.15.13 are released" link="https://groups.google.com/g/golang-announce/c/RgCMkAEQjSI/" lang="en" >}}
{{% quote %}}`ReverseProxy` in [`net/http/httputil`](https://pkg.go.dev/net/http/httputil) could be made to forward certain hop-by-hop headers, including Connection. In case the target of the `ReverseProxy` was itself a reverse proxy, this would let an attacker drop arbitrary headers, including those set by the `ReverseProxy.Director`{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:L/A:N`
- 深刻度: 警告 (Score: 5.3)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | 低 |
| 可用性への影響 | なし |

## [CVE-2021-33195]

{{< fig-quote type="markdown" title="Go 1.16.5 and Go 1.15.13 are released" link="https://groups.google.com/g/golang-announce/c/RgCMkAEQjSI/" lang="en" >}}
{{% quote %}}The `LookupCNAME`, `LookupSRV`, `LookupMX`, `LookupNS`, and `LookupAddr` functions in [`net`](https://pkg.go.dev/net), and their respective methods on the [`Resolver`](https://pkg.go.dev/net#Resolver) type may return arbitrary values retrieved from DNS which do not follow the established [RFC 1035](https://datatracker.ietf.org/doc/html/rfc1035) rules for domain names. If these names are used without further sanitization, for instance unsafely included in HTML, they may allow for injection of unexpected content. Note that `LookupTXT` may still return arbitrary values that could require sanitization before further use{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L`
- 深刻度: 重要 (Score: 7.3)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 低 |
| 完全性への影響 | 低 |
| 可用性への影響 | 低 |

## [CVE-2021-33196]

{{< fig-quote type="markdown" title="Go 1.16.5 and Go 1.15.13 are released" link="https://groups.google.com/g/golang-announce/c/RgCMkAEQjSI/" lang="en" >}}
{{% quote %}}The `NewReader` and `OpenReader` functions in [`archive/zip`](https://pkg.go.dev/archive/zip) can cause a panic or an unrecoverable fatal error when reading an archive that claims to contain a large number of files, regardless of its actual size{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | なし |
| 可用性への影響 | 高 |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.16.5.linux-amd64.tar.gz`](https://golang.org/dl/go1.16.5.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.16.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.16.5.linux-amd64.tar.gz
$ sudo mv go go1.16.5
$ sudo ln -s go1.16.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.16.5 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-33198]: https://nvd.nist.gov/vuln/detail/CVE-2021-33198
[CVE-2021-33197]: https://nvd.nist.gov/vuln/detail/CVE-2021-33197
[CVE-2021-33195]: https://nvd.nist.gov/vuln/detail/CVE-2021-33195
[CVE-2021-33196]: https://nvd.nist.gov/vuln/detail/CVE-2021-33196

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
