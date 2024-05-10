+++
title = "Go 1.18.5 のリリース【セキュリティ・アップデート】"
date =  "2022-08-03T07:57:41+09:00"
description = "今回は1件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まだ先だろうと高をくくって更新をサボってたら [Go] 1.19 が出てるよ。
というわけで，まずは 1.18 系のセキュリティ・アップデートの記事からやっつける。

- [[security] Go 1.18.5 and Go 1.17.13 are released](https://groups.google.com/g/golang-announce/c/YqYYG87xB10)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://go.dev/doc/devel/release#go1.18.minor" lang="en" >}}
{{% quote %}}go1.18.5 (released 2022-08-01) includes security fixes to the `encoding/gob` and `math/big` packages, as well as bug fixes to the compiler, the go command, the runtime, and the testing package. See the [Go 1.18.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.5+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2022-32189] encoding/gob & math/big: decoding big.Float and big.Rat can panic

{{< fig-quote type="markdown" title="Go 1.18.5 and Go 1.17.13 are released" link="https://groups.google.com/g/golang-announce/c/YqYYG87xB10" lang="en" >}}
{{% quote %}}Decoding big.Float and big.Rat types can panic if the encoded message is too short{{% /quote %}}.
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

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.18.5.linux-amd64.tar.gz`](https://go.dev/dl/go1.18.5.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.18.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.18.5.linux-amd64.tar.gz
$ sudo mv go go1.18.5
$ sudo ln -s go1.18.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.18.5 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-32189]: https://nvd.nist.gov/vuln/detail/CVE-2022-32189

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
