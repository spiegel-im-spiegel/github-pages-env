+++
title = "go-cvss パッケージ v1.4.2 をリリースした"
date =  "2023-01-28T12:47:48+09:00"
description = "さらに今回は fuzzing テストまでしてもらって，ホンマに「マジすんません」って感じである。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "package", "module", "golang", "cvss" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

CVSS ベクタ文字列を可視化したいという軽い動機で作った，拙作の [`github.com/goark/go-cvss`] パッケージだが，微妙に使って頂いてるようで，バグ報告をいくつか頂いたため，修正版をリリースした。

- [Release v1.4.2 · goark/go-cvss · GitHub](https://github.com/goark/go-cvss/releases/tag/v1.4.2)

CVSS のベクタ文字列（`CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H`[^cve1] など）のバリデーションを厳密に行うようにした。
今までベクタ文字列の parse はかなり緩くしていて， metric 名が重複してても（後勝ちで）有効にしてたし，大文字小文字も関係なく有効にしていたが，仕様的にあかんやろ，ということで。
これに伴い，古いコードは drop した（コード管理が煩雑になるので）。

[^cve1]: [CVE-2022-3515](https://nvd.nist.gov/vuln/detail/CVE-2022-3515) より

使い方は今までと変わらず

```go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/goark/go-cvss/v3/metric"
	"github.com/goark/go-cvss/v3/report"
	"golang.org/x/text/language"
)

var template = "- `{{ .Vector }}`" + `
- {{ .SeverityName }}: {{ .SeverityValue }} (Score: {{ .BaseScore }})

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
`

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Set CVSS vector")
		return
	}
	bm, err := metric.NewBase().Decode(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	r, err := report.NewBase(bm, report.WithOptionsLanguage(language.Japanese)).ExportWithString(template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
}
```

などとしておけば

```text
$ go run main.go "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H"
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
```

と出力される。

この [`github.com/goark/go-cvss`] パッケージってコードの半分くらい貰いものだし，さらに今回は fuzzing テストまでしてもらって，ホンマに「マジすんません」って感じである。ありがたや {{< emoji "ペコン" >}}

でも CVSS ってあくまでもリスクの「評価基準」のひとつであって，そこから「どうする」ってのはまた別の話なんだよね。
個人なら CVSS の Base metrics 情報を見て都度判断すればいいけど，組織では [SSVC (Stakeholder-Specific Vulnerability Categorization)](https://resources.sei.cmu.edu/library/asset-view.cfm?assetid=653459 "Prioritizing Vulnerability Response: A Stakeholder-Specific Vulnerability Categorization (Version 2.0)") なんかと組み合わせる必要があるかもしれない。

道具は適材適所で使いましょう，ということで。

## 【2022-01-29 追記】 v1.4.4 をリリースした

またバグ報告があったので修正版をリリースした。
とほほ...

- [Release v1.4.4 · goark/go-cvss · GitHub](https://github.com/goark/go-cvss/releases/tag/v1.4.4)

今回のついでにスコアの計算周りのリファクタリングを行った。
ちょっとスッキリ！

## ブックマーク

- [CVSS v3.0 Specification Document](https://www.first.org/cvss/v3.0/specification-document)
- [CVSS v3.1 Specification Document](https://www.first.org/cvss/v3.1/specification-document)
- [共通脆弱性評価システムCVSS v3概説：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/vuln/CVSSv3.html)

- [オープンソース製品とソフトウェア部品表]({{< ref "/remark/2022/08/software-bills-of-materials.md" >}})

[`github.com/goark/go-cvss`]: https://github.com/goark/go-cvss "goark/go-cvss: Common Vulnerability Scoring System (CVSS)"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
