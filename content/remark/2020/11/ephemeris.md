+++
title = "2020年11月の暦"
date =  "2020-11-03T09:40:46+09:00"
description = "『天文年鑑』が出ると年末って感じ。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2020年10月も [spiegel-im-spiegel/koyomi] パッケージを作って暦を浚ってみる。
なお [spiegel-im-spiegel/koyomi] パッケージは v0.1.6 をリリースしている。
[脆弱性チェックの GitHub Action]({{< ref "/remark/2020/10/github-code-scanning-with-golang.md" >}} "Go のコードでも GitHub Code Scanning が使えるらしい") を追加しただけなので，中身は変わらない。

国立天文台のデータを使っているため，暦象の基準が日本の暦になっている点に注意。

| 日付       | 内容               |
| ---------- | ------------------ |
| 2020-11-01 | 天王星が衝         |
| 2020-11-03 | 文化の日           |
| 2020-11-03 | 水星が留           |
| 2020-11-07 | 立冬               |
| 2020-11-08 | 下弦               |
| 2020-11-11 | 水星が西方最大離角 |
| 2020-11-15 | 朔(新月)           |
| 2020-11-16 | 火星が留           |
| 2020-11-22 | 上弦               |
| 2020-11-22 | 小雪               |
| 2020-11-23 | 勤労感謝の日       |
| 2020-11-29 | 海王星が留         |
| 2020-11-30 | 望(満月)           |

## 天象

各用語の意味は以下の通り。

{{< fig-img-quote class="lightmode" src="https://eco.mtk.nao.ac.jp/koyomi/faq/image/phenom_phase.png" title="天象 - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/phenomena_f.cgi" >}}

「留」とは「[惑星の地心視赤経の時間変化が0、すなわち赤経方向の動きが止まる瞬間](https://eco.mtk.nao.ac.jp/koyomi/wiki/CFC7C0B12FCEB1.html "暦Wiki/惑星/留 - 国立天文台暦計算室")」，もっと簡単に言うと天球上の（見かけの）惑星の動きが反転する瞬間を指す。

### しし座流星群

2020-11-17 にしし座流星群が極大になる（予定）。
HR 15 とたくさん流れるわけではないが，しし群は稀に大量出現することがあるので注目しておいてもいいだろう。

新月後で条件もいいし。

### 30日の満月は半影食だが...

2020-11-30 の望（満月）には半影月食がある。
ただし最大食分の時刻が18時半過ぎとかなり早いうえ，半影食は肉眼ではまず分からない。

「そーゆーのがある」くらいの認識でいいだろう（笑）

## 『天文年鑑』が出ると年末って感じ

今年は某パンデミックの影響で出版遅延が多かったが，順調に行けば今月末に[『天文年鑑』の2021年版](https://www.amazon.co.jp/dp/4416620616?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "天文年鑑 2021年版 | 天文年鑑 編集委員会 |本 | 通販 | Amazon")が出る。

{{% review-paapi "4416620616" %}} <!-- 天文年鑑 2021年版 -->

私は既に Amazon に予約発注している。
これが出るといよいよ年末って感じだよねぇ。


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
	start, _ := koyomi.DateFrom("2020-11-01")
	end, _ := koyomi.DateFrom("2020-11-30")
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
