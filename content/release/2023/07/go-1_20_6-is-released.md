+++
title = "Go 1.20.6 のリリース【セキュリティ・アップデート】"
date =  "2023-07-12T06:12:51+09:00"
description = "今回は CVE ID ベースで1件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

珍しく予告なしのセキュリティ・アップデート。
v1.21 リリース前の駆け込みかな？ というわけで， [Go] 1.20.6 がリリースされた。

- [[security] Go 1.20.6 and Go 1.19.11 are released](https://groups.google.com/g/golang-announce/c/2q13H6LEEx0)

今回は CVE ID ベースで1件の脆弱性修正を含んでいる。

## [CVE-2023-29406] net/http: insufficient sanitization of Host header

{{< fig-quote type="markdown" title="Go 1.20.6 and Go 1.19.11 are released" link="https://groups.google.com/g/golang-announce/c/2q13H6LEEx0" lang="en" >}}
The HTTP/1 client did not fully validate the contents of the Host header. A maliciously crafted Host header could inject additional headers or entire requests. The HTTP/1 client now refuses to send requests containing an invalid `Request.Host` or `Request.URL.Host` value.
{{< /fig-quote >}}

ありゃりゃー。
これはアカンやつやね。

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:U/C:N/I:H/A:N`
- 深刻度: 警告 (Score: 6.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.6.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.6.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.6.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.6.linux-amd64.tar.gz
$ sudo mv go go1.20.6
$ sudo ln -s go1.20.6 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.6 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.20.6@latest
$ go1.20.6 download
$ go1.20.6 version
go version go1.20.6 linux/amd64
```

てな感じで導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-29406]: https://nvd.nist.gov/vuln/detail/CVE-2023-29406

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
