+++
title = "2025年の暦"
date =  "2024-02-03T10:43:46+09:00"
description = "2025年は日食2回と月食2回がある。"
image = "/images/attention/kitten.jpg"
tags = [ "calendar", "astronomy", "ephemeris" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今年も[国立天文台]より翌年（2025年）の暦要項が発表された[^na1]。

[^na1]: [国立天文台]では毎年2月最初の官報で翌年の暦を発表する。

- [令和7（2025）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2024/20240201-rekiyoko.html)
- {{< pdf-file title="令和 7年 (2025) 暦要項" link="https://eco.mtk.nao.ac.jp/koyomi/yoko/pdf/yoko2025.pdf" >}}

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
    start, _ := koyomi.DateFrom("2025-01-01")
    end, _ := koyomi.DateFrom("2025-12-31")
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
| 2025-01-01 | 水 | 元日 |
| 2025-01-13 | 月 | 成人の日 |
| 2025-02-11 | 火 | 建国記念の日 |
| 2025-02-23 | 日 | 天皇誕生日 |
| 2025-02-24 | 月 | 休日 |
| 2025-03-20 | 木 | 春分の日 |
| 2025-04-29 | 火 | 昭和の日 |
| 2025-05-03 | 土 | 憲法記念日 |
| 2025-05-04 | 日 | みどりの日 |
| 2025-05-05 | 月 | こどもの日 |
| 2025-05-06 | 火 | 休日 |
| 2025-07-21 | 月 | 海の日 |
| 2025-08-11 | 月 | 山の日 |
| 2025-09-15 | 月 | 敬老の日 |
| 2025-09-23 | 火 | 秋分の日 |
| 2025-10-13 | 月 | スポーツの日 |
| 2025-11-03 | 月 | 文化の日 |
| 2025-11-23 | 日 | 勤労感謝の日 |
| 2025-11-24 | 月 | 休日 |
```

うーむ。
祝日が土曜日だと損した気分だな。
特に GW のあたり（笑）

## 2025年の暦象

2025年は日食2回と月食2回がある。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2025-03-14 | 金 | 皆既月食 |
| 2025-03-29 | 土 | 部分日食 |
| 2025-09-08 | 月 | 皆既月食 |
| 2025-09-22 | 月 | 部分日食 |

皆既月食はいずれも日本で見えるが 2025-03-14 の皆既月食は西日本が中心のようだ。
2回の部分日食はいずれも日本では見れないそうな。

他にはこんな感じ。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2025-02-03 | 月 | 立春 |
| 2025-03-20 | 木 | 春分 |
| 2025-05-05 | 月 | 立夏 |
| 2025-06-21 | 土 | 夏至 |
| 2025-08-07 | 木 | 立秋 |
| 2025-09-23 | 火 | 秋分 |
| 2025-11-07 | 金 | 立冬 |
| 2025-12-22 | 月 | 冬至 |

ちなみに2025年の「[伝統的七夕](https://www.nao.ac.jp/faq/a0310.html "質問3-10）伝統的七夕について教えて | 国立天文台")」は8月29日である。
って，8月終わるがな。

## ブックマーク

- [カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない]({{< ref "/remark/2019/05/google-ephemeris.md" >}})
- [国立天文台から最新の暦情報を取ってくる【広告記事】](https://zenn.dev/spiegel/articles/20201205-koyomi)

[国立天文台]: https://www.nao.ac.jp/ "国立天文台(NAOJ)"
[暦計算室]: https://eco.mtk.nao.ac.jp/koyomi/ "国立天文台 天文情報センター 暦計算室"
[`github.com/goark/koyomi`]: https://github.com/goark/koyomi "GitHub - goark/koyomi: 日本のこよみ"

## 参考図書

{{% review-paapi "4416623410" %}} <!-- 天文年鑑 2024年版 -->
{{% review-paapi "B0C68VLQMZ" %}} <!-- 卓上カレンダー 2024年版 -->
{{% review-paapi "B07CZZRBGV" %}} <!-- 壁掛けカレンダー 2024年版 -->
