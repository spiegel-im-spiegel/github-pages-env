+++
title = "Go 1.17.3 のリリース【セキュリティ・アップデート】"
date =  "2021-11-07T12:44:04+09:00"
description = "2件の脆弱性の修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go] 1.17.3 がリリースされたので，遅まきながら書いておく。

- [[security] Go 1.17.3 and Go 1.16.10 are released](https://groups.google.com/g/golang-announce/c/0fM21h43arc/m/a2Hu_dWRAAAJ)

今回は2件の脆弱性の修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.3 (released 2021-11-04) includes security fixes to the `archive/zip` and `debug/macho` packages, as well as bug fixes to the compiler, linker, runtime, the go command, the `misc/wasm` directory, and to the `net/http` and `syscall` packages. See the [Go 1.17.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.3+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-41772]

{{< fig-quote type="markdown" title="Go 1.17.3 and Go 1.16.10 are released" link="https://groups.google.com/g/golang-announce/c/0fM21h43arc/m/a2Hu_dWRAAAJ" lang="en" >}}
{{% quote %}}`Reader.Open` (the API implementing `io/fs.FS` introduced in Go 1.16) can be made to panic by an attacker providing either a crafted ZIP archive containing completely invalid names or an empty filename argument{{% /quote %}}.
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

## [CVE-2021-41771]

{{< fig-quote type="markdown" title="Go 1.17.3 and Go 1.16.10 are released" link="https://groups.google.com/g/golang-announce/c/0fM21h43arc/m/a2Hu_dWRAAAJ" lang="en" >}}
{{% quote %}}Malformed binaries parsed using `Open` or `OpenFat` can cause a panic when calling `ImportedSymbols`, due to an out-of-bounds slice operation{{% /quote %}}.
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

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.17.3.linux-amd64.tar.gz`](https://golang.org/dl/go1.17.3.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.17.3.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.3.linux-amd64.tar.gz
$ sudo mv go go1.17.3
$ sudo ln -s go1.17.3 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.17.3 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-41772]: https://nvd.nist.gov/vuln/detail/CVE-2021-41772
[CVE-2021-41771]: https://nvd.nist.gov/vuln/detail/CVE-2021-41771

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
