+++
title = "go-cvss パッケージ正式版 v1.0.0 をリリースした"
date =  "2022-02-20T09:47:18+09:00"
description = "PR をいただいたコードをベースにちょろんと機能を足して正式版 v1.0.0 としてリリースした。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "package", "module", "golang", "cvss" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私がモタモタしている間に CVSS v3.1 の環境評価基準（Environmental Metrics）のコードを PR でいただきました。
ありがたや {{% emoji "ペコン" %}}

いただいたコードをベースにちょろんと機能を足して正式版 v1.0.0 としてリリースした。

- [Release v1.0.0 · spiegel-im-spiegel/go-cvss · GitHub](https://github.com/spiegel-im-spiegel/go-cvss/releases/tag/v1.0.0)

当初「欲しい」と思っていた機能は網羅されたので正式版でいいかな，と。
まぁ，半分以上貰ったコードなんだけどね（笑） 改めまして，ありがとうございます {{% emoji "ペコン" %}}

ベンダ等が出す基本評価基準（Base Metrics）の処理は今までどおり。


```go
//go:build run
// +build run

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/go-cvss/v3/metric"
	"github.com/spiegel-im-spiegel/go-cvss/v3/report"
)

var template = `| {{ .BaseMetrics }} | {{ .BaseMetricValue }} |
|--------|-------|
| {{ .AVName }} | {{ .AVValue }} |
| {{ .ACName }} | {{ .ACValue }} |
| {{ .PRName }} | {{ .PRValue }} |
| {{ .UIName }} | {{ .UIValue }} |
| {{ .SName }} | {{ .SValue }} |
| {{ .CName }} | {{ .CValue }} |
| {{ .IName }} | {{ .IValue }} |
| {{ .AName }} | {{ .AValue }} |
`

func main() {
	bm, err := metric.NewBase().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H") //CVE-2021-44716: net/http: limit growth of header canonicalization cache
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	r, err := report.NewBase(bm).ExportWithString(template)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
```

てな感じに書けば

```text
$ go run sample1.go 
| Base Metrics | Metric Value |
|--------|-------|
| Attack Vector | Network |
| Attack Complexity | Low |
| Privileges Required | None |
| User Interaction | None |
| Scope | Changed |
| Confidentiality Impact | High |
| Integrity Impact | High |
| Availability Impact | High |
```

と出力される。
また

```go
r, err := report.NewBase(bm).ExportWithString(template)
```

を

```go
r, err := report.NewBase(bm, report.WithOptionsLanguage(language.Japanese)).ExportWithString(template)
```

と書き換えれば

```text
$ go run sample1.go 
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
```

と日本語になる（英語と日本語しか対応してません）。

今回の環境評価基準までのベクタ値を含めた評価であれば

```go
//go:build run
// +build run

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/go-cvss/v3/metric"
	"github.com/spiegel-im-spiegel/go-cvss/v3/report"
	"golang.org/x/text/language"
)

var template = `- CVSS Version {{ .Version }}
- Vector: {{ .Vector }}

## Base Metrics

- Base Score: {{ .BaseScore }}

| {{ .BaseMetrics }} | {{ .BaseMetricValue }} |
|--------|-------|
| {{ .AVName }} | {{ .AVValue }} |
| {{ .ACName }} | {{ .ACValue }} |
| {{ .PRName }} | {{ .PRValue }} |
| {{ .UIName }} | {{ .UIValue }} |
| {{ .SName }} | {{ .SValue }} |
| {{ .CName }} | {{ .CValue }} |
| {{ .IName }} | {{ .IValue }} |
| {{ .AName }} | {{ .AValue }} |

## Temporal Metrics

- Temporal Score: {{ .TemporalScore }}
- {{ .SeverityName }}: {{ .SeverityValue }}

| {{ .TemporalMetrics }} | {{ .TemporalMetricValue }} |
|--------|-------|
| {{ .EName }} | {{ .EValue }} |
| {{ .RLName }} | {{ .RLValue }} |
| {{ .RCName }} | {{ .RCValue }} |

## Environmental Metrics

- {{ .SeverityName }}: {{ .SeverityValue }} ({{ .EnvironmentalScore }})

| {{ .EnvironmentalMetrics }} | {{ .EnvironmentalMetricValue }} |
|--------|-------|
| {{ .CRName }} | {{ .CRValue }} |
| {{ .IRName }} | {{ .IRValue }} |
| {{ .ARName }} | {{ .ARValue }} |
| {{ .MAVName }} | {{ .MAVValue }} |
| {{ .MACName }} | {{ .MACValue }} |
| {{ .MPRName }} | {{ .MPRValue }} |
| {{ .MUIName }} | {{ .MUIValue }} |
| {{ .MSName }}  | {{ .MSValue }} |
| {{ .MCName }}  | {{ .MCValue }} |
| {{ .MIName }}  | {{ .MIValue }} |
| {{ .MAName }}  | {{ .MAValue }} |
`

func main() {
	em, err := metric.NewEnvironmental().Decode("CVSS:3.1/AV:P/AC:H/PR:H/UI:N/S:U/C:H/I:H/A:H/E:F/RL:U/RC:C/CR:M/IR:H/AR:M/MAV:L/MAC:H/MPR:L/MUI:R/MS:U/MC:L/MI:H/MA:L") //Random CVSS Vector
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	r, err := report.NewEnvironmental(em, report.WithOptionsLanguage(language.Japanese)).ExportWith(strings.NewReader(template))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
```

などとすれば

```text
$ go run sample2.go 
- CVSS Version 3.1
- Vector: CVSS:3.1/AV:P/AC:H/PR:H/UI:N/S:U/C:H/I:H/A:H/CR:M/IR:H/AR:M/MAV:L/MAC:H/MPR:L/MUI:R/MS:U/MC:L/MI:H/MA:L

## Base Metrics

- Base Score: 6.1

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | 物理 |
| 攻撃条件の複雑さ | 高 |
| 必要な特権レベル | 高 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## Temporal Metrics

- Temporal Score: 6
- 深刻度: 警告

| 現状評価基準 | 評価値 |
|--------|-------|
| 攻撃される可能性 | 攻撃可能 |
| 利用可能な対策のレベル | なし |
| 脆弱性情報の信頼性 | 確認済 |

## Environmental Metrics

- 深刻度: 警告 (6.5)

| 環境評価基準 | 評価値 |
|--------|-------|
| 機密性の要求度 | 中 |
| 完全性の要求度 | 高 |
| 可用性の要求度 | 中 |
| 調整後の攻撃元区分 | ローカル |
| 調整後の攻撃条件の複雑さ | 高 |
| 調整後の必要な特権レベル | 低 |
| 調整後のユーザ関与レベル | 要 |
| 調整後のスコープ  | 変更なし |
| 調整後の機密性への影響  | 低 |
| 調整後の完全性への影響  | 高 |
| 調整後の可用性への影響  | 低 |
```

などと出力される。

まぁ，個人で使う場合は現状評価基準や環境評価基準の評価はしないけどね。
あれは組織が脆弱性のリスク管理に使うものだし。

なお環境評価基準は v3.0 と v3.1 で評価方法が若干異なるがスコアの算出は v3.1 のロジックで行っている（筈）。
まぁ，今時わざわざ v3.0 を使う組織はないじゃろう。

というわけで個人の手遊びで書くのを躊躇していたのだが，多少でも使っていただいている人がいて PR までいただけるのはありがたい話である（大事なことなので何度でも繰り返すw）。

## ブックマーク

- [CVSS v3.0 Specification Document](https://www.first.org/cvss/v3.0/specification-document)
- [CVSS v3.1 Specification Document](https://www.first.org/cvss/v3.1/specification-document)
- [共通脆弱性評価システムCVSS v3概説：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/vuln/CVSSv3.html)
- [CVSS v3.1]({{< ref "/remark/2020/01/cvss-v3_1.md" >}})

[spiegel-im-spiegel/go-cvss]: https://github.com/spiegel-im-spiegel/go-cvss "spiegel-im-spiegel/go-cvss: Common Vulnerability Scoring System (CVSS)"
