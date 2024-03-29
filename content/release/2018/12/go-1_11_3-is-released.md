+++
title = "Go 1.11.3 のリリース 【セキュリティ・アップデート】"
date = "2018-12-14T22:10:04+09:00"
description = "3つのインシデントに対して修正が行われている。深刻度は高めで早めのアップデート推奨。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "security", "vulnerability" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]コンパイラの 1.11.3 および 1.10.6 がリリースされた。
今回はセキュリティ・アップデートを含んでいるので必ずアップデートすること。

- [[security] Go 1.11.3 and Go 1.10.6 are released - Google group](https://groups.google.com/forum/#%21topic/golang-announce/Kw31K8G7Fi0)
- [Go 1.11.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.11.3)

以下の3つのインシデントに対して修正が行われている。

{{% fig-quote type="markdown" title="Go 1.11.3 and Go 1.10.6 are released" link="https://groups.google.com/forum/#%21topic/golang-announce/Kw31K8G7Fi0" %}}
- cmd/go: remote command execution during "`go get -u`" : The issue is CVE-2018-16873 and Go issue [golang.org/issue/29230](https://golang.org/issue/29230). See the Go issue for details. Thanks to Etienne Stalmans from the Heroku platform security team for discovering and reporting this issue.
- cmd/go: directory traversal in "`go get`" via curly braces in import paths : The issue is CVE-2018-16874 and Go issue [golang.org/issue/29231](https://golang.org/issue/29231). See the Go issue for details. Thanks to ztz of Tencent Security Platform for discovering and reporting this issue.
- crypto/x509: CPU denial of service in chain validation : The issue is CVE-2018-16875 and Go issue [golang.org/issue/29233](https://golang.org/issue/29233). See the Go issue for details. Thanks to Netflix for discovering and reporting this issue.
{{% /fig-quote %}}

## 想定される影響

### [CVE-2018-16873](https://nvd.nist.gov/vuln/detail/CVE-2018-16873)

- [CVE-2018-16873 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-16873) 

`CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:U/C:H/I:H/A:H` (深刻度7.5) : 

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

### [CVE-2018-16874](https://nvd.nist.gov/vuln/detail/CVE-2018-16874)

- [CVE-2018-16874 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-16874)

`CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:U/C:H/I:H/A:N` (深刻度6.8) : 

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 高           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 要           |
| スコープ         | 変更なし     |
| 機密性への影響   | 高           |
| 完全性への影響   | 高           |
| 可用性への影響   | なし         |

### [CVE-2018-16875](https://nvd.nist.gov/vuln/detail/CVE-2018-16875)

- [CVE-2018-16875 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-16875)

`CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:H` (深刻度5.9) : 

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

## その他

これとは別に “`go get github.com/golang/pkg/...`” でコードを再帰的にダウンロードしてしまう問題があるそうで，これは次回の 1.11.4 および Go 1.10.7 で解消されるとのこと。

アップデートは計画的に。

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
