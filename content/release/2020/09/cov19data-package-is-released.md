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
    "github.com/spiegel-im-spiegel/cov19data/filter"
    "github.com/spiegel-im-spiegel/cov19data/histogram"
    "github.com/spiegel-im-spiegel/cov19data/values"
    "github.com/spiegel-im-spiegel/errs"
)

func getHist() ([]*histogram.HistData, error) {
    impt, err := cov19data.NewWeb(client.Default())
    if err != nil {
        return nil, errs.Wrap(err)
    }
    defer impt.Close()
    return impt.Histogram(
        values.NewPeriod(
            values.NewDate(2020, time.Month(9), 1),
            values.NewDate(2020, time.Month(9), 28),
        ),
        7,
        filter.WithCountryCode(values.CC_JP),
        filter.WithRegionCode(values.WPRO),
    )
}

func main() {
    hist, err := getHist()
    if err != nil {
        fmt.Fprintf(os.Stderr, "%+v\n", err)
        return
    }
    b, err := histogram.ExportCSV(hist)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        return
    }
    if _, err := io.Copy(os.Stdout, bytes.NewReader(b)); err != nil {
        fmt.Println(err)
    }
    // Output:
    // Date_from,Date_to,Cases,Deaths
    // 2020-09-01,2020-09-07,3991,84
    // 2020-09-08,2020-09-14,3801,79
    // 2020-09-15,2020-09-21,3483,58
    // 2020-09-22,2020-09-28,2991,48
}
```

てな感じに書けば，9月の日本のデータを7日ごとに集計できたりする。

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

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
