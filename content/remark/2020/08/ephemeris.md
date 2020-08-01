+++
title = "2020年8月の暦"
date =  "2020-08-01T19:37:02+09:00"
description = "今年の伝統的七夕は 2020-08-25"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2020年8月も [spiegel-im-spiegel/koyomi] パッケージを作って暦を浚ってみる。
国立天文台のデータを使っているため，暦象の基準が日本の暦になっている点に注意。

| 日付       | 内容               |
| ---------- | ------------------ |
| 2020-08-04 | 望(満月)           |
| 2020-08-07 | 立秋               |
| 2020-08-10 | 山の日             |
| 2020-08-12 | 下弦               |
| 2020-08-13 | 金星が西方最大離角 |
| 2020-08-16 | 天王星が留         |
| 2020-08-18 | 水星が外合         |
| 2020-08-19 | 朔(新月)           |
| 2020-08-23 | 処暑               |
| 2020-08-26 | 上弦               |
| 2020-08-31 | 二百十日           |

## 天象

各用語の意味は以下の通り。

{{< fig-img src="https://eco.mtk.nao.ac.jp/koyomi/faq/image/phenom_phase.png" title="「天象 - 国立天文台暦計算室」より" link="https://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/phenomena_f.cgi" >}}

「留」とは「[惑星の地心視赤経の時間変化が0、すなわち赤経方向の動きが止まる瞬間](https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FCEB1.html "暦Wiki/惑星/留 - 国立天文台暦計算室")」，もっと簡単に言うと天球上の（見かけの）惑星の動きが反転する瞬間を指す。

## ペルセウス座流星群

2020年のペルセウス座流星群は8月12日22時頃に極大になる予定。

夏のペルセウス座流星群はHR50（目視で1時間に50前後流れる）くらいで毎年安定的に流れるのが特徴である。
輻射点のあるペルセウス座が昇るタイミングで夜半すぎからが本来の見頃なのだが，今年は下弦の月と重なっているため条件が悪い。
夜半までの時間でじっくり観望・観測するのがいいだろう。

- [2020年のペルセウス座流星群の情報](http://meteor.kaicho.net/per2020.html)

せっかくの夏休み。
今年は思いっきり外で遊べず大変だろうが，お子さんのいる方は是非この機会に巻き込んでみてはいかがだろう。

## まだ見える NEOWISE 彗星

今年発見され話題になっている [NEOWISE 彗星]({{< ref "/remark/2020/07/comet-neowise.md" >}} "C/2020 F3 (NEOWISE) が見頃")は宵の空でまだしばらくは見れるらしい。

## 伝統的七夕

[先月も書いた]({{< ref "/remark/2020/07/ephemeris.md" >}} "2020年7月の暦")が
「七夕」を含む五節句は明治の改暦で廃止されて以降，正式な暦としては定義されていない。
その代わりと言ってはナニだが，国立天文台では旧暦の七夕に近い日を「伝統的七夕」としてキャンペーンを行っている。

「伝統的七夕」の定義は以下の通り。

{{< fig-quote type="markdown" title="質問3-10）伝統的七夕について教えて | 国立天文台(NAOJ)" link="https://www.nao.ac.jp/faq/a0310.html" >}}
{{% quote %}}二十四節気の処暑（しょしょ＝太陽黄経が150度になる瞬間）を含む日かそれよりも前で、処暑に最も近い朔（さく＝新月）の瞬間を含む日から数えて7日目が「伝統的七夕」の日です{{% /quote %}}
{{< /fig-quote >}}

この定義に従うと，2020年の伝統的七夕は 2020-08-19 の朔から6日後の 2020-08-25 となる。

## 今年は「土用の丑の日」がもう1回ある

[先月も書いた]({{< ref "/remark/2020/07/ephemeris.md" >}} "2020年7月の暦")とおり，こよみ上の秋の直前の「土用」は，今年は 2020-07-19 から立秋前日の 2020-08-07 までだが，この間の「土用の丑の日」は 2020-07-21 および 2020-08-02 の2回ある。 

ちなみに私はウナギよりアナゴが好きです（笑）

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
    start, _ := koyomi.DateFrom("2020-08-01")
    end, _ := koyomi.DateFrom("2020-08-31")
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

- [暦Wiki/季節/雑節とは？ - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FBBA8C0E1A4C8A4CFA1A9.html)

- [2020年の主な暦象]({{< ref "/remark/2019/11/astronomical-calendar-2020.md" >}})
- [日本の暦情報を取得するパッケージを作ってみた]({{< ref "/release/2020/05/koyomi.md" >}})

[spiegel-im-spiegel/koyomi]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4416719485" %}} <!-- 天文年鑑 2020年版 -->
