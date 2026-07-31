+++
title = "Go 1.17.7 のリリース【セキュリティ・アップデート】"
date =  "2022-02-11T18:30:25+09:00"
description = "今回は3件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.7 がリリースされた。

- [[security] Go 1.17.7 and Go 1.16.14 are released](https://groups.google.com/g/golang-announce/c/SUsQn0aSgPQ)

今回は3件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.7 (released 2022-02-10) includes security fixes to the `crypto/elliptic`, `math/big` packages and to the go command, as well as bug fixes to the compiler, linker, runtime, the go command, and the `debug/macho`, `debug/pe`, and `net/http/httptest` packages. See the [Go 1.17.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.7+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2022-23806]: crypto/elliptic: fix IsOnCurve for big.Int values that are not valid coordinates

{{< fig-quote type="markdown" title="Go 1.17.7 and Go 1.16.14 are released" link="https://groups.google.com/g/golang-announce/c/SUsQn0aSgPQ" lang="en" >}}
{{% quote %}}Some `big.Int` values that are not valid field elements (negative or overflowing) might cause `Curve.IsOnCurve` to incorrectly return true. Operating on those values may cause a panic or an invalid curve operation. Note that `Unmarshal` will never return such values{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:H`
- 深刻度: 緊急 (Score: 9.1)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## [CVE-2022-23772]: math/big: prevent large memory consumption in Rat.SetString

{{< fig-quote type="markdown" title="Go 1.17.7 and Go 1.16.14 are released" link="https://groups.google.com/g/golang-announce/c/SUsQn0aSgPQ" lang="en" >}}
{{% quote %}}An attacker can cause unbounded memory growth in a program using `(*Rat).SetString` due to an unhandled overflow{{% /quote %}}.
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

## [CVE-2022-23773]: cmd/go: prevent branches from materializing into versions

{{< fig-quote type="markdown" title="Go 1.17.7 and Go 1.16.14 are released" link="https://groups.google.com/g/golang-announce/c/SUsQn0aSgPQ" lang="en" >}}
{{% quote %}}A branch whose name resembles a version tag (such as "`v1.0.0`" or "`subdir/v2.0.0-dev`") can be considered a valid version by the go command. Materializing versions from branches might be unexpected and bypass ACLs that limit the creation of tags but not branches{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:N`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.17.7.linux-amd64.tar.gz`](https://go.dev/dl/go1.17.7.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.17.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.7.linux-amd64.tar.gz
$ sudo mv go go1.17.7
$ sudo ln -s go1.17.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.7 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2022-23806]: https://nvd.nist.gov/vuln/detail/CVE-2022-23806
[CVE-2022-23772]: https://nvd.nist.gov/vuln/detail/CVE-2022-23772
[CVE-2022-23773]: https://nvd.nist.gov/vuln/detail/CVE-2022-23773

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
