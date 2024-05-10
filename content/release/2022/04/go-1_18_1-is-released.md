+++
title = "Go 1.18.1 のリリース【セキュリティ・アップデート】"
date =  "2022-04-13T20:54:58+09:00"
description = "今回は3件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.18.1 がリリースされた。

- [[security] Go 1.18.1 and Go 1.17.9 are released](https://groups.google.com/g/golang-announce/c/oecdBNLOml8)

今回は3件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://go.dev/doc/devel/release#go1.18.minor" lang="en" >}}
{{% quote %}}go1.18.1 (released 2022-04-12) includes security fixes to the `crypto/elliptic`, `crypto/x509`, and `encoding/pem` packages, as well as bug fixes to the compiler, linker, runtime, the go command, vet, and the `bytes`, `crypto/x509`, and `go/types` packages. See the Go 1.18.1 milestone on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2022-24675]: encoding/pem: fix stack overflow in Decode

{{< fig-quote type="markdown" title="Go 1.18.1 and Go 1.17.9 are released" link="https://groups.google.com/g/golang-announce/c/oecdBNLOml8" lang="en" >}}
{{% quote %}}A large (more than 5 MB) PEM input can cause a stack overflow in `Decode`, leading the program to crash{{% /quote %}}.
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

## [CVE-2022-28327]: crypto/elliptic: tolerate all oversized scalars in generic P-256

{{< fig-quote type="markdown" title="Go 1.18.1 and Go 1.17.9 are released" link="https://groups.google.com/g/golang-announce/c/oecdBNLOml8" lang="en" >}}
{{% quote %}}A crafted scalar input longer than 32 bytes can cause `P256().ScalarMult` or `P256().ScalarBaseMult` to panic. Indirect uses through `crypto/ecdsa` and `crypto/tls` are unaffected. amd64, arm64, ppc64le, and s390x are unaffected{{% /quote %}}.
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

## [CVE-2022-27536]: crypto/x509: non-compliant certificates can cause a panic in Verify on macOS in Go 1.18

{{< fig-quote type="markdown" title="Go 1.18.1 and Go 1.17.9 are released" link="https://groups.google.com/g/golang-announce/c/oecdBNLOml8" lang="en" >}}
Verifying certificate chains containing certificates which are not compliant with RFC 5280 causes Certificate.Verify to panic on macOS.

These chains can be delivered through TLS and can cause a `crypto/tls` or `net/http` client to crash.
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

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.18.1.linux-amd64.tar.gz`](https://go.dev/dl/go1.18.1.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.18.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.18.1.linux-amd64.tar.gz
$ sudo mv go go1.18.1
$ sudo ln -s go1.18.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.18.1 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-24675]: https://nvd.nist.gov/vuln/detail/CVE-2022-24675
[CVE-2022-28327]: https://nvd.nist.gov/vuln/detail/CVE-2022-28327
[CVE-2022-27536]: https://nvd.nist.gov/vuln/detail/CVE-2022-27536

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
