+++
title = "Go 1.20.4 のリリース【セキュリティ・アップデート】"
date =  "2023-05-03T12:44:26+09:00"
description = "今回は3件の脆弱性修正を含んでいる。今回は html/template の不具合ばっかりだな。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/vFRFE07dbB8 "[security] Go 1.20.4 and Go 1.19.9 pre-announcement")通り [Go] 1.20.4 がリリースされた。

- [[security] Go 1.20.4 and Go 1.19.9 are released](https://groups.google.com/g/golang-announce/c/MEb0UyuSMsU)

今回は3件の脆弱性修正を含んでいる。
今回は [`html/template`](https://pkg.go.dev/html/template "template package - html/template - Go Packages") の不具合ばっかりだな。
Web アプリケーションを組んでいる方は特に気をつけましょう。

## [CVE-2023-24539] html/template: improper sanitization of CSS values

{{< fig-quote type="markdown" title="Go 1.20.4 and Go 1.19.9 are released" link="https://groups.google.com/g/golang-announce/c/MEb0UyuSMsU" lang="en" >}}
Angle brackets (`<>`) were not considered dangerous characters when inserted into CSS contexts. Templates containing multiple actions separated by a '`/`' character could result in unexpectedly closing the CSS context and allowing for injection of unexpected HMTL, if executed with untrusted input.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L`
- 深刻度: 重要 (Score: 7.3)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 低 |
| 完全性への影響 | 低 |
| 可用性への影響 | 低 |

## [CVE-2023-24540] html/template: improper handling of JavaScript whitespace

{{< fig-quote type="markdown" title="Go 1.20.4 and Go 1.19.9 are released" link="https://groups.google.com/g/golang-announce/c/MEb0UyuSMsU" lang="en" >}}
Not all valid JavaScript whitespace characters were considered to be whitespace. Templates containing whitespace characters outside of the character set `"\t\n\f\r\u0020\u2028\u2029"` in JavaScript contexts that also contain actions may not be properly sanitized during execution.
{{< /fig-quote >}}

うわぁ，面倒くさい。

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

## [CVE-2023-29400] html/template: improper handling of empty HTML attributes

{{< fig-quote type="markdown" title="Go 1.20.4 and Go 1.19.9 are released" link="https://groups.google.com/g/golang-announce/c/MEb0UyuSMsU" lang="en" >}}
Templates containing actions in unquoted HTML attributes (e.g. `"attr={{.}}"`) executed with empty input could result in output that would have unexpected results when parsed due to HTML normalization rules. This may allow injection of arbitrary attributes into tags.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:L`
- 深刻度: 重要 (Score: 7.3)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 低 |
| 完全性への影響 | 低 |
| 可用性への影響 | 低 |

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.20.4.linux-amd64.tar.gz`](https://go.dev/dl/go1.20.4.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.20.4.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.20.4.linux-amd64.tar.gz
$ sudo mv go go1.20.4
$ sudo ln -s go1.20.4 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.20.4 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.20.4@latest
$ go1.20.4 download
$ go1.20.4 version
go version go1.20.4 linux/amd64
```

てな感じで導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-24539]: https://nvd.nist.gov/vuln/detail/CVE-2023-24539
[CVE-2023-24540]: https://nvd.nist.gov/vuln/detail/CVE-2023-24540
[CVE-2023-29400]: https://nvd.nist.gov/vuln/detail/CVE-2023-29400

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
