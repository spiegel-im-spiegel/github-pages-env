+++
title = "Go 1.21.1 のリリース【セキュリティ・アップデート】"
date =  "2023-09-18T10:27:45+09:00"
description = "CVE ID ベースで5件の脆弱性の修正がある。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらでごめんペコン。
[予告](https://groups.google.com/g/golang-announce/c/UXJQvKffcao "[security] Go 1.21.1 and Go 1.20.8 pre-announcement")どおり [Go] 1.21.1 がリリースされた。

- [[security] Go 1.21.1 and Go 1.20.8 are released](https://groups.google.com/g/golang-announce/c/Fm51GRLNRvM)

今回は CVE ID ベースで5件の脆弱性修正を含んでいる。

## [CVE-2023-39320] cmd/go: go.mod toolchain directive allows arbitrary execution

{{< fig-quote type="markdown" title="Go 1.21.1 and Go 1.20.8 are released" link="https://groups.google.com/g/golang-announce/c/Fm51GRLNRvM" lang="en" >}}
The `go.mod` toolchain directive, introduced in Go 1.21, could be leveraged to execute scripts and binaries relative to the root of the module when the "`go`" command was executed within the module. This applies to modules downloaded using the "`go`" command from the module proxy, as well as modules downloaded directly using VCS software.
{{< /fig-quote >}}

またコマンドの任意実行か。
ヤバいやつ。

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

## [CVE-2023-39318] html/template: improper handling of HTML-like comments within script contexts

{{< fig-quote type="markdown" title="Go 1.21.1 and Go 1.20.8 are released" link="https://groups.google.com/g/golang-announce/c/Fm51GRLNRvM" lang="en" >}}
The `html/template` package did not properly handle HMTL-like "`<!--`" and "`-->`" comment tokens, nor hashbang "`#!`" comment tokens, in `<script>` contexts. This may cause the template parser to improperly interpret the contents of `<script>` contexts, causing actions to be improperly escaped. This could be leveraged to perform an XSS attack.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N`
- 深刻度: 警告 (Score: 6.1)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更あり |
| 機密性への影響 | 低 |
| 完全性への影響 | 低 |
| 可用性への影響 | なし |

## [CVE-2023-39319] crypto/tls: panic when processing post-handshake message on QUIC connections

{{< fig-quote type="markdown" title="Go 1.21.1 and Go 1.20.8 are released" link="https://groups.google.com/g/golang-announce/c/Fm51GRLNRvM" lang="en" >}}
The `html/template` package did not apply the proper rules for handling occurrences of "`<script`", "`<!--`", and "`</script`" within JS literals in `<script>` contexts. This may cause the template parser to improperly consider script contexts to be terminated early, causing actions to be improperly escaped. This could be leveraged to perform an XSS attack.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N`
- 深刻度: 警告 (Score: 6.1)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 要 |
| スコープ | 変更あり |
| 機密性への影響 | 低 |
| 完全性への影響 | 低 |
| 可用性への影響 | なし |

## [CVE-2023-39321], [CVE-2023-39322] html/template: improper handling of special tags within script contexts

{{< fig-quote type="markdown" title="Go 1.21.1 and Go 1.20.8 are released" link="https://groups.google.com/g/golang-announce/c/Fm51GRLNRvM" lang="en" >}}
Processing an incomplete post-handshake message for a QUIC connection caused a panic.
{{< /fig-quote >}}

CVE ID ベースでは2件だけど CVSS 評価は同じようだ。

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

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.21.1.linux-amd64.tar.gz`](https://go.dev/dl/go1.21.1.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.21.1.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.21.1.linux-amd64.tar.gz
$ sudo mv go go1.21.1
$ sudo ln -s go1.21.1 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.21.1 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.21.1@latest
$ go1.21.1 download
$ go1.21.1 version
go version go1.21.1 linux/amd64
```

てな感じに導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-39320]: https://nvd.nist.gov/vuln/detail/CVE-2023-39320
[CVE-2023-39318]: https://nvd.nist.gov/vuln/detail/CVE-2023-39318
[CVE-2023-39319]: https://nvd.nist.gov/vuln/detail/CVE-2023-39319
[CVE-2023-39321]: https://nvd.nist.gov/vuln/detail/CVE-2023-39321
[CVE-2023-39322]: https://nvd.nist.gov/vuln/detail/CVE-2023-39322

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
