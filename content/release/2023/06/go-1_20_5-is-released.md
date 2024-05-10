+++
title = "Go 1.20.5 のリリース【セキュリティ・アップデート】"
date =  "2023-06-11T09:51:55+09:00"
description = "今回は CVE ID ベースで4件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/1AItFMBjrfw "[security] Go 1.20.5 and Go 1.19.10 pre-announcement")通り [Go] 1.20.5 がリリースされた。

- [[security] Go 1.20.5 and Go 1.19.10 are released](https://groups.google.com/g/golang-announce/c/q5135a9d924)

今回は CVE ID ベースで4件の脆弱性修正を含んでいる。

## [CVE-2023-29402] cmd/go: cgo code injection

{{< fig-quote type="markdown" title="Go 1.20.5 and Go 1.19.10 are released" link="https://groups.google.com/g/golang-announce/c/q5135a9d924" lang="en" >}}
The go command may generate unexpected code at build time when using cgo. This may result in unexpected behavior when running a go program which uses cgo.

This may occur when running an untrusted module which contains directories with newline characters in their names. Modules which are retrieved using the go command, i.e. via "`go get`", are not affected (modules retrieved using GOPATH-mode, i.e. `GO111MODULE=off`, may be affected).
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H`
- 深刻度: 緊急 (Score: 9.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## [CVE-2023-29403] runtime: unexpected behavior of setuid/setgid binaries

{{< fig-quote type="markdown" title="Go 1.20.5 and Go 1.19.10 are released" link="https://groups.google.com/g/golang-announce/c/q5135a9d924" lang="en" >}}
The Go runtime didn't act any differently when a binary had the `setuid/setgid` bit set. On Unix platforms, if a `setuid/setgid` binary was executed with standard I/O file descriptors closed, opening any files could result in unexpected content being read/written with elevated prilieges. Similarly if a `setuid/setgid` program was terminated, either via panic or signal, it could leak the contents of its registers.
{{< /fig-quote >}}

- `CVSS:3.1/AV:L/AC:L/PR:N/UI:R/S:U/C:H/I:H/A:H`
- 深刻度: 重要 (Score: 7.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ローカル |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## [CVE-2023-29404], [CVE-2023-29405] cmd/go: improper sanitization of LDFLAGS

{{< fig-quote type="markdown" title="Go 1.20.5 and Go 1.19.10 are released" link="https://groups.google.com/g/golang-announce/c/q5135a9d924" lang="en" >}}
The go command may execute arbitrary code at build time when using cgo. This may occur when running "`go get`" on a malicious module, or when running any other command which builds untrusted code. This is can by triggered by linker flags, specified via a "`#cgo LDFLAGS`" directive.
{{< /fig-quote >}}

これについては追加情報があって

{{< fig-quote type="markdown" title="Go 1.20.5 and Go 1.19.10 are released" link="https://groups.google.com/g/golang-announce/c/q5135a9d924" lang="en" >}}
Due to an unfortunate mistake, this change will break the use of "`#cgo LDFLAGS`" directives when using `-compiler=gccgo`. Most people using gccgo or GoLLVM use the `cmd/go` that is distributed with those tools, and that is unaffected. Therefore, we will fix this in the next minor release. The current minor releases 1.20.5 and 1.19.10 are unfortunately broken for some cases when using gccgo or GoLLVM. Our apologies for the mishap. Thanks to Jeffrey Tolar for spotting the problem.
{{< /fig-quote >}}

とのこと。
gccgo や GoLLVM を使っている人は要注意である。

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H` ([CVE-2023-29404])
- 深刻度: 緊急 (Score: 9.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H` ([CVE-2023-29405])
- 深刻度: 緊急 (Score: 9.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |


## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.5.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.5.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.5.linux-amd64.tar.gz
$ sudo mv go go1.20.5
$ sudo ln -s go1.20.5 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.5 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.20.5@latest
$ go1.20.5 download
$ go1.20.5 version
go version go1.20.5 linux/amd64
```

てな感じで導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-29402]: https://nvd.nist.gov/vuln/detail/CVE-2023-29402
[CVE-2023-29403]: https://nvd.nist.gov/vuln/detail/CVE-2023-29403
[CVE-2023-29404]: https://nvd.nist.gov/vuln/detail/CVE-2023-29404
[CVE-2023-29405]: https://nvd.nist.gov/vuln/detail/CVE-2023-29405

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
