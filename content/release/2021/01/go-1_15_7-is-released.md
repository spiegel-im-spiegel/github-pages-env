+++
title = "Go 1.15.7 のリリース【セキュリティ・アップデート】"
date =  "2021-01-21T19:45:21+09:00"
description = "今回は複数の脆弱性について改修されている。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/KvrRblbXp_w "[security] Go 1.15.7 and Go 1.14.14 pre-announcement")通り， [Go] 1.15.7 がリリースされた。

- [[security] Go 1.15.7 and Go 1.14.14 are released](https://groups.google.com/g/golang-announce/c/mperVMGa98w)

今回は複数の脆弱性について改修されている。

## cmd/go: packages using cgo can cause arbitrary code execution at build time ([CVE-2021-3115])

{{< fig-quote type="markdown" title="Go 1.15.7 and Go 1.14.14 are released" link="https://groups.google.com/g/golang-announce/c/mperVMGa98w" lang="en" >}}
The go command may execute arbitrary code at build time when cgo is in use on Windows. This may occur when running “`go get`”, or any other command that builds code. Only users who build untrusted code (and don’t execute it) are affected.

In addition to Windows users, this can also affect Unix users who have “`.`” listed explicitly in their `PATH` and are running “`go get`” or build commands outside of a module or with module mode disabled.
{{< /fig-quote >}}

というわけで， Windows だけでなく UNIX 系のプラットフォームでも環境変数 `PATH` にカレントディレクトリ “`.`” が設定されているとヤバいので（そんなヤツおらんやろけど），きちんと対処すること。

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:R/S:U/C:H/I:H/A:H`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 高 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## crypto/elliptic: incorrect operations on the P-224 curve ([CVE-2021-3114])

{{< fig-quote type="markdown" title="Go 1.15.7 and Go 1.14.14 are released" link="https://groups.google.com/g/golang-announce/c/mperVMGa98w" lang="en" >}}
{{% quote %}}The `P224()` Curve implementation can in rare circumstances generate incorrect outputs, including returning invalid points from `ScalarMult`{{% /quote %}}.
{{< /fig-quote >}}

ただし

{{< fig-quote type="markdown" title="Go 1.15.7 and Go 1.14.14 are released" link="https://groups.google.com/g/golang-announce/c/mperVMGa98w" lang="en" >}}
{{% quote %}}The `crypto/x509` and [`golang.org/x/crypto/ocsp`](http://golang.org/x/crypto/ocsp) (but not `crypto/tls`) packages support P-224 ECDSA keys, but they are not supported by publicly trusted certificate authorities. No other standard library or [`golang.org/x/crypto`](http://golang.org/x/crypto) package supports or uses the P-224 curve.{{% /quote %}}.
{{< /fig-quote >}}

ということなので，実質的なインパクトは小さい？

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:N`
- 深刻度: 警告 (Score: 6.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 低 |
| 完全性への影響 | 低 |
| 可用性への影響 | なし |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.15.7.linux-amd64.tar.gz`](https://golang.org/dl/go1.15.7.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.15.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.15.7.linux-amd64.tar.gz
$ sudo mv go go1.15.7
$ sudo ln -s go1.15.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.15.7 linux/amd64
```

アップデートは計画的に。

## ブックマーク

- [Command PATH security in Go - The Go Blog](https://blog.golang.org/path-security)
- [Go でサブプロセスを起動する際は LookPath に気をつけろ！](https://zenn.dev/spiegel/articles/20201107-lookpath-by-golang)

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-3115]: https://nvd.nist.gov/vuln/detail/CVE-2021-3115
[CVE-2021-3114]: https://nvd.nist.gov/vuln/detail/CVE-2021-3114

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
