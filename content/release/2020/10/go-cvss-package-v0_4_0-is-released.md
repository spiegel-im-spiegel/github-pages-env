+++
title = "えんやらやっと go-cvss パッケージ v0.4.0 をリリースした"
date =  "2020-10-07T16:48:21+09:00"
description = "PR もらっていったんリリースしたのだが，アレなコードでマジすんません 🙇‍♂️ ということで書き直した。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "package", "module", "golang", "cvss" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

絶賛放置プレイ中の [spiegel-im-spiegel/go-cvss] パッケージだが，どうも使っていただいてる方がいて「Temporal Metrics は実装しないの？」と言われ「PR くれるならマージするよ」と下手くそな英語で返したのだが，本当に PR をくださって恐縮です。

で，まぁ，ありがたくマージしていったん [v0.3.0 をリリース](https://github.com/spiegel-im-spiegel/go-cvss/releases/tag/v0.3.0)したのだが，なんせベースのコードが場当たりなやっつけコードなので，マジすんません 🙇‍♂️ て感じ。

そこで，一念発起してちゃんと書き直すことにした。
もっとも，コード自体はほとんど使いまわしで構成を変えただけなんだけどね。

というわけで v0.4.0 をリリースしました。

- [Release v0.4.0 · spiegel-im-spiegel/go-cvss · GitHub](https://github.com/spiegel-im-spiegel/go-cvss/releases/tag/v0.4.0)

以前のコードは残してあるが deprecated なコードとしていずれ削除する予定。
書き直したコードはこんな風に使う。

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/go-cvss/v3/metric"
)

func main() {
    bm, err := metric.NewBase().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H") //CVE-2020-1472: ZeroLogon
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("Severity: %v (%v)\n", bm.Severity(), bm.Score())
    // Output:
    // Severity: Critical (10)
}
```

`metric.NewBase()` の部分を `metric.NewTemporal()` に変えれば Temporal Metrics として処理できる。

```go {hl_lines=[2, "7-8"]}
func main() {
    tm, err := metric.NewTemporal().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H/E:F/RL:W/RC:R") //CVE-2020-1472: ZeroLogon
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("Base Severity: %v (%v)\n", tm.BaseMetrics().Severity(), tm.BaseMetrics().Score())
    fmt.Printf("Temporal Severity: %v (%v)\n", tm.Severity(), tm.Score())
    // Output:
    // Base Severity: Critical (10)
    // Temporal Severity: Critical (9.1)
}
```

評価の内容をテンプレートを使って吐き出す処理は別のサブパッケージに分離した（上の PR くれた方はこの機能を全く使ってないみたいなので）。

たとえば

```go
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
```

みたいに markdown の表形式でテンプレートを作って

```go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/go-cvss/v3/metric"
	"github.com/spiegel-im-spiegel/go-cvss/v3/report"
	"golang.org/x/text/language"
)

var template = `
...
`

func main() {
	bm, err := metric.NewBase().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H") //CVE-2020-1472: ZeroLogon
	if err != nil {
        fmt.Fprintln(os.Stderr, err)
		return
	}
	r, err := report.NewBase(bm, report.WithOptionsLanguage(language.Japanese)).ExportWithString(template)
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
$ go run main.go
| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更あり |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |
```

のように出力してくれる（ちなみに英語と日本語にしか対応してません`w`）。

さて。
Environmental Metrics に対応するかどうか。
やるなら今のうちなんだよなぁ...

## ブックマーク

- [CVSS v3.0 Specification Document](https://www.first.org/cvss/v3.0/specification-document)
- [CVSS v3.1 Specification Document](https://www.first.org/cvss/v3.1/specification-document)
- [共通脆弱性評価システムCVSS v3概説：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/vuln/CVSSv3.html)
- [CVSS v3.1]({{< ref "/remark/2020/01/cvss-v3_1.md" >}})

[spiegel-im-spiegel/go-cvss]: https://github.com/spiegel-im-spiegel/go-cvss "spiegel-im-spiegel/go-cvss: Common Vulnerability Scoring System (CVSS)"
<!-- eof -->
