+++
title = "WHO COVID-2019 データを取得するパッケージを作ってみた"
date =  "2020-09-08T18:29:42+09:00"
description = "データ取得部分を外だしにしてパッケージ化した。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "infection" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちょっと思いついて WHO COVID-2019 データを取得するパッケージを作ってみた。

- [spiegel-im-spiegel/cov19data: Importing WHO COVID-2019 Cases Global Data](https://github.com/spiegel-im-spiegel/cov19data)

ろくすっぽテストもしてないが，汎用で使えるとも思えないし，まぁいいだろう。

“[WHO Coronavirus Disease (COVID-19) Dashboard](https://covid19.who.int/)” で公開されている CSV データを取ってくるだけの簡単なお仕事である。

それだけではナニなので，簡単な集計もできるようにした。
たとえば

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/cov19data"
    "github.com/spiegel-im-spiegel/cov19data/client"
    "github.com/spiegel-im-spiegel/cov19data/entity"
    "github.com/spiegel-im-spiegel/cov19data/histogram"
    "github.com/spiegel-im-spiegel/cov19data/values"
)

func main() {
    h, err := cov19data.MakeHistogramWHO(
        client.Default(),
        values.NewPeriod(
            values.NewDate(2020, time.Month(8), 1),
            values.Yesterday(),
        ),
        7, //step by 7 days
        entity.WithCountryCode(values.CC_JP),
        entity.WithRegionCode(values.WPRO),
    )
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }

    b, err := histogram.ExportHistCSV(h)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
        fmt.Println()
    }
    // Output:
    // Date_from,Date_to,Cases,Deaths
    // 2020-07-28,2020-08-03,8698,16
    // 2020-08-04,2020-08-10,9303,35
    // 2020-08-11,2020-08-17,7677,52
    // 2020-08-18,2020-08-24,6840,82
    // 2020-08-25,2020-08-31,5358,98
    // 2020-09-01,2020-09-07,3991,84
}
```

てな感じに書けば，8月以降の日本のデータを7日ごとに集計できたりする。

このパッケージで集計したデータを使えば

{{< fig-img src="./covid-2019-new-cases-histgram-in-japan.png" title="Confirmed COVID-2019 Cases in Japan" link="./covid-2019-new-cases-histgram-in-japan.png" width="747" >}}

なんてなグラフも描ける。

元々は SARS-CoV-2 ウイルス感染の動向をチェックするために，ネットの隅っこで

- [spiegel-im-spiegel/covid-2019-report: 日本における COVID-2019 確認発症者のレポート](https://github.com/spiegel-im-spiegel/covid-2019-report)

を作って手作業でチマチマ集計していたのだが，徐々に自動化できるようになったので，データ取得部分を外だしにしてパッケージ化したというわけ。

「とりあえず動いてるからいいか」とあまりテストを書いてないし，設計も雑で如何にもやっつけなコードだけど，よろしかったらどうぞ。

## ブックマーク

- [Go 言語でもグラフを描きたい！]({{< ref "/golang/chart-with-golang.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
