+++
title = "Go 1.19.3 のリリース【セキュリティ・アップデート】"
date =  "2022-11-04T09:40:28+09:00"
description = "今回は1件の脆弱性修正を含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/dRtDK7WS78g "[security] Go 1.19.3 and Go 1.18.8 pre-announcement")通り 2022-11-01 (現地時間) に [Go] 1.19.3 がリリースされた。

- [[security] Go 1.19.3 and Go 1.18.8 are released](https://groups.google.com/g/golang-announce/c/mbHY1UY3BaM)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://go.dev/doc/devel/release#go1.19.minor" lang="en" >}}
go1.19.3 (released 2022-11-01) includes security fixes to the `os/exec` and `syscall` packages, as well as bug fixes to the compiler and the runtime. See the [Go 1.19.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.3+label%3ACherryPickApproved) on our issue tracker for details.
{{< /fig-quote >}}

## [CVE-2022-41716] syscall, os/exec: unsanitized NUL in environment variables

{{< fig-quote type="markdown" title="Go 1.19.3 and Go 1.18.8 are released" link="https://groups.google.com/g/golang-announce/c/mbHY1UY3BaM" lang="en" >}}
On Windows, `syscall.StartProcess` and `os/exec.Cmd` did not properly check for invalid environment variable values. A malicious environment variable value could exploit this behavior to set a value for a different environment variable. For example, the environment variable string `"A=B\x00C=D"` set the variables `"A=B"` and `"C=D"`.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:N`
- 深刻度: 重要 (Score: 7.5)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | なし |
| 完全性への影響 | 高 |
| 可用性への影響 | なし |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.19.3.linux-amd64.tar.gz`](https://go.dev/dl/go1.19.3.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.19.3.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.19.3.linux-amd64.tar.gz
$ sudo mv go go1.19.3
$ sudo ln -s go1.19.3 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.19.3 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-41716]: https://nvd.nist.gov/vuln/detail/CVE-2022-41716

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
