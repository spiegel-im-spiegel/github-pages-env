+++
title = "Go 1.17.8 のリリース【セキュリティ・アップデート】"
date =  "2022-03-04T20:20:52+09:00"
description = "今回は1件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.8 がリリースされた。

- [[security] Go 1.17.8 and Go 1.16.15 are released](https://groups.google.com/g/golang-announce/c/RP1hfrBYVuk)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.8 (released 2022-03-03) includes a security fix to the `regexp/syntax` package, as well as bug fixes to the compiler, runtime, the go command, and the `crypto/x509`, and net packages. See the [Go 1.17.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.8+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2022-24921]: regexp: stack exhaustion compiling deeply nested expressions

{{< fig-quote type="markdown" title="Go 1.17.7 and Go 1.16.14 are released" link="https://groups.google.com/g/golang-announce/c/SUsQn0aSgPQ" lang="en" >}}
{{% quote %}}On 64-bit platforms, an extremely deeply nested expression can cause `regexp.Compile` to cause goroutine stack exhaustion, forcing the program to exit. Note this applies to very large expressions, on the order of 2MB{{% /quote %}}.
{{< /fig-quote >}}

2MBの式とか（笑） まぁ，でも， SQL 文とかでも，とてつもなくデカイ式を書く奴とかいるからなぁ。

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

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.17.8.linux-amd64.tar.gz`](https://go.dev/dl/go1.17.8.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.17.8.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.8.linux-amd64.tar.gz
$ sudo mv go go1.17.8
$ sudo ln -s go1.17.8 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.8 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2022-24921]: https://nvd.nist.gov/vuln/detail/CVE-2022-24921

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
