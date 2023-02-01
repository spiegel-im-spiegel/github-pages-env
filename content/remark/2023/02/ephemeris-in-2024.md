+++
title = "2024年の暦"
date =  "2023-02-01T11:06:23+09:00"
description = "2024年は振替休日がいっぱいあるねぇ。"
image = "/images/attention/kitten.jpg"
tags = [ "calendar", "astronomy", "ephemeris" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今年も[国立天文台]より翌年（2024年）の暦要項が発表された[^na1]。

[^na1]: [国立天文台]では毎年2月最初の官報で翌年の暦を発表する。

- [令和6（2024）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2023/20230201-rekiyoko.html)

以降，いくつか抜粋してみる。

## 祝日・休日[^hd1]

[^hd1]: 「国民の祝日」は休日となる。また「「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする」（「国民の祝日に関する法律」より）。

日本の暦情報は[国立天文台]の[暦計算室]で見ることができる。
この中の「[今月のこよみ powered by Google Calendar](https://eco.mtk.nao.ac.jp/koyomi/cande/calendar.html)」で Google Calendar と連携させることができるのだが，この機能を利用した [`github.com/goark/koyomi`] パッケージを公開している。
こんな感じに使える。

```go
package main

import (
	"fmt"
	"os"

	"github.com/goark/koyomi"
)

var weekShortNames = [7]string{"日", "月", "火", "水", "木", "金", "土"}

func WeekShortNameJp(dt koyomi.DateJp) string {
	return weekShortNames[dt.Weekday()%7]
}

func main() {
	start, _ := koyomi.DateFrom("2024-01-01")
	end, _ := koyomi.DateFrom("2024-12-31")
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("| 日付 | 曜日 | 内容 |")
	fmt.Println("| ---- |:----:| ---- |")
	for _, e := range k.Events() {
		fmt.Printf("| %v | %v | %v |\n", e.Date, WeekShortNameJp(e.Date), e.Title)
	}
}
```

これを実行すると以下の出力が得られる。

```text
$ go run main.go 
| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2024-01-01 | 月 | 元日 |
| 2024-01-08 | 月 | 成人の日 |
| 2024-02-11 | 日 | 建国記念の日 |
| 2024-02-12 | 月 | 休日 |
| 2024-02-23 | 金 | 天皇誕生日 |
| 2024-03-20 | 水 | 春分の日 |
| 2024-04-29 | 月 | 昭和の日 |
| 2024-05-03 | 金 | 憲法記念日 |
| 2024-05-04 | 土 | みどりの日 |
| 2024-05-05 | 日 | こどもの日 |
| 2024-05-06 | 月 | 休日 |
| 2024-07-15 | 月 | 海の日 |
| 2024-08-11 | 日 | 山の日 |
| 2024-08-12 | 月 | 休日 |
| 2024-09-16 | 月 | 敬老の日 |
| 2024-09-22 | 日 | 秋分の日 |
| 2024-09-23 | 月 | 休日 |
| 2024-10-14 | 月 | スポーツの日 |
| 2024-11-03 | 日 | 文化の日 |
| 2024-11-04 | 月 | 休日 |
| 2024-11-23 | 土 | 勤労感謝の日 |
```

ふむむー。
2024年は振替休日がいっぱいあるねぇ。
連休バンザイ（笑）

## 2024年の暦象

2024年は日食2回と月食1回がある。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2024-04-09 | 火 | 皆既日食 |
| 2024-09-18 | 水 | 部分月食 |
| 2024-10-03 | 木 | 金環日食 |

しかし，いずれも日本では見れないらしい。

他にはこんな感じ。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2024-02-04 | 日 | 立春 |
| 2024-03-20 | 水 | 春分 |
| 2024-05-05 | 日 | 立夏 |
| 2024-06-21 | 金 | 夏至 |
| 2024-08-07 | 水 | 立秋 |
| 2024-09-22 | 日 | 秋分 |
| 2024-11-07 | 木 | 立冬 |
| 2024-12-21 | 土 | 冬至 |

ちなみに2024年の「[伝統的七夕](https://www.nao.ac.jp/faq/a0310.html "質問3-10）伝統的七夕について教えて | 国立天文台")」は8月10日である。

そうそう。
2024年は閏年っス。

## ブックマーク

- [カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない]({{< ref "/remark/2019/05/google-ephemeris.md" >}})
- [国立天文台から最新の暦情報を取ってくる【広告記事】](https://zenn.dev/spiegel/articles/20201205-koyomi)

[国立天文台]: https://www.nao.ac.jp/ "国立天文台(NAOJ)"
[暦計算室]: https://eco.mtk.nao.ac.jp/koyomi/ "国立天文台 天文情報センター 暦計算室"
[`github.com/goark/koyomi`]: https://github.com/goark/koyomi "GitHub - goark/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4416522940" %}} <!-- 天文年鑑 2023年版 -->
