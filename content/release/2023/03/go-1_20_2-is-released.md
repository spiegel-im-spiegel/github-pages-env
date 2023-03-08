+++
title = "Go 1.20.2 のリリース【セキュリティ・アップデート】"
date =  "2023-03-08T18:43:32+09:00"
description = "今回は1件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.20.2 がリリースされた。

- [[security] Go 1.20.2 and Go 1.19.7 are released](https://groups.google.com/g/golang-announce/c/3-TpUx48iQY)

今回は1件の脆弱性修正を含んでいる。

## [CVE-2023-24532] crypto/elliptic: incorrect P-256 ScalarMult and ScalarBaseMult results

{{< fig-quote type="markdown" title="Go 1.20.2 and Go 1.19.7 are released" link="https://groups.google.com/g/golang-announce/c/3-TpUx48iQY" lang="en" >}}
The `ScalarMult` and `ScalarBaseMult` methods of the P256 Curve may return an incorrect result if called with some specific unreduced scalars (a scalar larger than the order of the curve).

This does not impact usages of `crypto/ecdsa` or `crypto/ecdh`.
{{< /fig-quote >}}

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.2.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.2.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.2.linux-amd64.tar.gz
$ sudo mv go go1.20.2
$ sudo ln -s go1.20.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.2 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-24532]: https://nvd.nist.gov/vuln/detail/CVE-2023-24532

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
