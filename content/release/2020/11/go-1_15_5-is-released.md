+++
title = "Go 1.15.5 のリリース【セキュリティ・アップデート】"
date =  "2020-11-14T12:24:32+09:00"
description = "今回は複数の脆弱性について改修されている。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/kMa3eup0qhU "[security] Go 1.15.5 and Go 1.14.12 pre-announcement")通り， [Go] 1.15.5 がリリースされた。

- [[security] Go 1.15.5 and Go 1.14.12 are released](https://groups.google.com/g/golang-announce/c/NpBGTTmKzpM)

今回は複数の脆弱性について改修されている。

## [math/big][big]: panic during recursive division of very large numbers ([CVE-2020-28362])

{{< fig-quote type="markdown" title="Go 1.15.5 and Go 1.14.12 are released" link="https://groups.google.com/g/golang-announce/c/NpBGTTmKzpM" lang="en" >}}
{{% quote %}}A number of [`math/big.Int`](https://pkg.go.dev/math/big#Int) methods (`Div`, `Exp`, `DivMod`, `Quo`, `Rem`, `QuoRem`, `Mod`, `ModInverse`, `ModSqrt`, `Jacobi`, and `GCD`) can panic when provided crafted large inputs. For the panic to happen, the divisor or modulo argument must be larger than 3168 bits (on 32-bit architectures) or 6336 bits (on 64-bit architectures). Multiple [`math/big.Rat`](https://pkg.go.dev/math/big#Rat) methods are similarly affected{{% /quote %}}.
{{< /fig-quote >}}

[`math/big`][big] パッケージは暗号関連のパッケージと密な関係にあるため，暗号関係の処理全般にインパクトがあると考えたほうがいいだろう。

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | なし         |
| 完全性への影響   | なし         |
| 可用性への影響   | 高           |

## cmd/go: arbitrary code execution at build time through cgo ([CVE-2020-28366], [CVE-2020-28367])

{{< fig-quote type="markdown" title="Go 1.15.5 and Go 1.14.12 are released" link="https://groups.google.com/g/golang-announce/c/NpBGTTmKzpM" lang="en" >}}
{{% quote %}}The go command may execute arbitrary code at build time when cgo is in use. This may occur when running go get on a malicious package, or any other command that builds untrusted code{{% /quote %}}.
{{< /fig-quote >}}

うわっ，ヤベ！ 相変わらず cgo は（セキュリティ的に）鬼門だなぁ。

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:R/S:U/C:H/I:H/A:H`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | 高           |
| 可用性への影響   | 高           |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.15.5.linux-amd64.tar.gz`](https://golang.org/dl/go1.15.5.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.15.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.15.5.linux-amd64.tar.gz
$ sudo mv go go1.15.5
$ sudo ln -s go1.15.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.15.5 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2020-28362]: https://nvd.nist.gov/vuln/detail/CVE-2020-28362
[CVE-2020-28366]: https://nvd.nist.gov/vuln/detail/CVE-2020-28366
[CVE-2020-28367]: https://nvd.nist.gov/vuln/detail/CVE-2020-28367
[big]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
