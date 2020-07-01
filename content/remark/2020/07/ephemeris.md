+++
title = "2020年7月の暦"
date =  "2020-07-01T10:27:55+09:00"
description = "本来なら開催されるはずだったオリンピックのために祝日を無理やり寄せているのがもの悲しい感じ（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

折角 [spiegel-im-spiegel/koyomi] パッケージを作ったので，2020年7月の暦を浚ってみよう。

国立天文台のデータを使っているため，暦象の基準が日本の暦になっている点に注意。

| 日付       | 内容               |
| ---------- | ------------------ |
| 2020-07-01 | 半夏生             |
| 2020-07-01 | 水星が内合         |
| 2020-07-04 | 地球が遠日点通過   |
| 2020-07-05 | 望(満月)           |
| 2020-07-07 | 小暑               |
| 2020-07-10 | 金星が最大光度     |
| 2020-07-12 | 水星が留           |
| 2020-07-13 | 下弦               |
| 2020-07-14 | 木星が衝           |
| 2020-07-19 | 土用の入り         |
| 2020-07-21 | 朔(新月)           |
| 2020-07-21 | 土星が衝           |
| 2020-07-22 | 大暑               |
| 2020-07-23 | 海の日             |
| 2020-07-23 | 水星が西方最大離角 |
| 2020-07-24 | スポーツの日       |
| 2020-07-27 | 上弦               |

本来なら開催されるはずだったオリンピックのために祝日[^hd1] を無理やり寄せているのがもの悲しい感じ（笑）

[^hd1]: 「国民の祝日」は休日となる。また「「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする」（「国民の祝日に関する法律」より）。

## 天象

- 2020-07-05 の望に半影月食がある（日本からは見えない。まぁ，肉眼ではほぼ分からないけど）
- 明けの明星が明るい。おうし座のアルデバランの近くにいる
- 下旬の水星が見頃。明け方の空に注目

各用語の意味は以下の通り。

{{< fig-img src="https://eco.mtk.nao.ac.jp/koyomi/faq/image/phenom_phase.png" title="「天象 - 国立天文台暦計算室」より" link="https://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/phenomena_f.cgi" >}}

「留」とは「[惑星の地心視赤経の時間変化が0、すなわち赤経方向の動きが止まる瞬間](https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FCEB1.html "暦Wiki/惑星/留 - 国立天文台暦計算室")」，もっと簡単に言うと天球上の（見かけの）惑星の動きが反転する瞬間を指す。

## 土用の入り

「土用の入り」は雑節のひとつで，元々は「五行説」の考えから来たものである。

四季を「五行説」に当てはめようとするとひとつ足らなくなるので，各季節の終わりの1/5ずつを集めて土の季節つまり「土用」としたそうな。
無理やりですねぇ（笑）

現在の定義では太陽黄経が 27°, 117°, 207°, 297° となる日が「土用の入り」となる。
ちなみに2020年夏の「土用の丑の日」は 2020-07-21 および 2020-08-02 の2回ある。

## 七夕と五節句

「七夕」を含む五節句は明治の改暦で廃止されて以降，正式な暦としては定義されていない。
ただし国立天文台では旧暦の七夕に近い日を「伝統的七夕」としてキャンペーンを行っている。

- [質問3-10）伝統的七夕について教えて | 国立天文台(NAOJ)](https://www.nao.ac.jp/faq/a0310.html)

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
	start, _ := koyomi.DateFrom("2020-07-01")
	end, _ := koyomi.DateFrom("2020-07-31")
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

気が向いたらシリーズ化する？

## ブックマーク

- [暦Wiki/季節/雑節とは？ - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FBBA8C0E1A4C8A4CFA1A9.html)

- [2020年の主な暦象]({{< ref "/remark/2019/11/astronomical-calendar-2020.md" >}})
- [2020年の祝日休日]({{< ref "/remark/2020/01/2020-holidays.md" >}})
- [日本の暦情報を取得するパッケージを作ってみた]({{< ref "/release/2020/05/koyomi.md" >}})

[spiegel-im-spiegel/koyomi]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4416719485" %}} <!-- 天文年鑑 2020年版 -->
