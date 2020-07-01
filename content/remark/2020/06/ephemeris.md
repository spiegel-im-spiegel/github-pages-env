+++
title = "2020年6月の暦"
date =  "2020-06-01T11:04:00+09:00"
description = "2020-06-21 の金環日食は日本では見れず部分日食となる"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

折角 [spiegel-im-spiegel/koyomi] パッケージを作ったので，2020年6月の暦を浚ってみよう。

国立天文台のデータを使っているため，暦象の基準が日本の暦になっている点に注意。

| 日付       | 内容               |
| ---------- | ------------------ |
| 2020-06-04 | 金星が内合         |
| 2020-06-04 | 水星が東方最大離角 |
| 2020-06-05 | 芒種               |
| 2020-06-06 | 望(満月)           |
| 2020-06-10 | 入梅               |
| 2020-06-13 | 下弦               |
| 2020-06-18 | 水星が留           |
| 2020-06-21 | 朔(新月)           |
| 2020-06-21 | 夏至               |
| 2020-06-21 | 金環日食           |
| 2020-06-24 | 海王星が留         |
| 2020-06-25 | 金星が留           |
| 2020-06-28 | 上弦               |

- 水星が見頃。夕方の空に注目
- 2020-06-06 の望に半影月食がある（肉眼ではほぼ分からないけど）
- 2020-06-21 の金環日食は日本では見れず部分日食となる

各用語の意味は以下の通り。

{{< fig-img src="https://eco.mtk.nao.ac.jp/koyomi/faq/image/phenom_phase.png" title="「天象 - 国立天文台暦計算室」より" link="https://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/phenomena_f.cgi" >}}

ちなみに「留」とは「[惑星の地心視赤経の時間変化が0、すなわち赤経方向の動きが止まる瞬間](https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FCEB1.html "暦Wiki/惑星/留 - 国立天文台暦計算室")」，もっと簡単に言うと天球上の（見かけの）惑星の動きが反転する瞬間を指す。

気が向いたらシリーズ化する？

## コード

今回使用したコードは以下の通り。

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/koyomi"
)

func main() {
    start, _ := koyomi.DateFrom("2020-06-01")
    end, _ := koyomi.DateFrom("2020-06-30")
    k, err := koyomi.NewSource(
        koyomi.WithCalendarID(
            koyomi.Holiday,
            koyomi.MoonPhase,
            koyomi.SolarTerm,
            koyomi.Eclipse,
            koyomi.Planet,
        ),
        koyomi.WithStartDate(start),
        koyomi.WithEndDate(end),
    ).Get()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }

    fmt.Println("| 日付 | 内容 |")
    fmt.Println("| ---- | ---- |")
    for _, e := range k.Events() {
        fmt.Printf("| %v | %v |\n", e.Date, e.Title)
    }
}
```

## ブックマーク

- [日本全国で部分日食（2020年6月） | 国立天文台(NAOJ)](https://www.nao.ac.jp/astro/sky/2020/06-topics03.html)
- [［特設ページ］部分日食（2020年6月21日） | 国立天文台(NAOJ)](https://www.nao.ac.jp/notice/20200621-partial-eclipse.html)
- [APOD: 2020 June 22 - Moon Mountains Magnified during Ring of Fire Eclipse](https://apod.nasa.gov/apod/ap200622.html)
- [APOD: 2020 June 26 - Eclipse under the Bamboo](https://apod.nasa.gov/apod/ap200626.html)
- [APOD: 2020 June 27 - Eclipse under the ISS](https://apod.nasa.gov/apod/ap200627.html)

- [2020年の主な暦象]({{< ref "/remark/2019/11/astronomical-calendar-2020.md" >}})
- [日本の暦情報を取得するパッケージを作ってみた]({{< ref "/release/2020/05/koyomi.md" >}})

[spiegel-im-spiegel/koyomi]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4416719485" %}} <!-- 天文年鑑 2020年版 -->
{{% review-paapi "B07V2Y6B2P" %}} <!-- 日食グラス -->
