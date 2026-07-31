+++
title = "2020年9月の暦"
date =  "2020-09-01T22:54:19+09:00"
description = "来月は火星最接近。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2020年9月も [spiegel-im-spiegel/koyomi] パッケージを作って暦を浚ってみる。
なお [spiegel-im-spiegel/koyomi] パッケージは v0.1.4 をリリースしている。
まぁ，外部パッケージのバージョンを上げただけだけど。

国立天文台のデータを使っているため，暦象の基準が日本の暦になっている点に注意。

| 日付       | 内容       |
| ---------- | ---------- |
| 2020-09-02 | 望(満月)   |
| 2020-09-07 | 白露       |
| 2020-09-10 | 下弦       |
| 2020-09-10 | 火星が留   |
| 2020-09-12 | 海王星が衝 |
| 2020-09-13 | 木星が留   |
| 2020-09-17 | 朔(新月)   |
| 2020-09-19 | 彼岸の入り |
| 2020-09-21 | 敬老の日   |
| 2020-09-22 | 秋分の日   |
| 2020-09-22 | 秋分       |
| 2020-09-24 | 上弦       |
| 2020-09-29 | 土星が留   |

## 天象

各用語の意味は以下の通り。

{{< fig-img-quote class="lightmode" src="https://eco.mtk.nao.ac.jp/koyomi/faq/image/phenom_phase.png" title="天象 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/phenomena_f.cgi" >}}

「留」とは「[惑星の地心視赤経の時間変化が0、すなわち赤経方向の動きが止まる瞬間](https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FCEB1.html "暦Wiki/惑星/留 - 国立天文台暦計算室")」，もっと簡単に言うと天球上の（見かけの）惑星の動きが反転する瞬間を指す。

### 火星が見頃（10月に最接近）

地球から見て火星は2年2ヶ月毎に接近する。
今年は火星が 2020-10-06 が最接近日となる。
つっても中接近くらいだけど。

8月くらいから見頃になっていて，あちこちの天文系サイトで映像等が上がっている。
探してみるのも面白いだろう。

## 秋分とお彼岸

現在の定義では「秋分の日」の定義は「太陽黄経が180°になる瞬間を含む日」である。

{{< fig-img-quote class="lightmode" src="https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FC6F3BDBDBBCDC0E1B5A4A4CEC4EAA4E1CAFD24C0E1B5A4B5B0C6BBBFDE2.png" title="暦Wiki/季節/二十四節気の定め方 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FC6F3BDBDBBCDC0E1B5A4A4CEC4EAA4E1CAFD.html" >}}

そして秋分の日を挟む前後3日が「お彼岸」となる。
ちなみに今年の「彼岸の入り」は 2020-09-22 である。

実は最近まで知らなかったのだが，春分・秋分を「彼岸の中日」と定めたのは天保暦以後のことらしい。
天保暦以前はこんな感じ。

{{< fig-quote class="nobox smaller center" title="暦Wiki/季節/雑節とは？ - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/wiki/B5A8C0E12FBBA8C0E1A4C8A4CFA1A9.html" >}}
<table>
<tr>
    <th>暦法</th>
    <td>-5</td>
    <td>-4</td>
    <td>-3</td>
    <td>-2</td>
    <td>-1</td>
    <td>春秋分</td>
    <td>+1</td>
    <td>+2</td>
    <td>+3</td>
    <td>+4</td>
    <td>+5</td>
    <td>+6</td>
    <td>+7</td>
    <td>+8</td>
</tr><tr style="vertical-align:middle;">
    <th>宣明暦<br>貞享暦</th>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">入り</td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">中日</td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">明け</td>
</tr><tr>
    <th rowspan="2">宝暦暦<br>寛政暦</th>
    <td style="vertical-align:middle;">入り</td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">中日</td>
    <td></td>
    <td style="vertical-align:middle;">（春分）</td>
    <td style="vertical-align:middle;">明け</td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
</tr><tr>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">入り</td>
    <td style="vertical-align:middle;">（秋分）</td>
    <td></td>
    <td style="vertical-align:middle;">中日</td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">明け</td>
    <td></td>
    <td></td>
    <td></td>
</tr><tr>
    <th>天保暦<br>現在</th>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">入り</td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">中日</td>
    <td></td>
    <td></td>
    <td style="vertical-align:middle;">明け</td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
    <td></td>
</tr>
</table>
{{< /fig-quote >}}

宣明暦・貞享暦は「どうしてそうなった」って感じ（笑）

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
	start, _ := koyomi.DateFrom("2020-09-01")
	end, _ := koyomi.DateFrom("2020-09-30")
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
