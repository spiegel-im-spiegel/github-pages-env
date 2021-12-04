+++
title = "Go 1.14.5 のリリース【セキュリティ・アップデート】"
date =  "2020-07-15T09:47:40+09:00"
description = "今回は2件のセキュリティ・アップデートを含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先週[予告されていた](https://groups.google.com/g/golang-announce/c/f2c5bqrGH_g "[security] Go 1.14.5 and Go 1.13.13 pre-announcement")とおり， [Go] 1.14.5 がリリースされた。

- [[security] Go 1.14.5 and Go 1.13.13 are released](https://groups.google.com/g/golang-announce/c/XZNfaiwgt2w)

2件のセキュリティ・アップデートを含んでいる。

## 【[CVE-2020-15586]】 Data race in certain net/http servers including ReverseProxy

{{< fig-quote type="markdown" title="[security] Go 1.14.5 and Go 1.13.13 are released" link="https://groups.google.com/g/golang-announce/c/XZNfaiwgt2w" lang="en" >}}
Servers where the Handler concurrently reads the request body and writes a response can encounter a data race and crash. The `httputil.ReverseProxy` Handler is affected.<br>
Thanks to Mikael Manukyan, Andrew Kutz, Dave McClure, Tim Downey, Clay Kauzlaric, and Gabe Rosenhouse for reporting this issue.
This issue is CVE-2020-15586 and Go issue [golang.org/issue/34902](https://golang.org/issue/34902).
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 警告 (5.9)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | なし         |
| 完全性への影響   | なし         |
| 可用性への影響   | 高           |

## 【[CVE-2020-14039]】 X.509 verification ignores provided EKUs on Windows

{{< fig-quote type="markdown" title="[security] Go 1.14.5 and Go 1.13.13 are released" link="https://groups.google.com/g/golang-announce/c/XZNfaiwgt2w" lang="en" >}}
On Windows, if [`VerifyOptions.Roots`](https://pkg.go.dev/crypto/x509?tab=doc#VerifyOptions.Roots) is nil, [`Certificate.Verify`](https://pkg.go.dev/crypto/x509?tab=doc#VerifyOptions.Roots) does not check the EKU requirements specified in [`VerifyOptions.KeyUsages`](https://pkg.go.dev/crypto/x509?tab=doc#VerifyOptions.KeyUsages).
Thanks to Niall Newman for reporting this issue.
This issue is CVE-2020-14039 and Go issue [golang.org/issue/39360](https://golang.org/issue/39360).
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:L/A:N`
- 深刻度: 警告 (5.3)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | なし         |
| 完全性への影響   | 低           |
| 可用性への影響   | なし         |

## 例によって

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.14.5.linux-amd64.tar.gz`](https://golang.org/dl/go1.14.5.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する（もしくは自力でコンパイルするか）。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.14.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.14.5.linux-amd64.tar.gz
$ sudo mv go go1.14.5
$ sudo ln -s go1.14.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.14.5 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[CVE-2020-15586]: https://nvd.nist.gov/vuln/detail/CVE-2020-15586
[CVE-2020-14039]: https://nvd.nist.gov/vuln/detail/CVE-2020-14039
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
