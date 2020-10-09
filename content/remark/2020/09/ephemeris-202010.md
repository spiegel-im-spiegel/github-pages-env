+++
title = "2020年10月の暦"
date =  "2020-09-30T10:49:54+09:00"
description = "どっかの馬鹿な国際スポーツ大会のせいで今年の10月の祝日はなくなりました orz"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2020年10月も [spiegel-im-spiegel/koyomi] パッケージを作って暦を浚ってみる。
なお [spiegel-im-spiegel/koyomi] パッケージは v0.1.5 をリリースしている。
今回も外部パッケージのバージョンを上げただけだけど。

国立天文台のデータを使っているため，暦象の基準が日本の暦になっている点に注意。

| 日付       | 内容               |
| ---------- | ------------------ |
| 2020-10-02 | 望(満月)           |
| 2020-10-02 | 水星が東方最大離角 |
| 2020-10-06 | 火星が地球最近     |
| 2020-10-08 | 寒露               |
| 2020-10-10 | 下弦               |
| 2020-10-14 | 火星が衝           |
| 2020-10-14 | 水星が留           |
| 2020-10-17 | 朔(新月)           |
| 2020-10-20 | 土用の入り         |
| 2020-10-23 | 上弦               |
| 2020-10-23 | 霜降               |
| 2020-10-26 | 水星が内合         |
| 2020-10-31 | 望(満月)           |

そうそう。
どっかの馬鹿な国際スポーツ大会のせいで今年の10月の祝日はなくなりました `orz`

## 天象

各用語の意味は以下の通り。

{{< fig-img-quote class="lightmode" src="https://eco.mtk.nao.ac.jp/koyomi/faq/image/phenom_phase.png" title="天象 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/phenomena_f.cgi" >}}

「留」とは「[惑星の地心視赤経の時間変化が0、すなわち赤経方向の動きが止まる瞬間](https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FCEB1.html "暦Wiki/惑星/留 - 国立天文台暦計算室")」，もっと簡単に言うと天球上の（見かけの）惑星の動きが反転する瞬間を指す。

## 火星が最接近

既に十分見頃だが火星が 2020-10-06 に最接近となる。
つっても中接近くらいだけど。

## 中秋の名月とブルームーン

今年は 2020-10-01 が中秋の名月で，翌2日が望（満月）である。
そんで，10月最終日の日付が変わる直前（23:49）に2回目の望となる。
いわゆる「ブルームーン」である。

実は「月内2回目の満月」をブルームーンと呼ぶのは天文雑誌 “Sky & Telescope” による誤解なのだが，けっこう定着しちゃってる感がある（笑） 元々は “Once in a blue moon” といって「ありえないこと」とか「滅多にないこと」を指す慣用句で，世界各地に逸話や伝説があるらしい。

## 土用の入り

「土用の入り」は雑節のひとつで，元々は「五行説」の考えから来たものである。

四季を「五行説」に当てはめようとするとひとつ足らなくなるので，各季節の終わりの1/5ずつを集めて土の季節つまり「土用」としたそうな。
無理やりですねぇ（笑）

現在の定義では太陽黄経が 27°, 117°, 207°, 297° となる日が「土用の入り」となる。

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
	start, _ := koyomi.DateFrom("2020-10-01")
	end, _ := koyomi.DateFrom("2020-10-31")
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
- [中秋の名月（2020年10月） | 国立天文台(NAOJ)](https://www.nao.ac.jp/astro/sky/2020/10-topics01.html)
- [火星最接近2020 | 国立天文台(NAOJ)](https://www.nao.ac.jp/astro/feature/mars2020/)
    - [火星が地球に最接近（2020年10月） | 国立天文台(NAOJ)](https://www.nao.ac.jp/astro/sky/2020/10-topics02.html)
- [Mars Week 3 - 火星観測週間 2020](http://planetary.jp/marsweek/)

- [2020年の主な暦象]({{< ref "/remark/2019/11/astronomical-calendar-2020.md" >}})
- [日本の暦情報を取得するパッケージを作ってみた]({{< ref "/release/2020/05/koyomi.md" >}})

[spiegel-im-spiegel/koyomi]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4416719485" %}} <!-- 天文年鑑 2020年版 -->
