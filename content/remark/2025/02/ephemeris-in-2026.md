+++
title = "2026年の暦"
date =  "2025-02-03T21:01:10+09:00"
description = "2026-09-22 の休日は「国民の祝日に関する法律」第3条第3項のルールによるもの。2026年は日食2回と月食2回がある。"
image = "/images/attention/kitten.jpg"
tags = [ "calendar", "astronomy", "ephemeris" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今年も[国立天文台]より翌年（2026年）の暦要項が発表された[^na1]。

[^na1]: [国立天文台]では毎年2月最初の官報で翌年の暦を発表する。

- [令和8（2026）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2025/20250203-rekiyoko.html)
- [暦要項 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/yoko/index.html)
- {{< pdf-file title="令和 8年 (2026) 暦要項" link="https://eco.mtk.nao.ac.jp/koyomi/yoko/pdf/yoko2026.pdf" >}}

## 祝日・休日

日本の暦情報は[国立天文台]の[暦計算室]で見ることができる。
この中の「[今月のこよみ powered by Google Calendar](https://eco.mtk.nao.ac.jp/koyomi/cande/calendar.html)」で Google Calendar と連携させることができるのだが，この機能を利用した [`github.com/goark/koyomi`] パッケージを公開している。
こんな感じに使える。

```go
package main

import (
    "fmt"
    "os"

    "github.com/goark/koyomi"
    "github.com/goark/koyomi/value"
)

var weekShortNames = [7]string{"日", "月", "火", "水", "木", "金", "土"}

func WeekShortNameJp(dt value.DateJp) string {
    return weekShortNames[dt.Weekday()%7]
}

func main() {
    start, _ := value.DateFrom("2026-01-01")
    end, _ := value.DateFrom("2026-12-31")
    td, err := os.MkdirTemp(os.TempDir(), "blog")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer func() { _ = os.RemoveAll(td) }()
    k, err := koyomi.NewSource(
        koyomi.WithCalendarID(koyomi.Holiday), //祝日・休日
        koyomi.WithStartDate(start),
        koyomi.WithEndDate(end),
        koyomi.WithTempDir(td),
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
| 2026-01-01 | 木 | 元日 |
| 2026-01-12 | 月 | 成人の日 |
| 2026-02-11 | 水 | 建国記念の日 |
| 2026-02-23 | 月 | 天皇誕生日 |
| 2026-03-20 | 金 | 春分の日 |
| 2026-04-29 | 水 | 昭和の日 |
| 2026-05-03 | 日 | 憲法記念日 |
| 2026-05-04 | 月 | みどりの日 |
| 2026-05-05 | 火 | こどもの日 |
| 2026-05-06 | 水 | 休日 |
| 2026-07-20 | 月 | 海の日 |
| 2026-08-11 | 火 | 山の日 |
| 2026-09-21 | 月 | 敬老の日 |
| 2026-09-22 | 火 | 休日 |
| 2026-09-23 | 水 | 秋分の日 |
| 2026-10-12 | 月 | スポーツの日 |
| 2026-11-03 | 火 | 文化の日 |
| 2026-11-23 | 月 | 勤労感謝の日 |
```

2026-09-22 の休日は「[国民の祝日に関する法律](https://laws.e-gov.go.jp/law/323AC1000000178 "国民の祝日に関する法律 | e-Gov 法令検索")」第3条第3項のルールによるもの。

{{< fig-quote type="markdown" title="国民の祝日に関する法律" link="https://laws.e-gov.go.jp/law/323AC1000000178" >}}
第三条　「国民の祝日」は、休日とする。<br>
２　「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする。<br>
３　その前日及び翌日が「国民の祝日」である日（「国民の祝日」でない日に限る。）は、休日とする。
{{< /fig-quote >}}

当然ながら「祭日」については国家は規定しない。

## 2026年の暦象

2026年は日食2回と月食2回がある。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2026-02-17 | 火 | 金環日食 |
| 2026-03-03 | 火 | 皆既月食 |
| 2026-08-13 | 木 | 皆既日食 |
| 2026-08-28 | 金 | 部分月食 |

{{< fig-quote type="markdown" title="令和8（2026）年暦要項の発表" link="https://www.nao.ac.jp/news/topics/2025/20250203-rekiyoko.html" >}}
- 2月17日には金環日食がありますが、日本では見ることができません。
- 3月3日には皆既月食があり、日本では全国で皆既食を見ることができます。
- 8月13日には皆既日食がありますが、日本では見ることができません。
- 8月28日には部分月食がありますが、日本では見ることができません。
{{< /fig-quote >}}

他（二十四節気・雑節）はこんな感じ。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2026-01-05 | 月 | 小寒 |
| 2026-01-17 | 土 | 土用の入り |
| 2026-01-20 | 火 | 大寒 |
| 2026-02-03 | 火 | 節分 |
| 2026-02-04 | 水 | 立春 |
| 2026-02-19 | 木 | 雨水 |
| 2026-03-05 | 木 | 啓蟄 |
| 2026-03-17 | 火 | 彼岸の入り |
| 2026-03-20 | 金 | 春分 |
| 2026-04-05 | 日 | 清明 |
| 2026-04-17 | 金 | 土用の入り |
| 2026-04-20 | 月 | 穀雨 |
| 2026-05-02 | 土 | 八十八夜 |
| 2026-05-05 | 火 | 立夏 |
| 2026-05-21 | 木 | 小満 |
| 2026-06-06 | 土 | 芒種 |
| 2026-06-11 | 木 | 入梅 |
| 2026-06-21 | 日 | 夏至 |
| 2026-07-02 | 木 | 半夏生 |
| 2026-07-07 | 火 | 小暑 |
| 2026-07-20 | 月 | 土用の入り |
| 2026-07-23 | 木 | 大暑 |
| 2026-08-07 | 金 | 立秋 |
| 2026-08-23 | 日 | 処暑 |
| 2026-09-01 | 火 | 二百十日 |
| 2026-09-07 | 月 | 白露 |
| 2026-09-20 | 日 | 彼岸の入り |
| 2026-09-23 | 水 | 秋分 |
| 2026-10-08 | 木 | 寒露 |
| 2026-10-20 | 火 | 土用の入り |
| 2026-10-23 | 金 | 霜降 |
| 2026-11-07 | 土 | 立冬 |
| 2026-11-22 | 日 | 小雪 |
| 2026-12-07 | 月 | 大雪 |
| 2026-12-22 | 火 | 冬至 |

## その他

2026年の「[伝統的七夕](https://www.nao.ac.jp/faq/a0310.html "質問3-10）伝統的七夕について教えて | 国立天文台")」は8月19日である。
お盆過ぎですな。

あと2026年って{{% ruby "ひのえうま" %}}丙午{{% /ruby %}}なんだよな。
昭和時代までのジンクスというか迷信は今でも通用したりするのだろうか。
1966年の丙午のときは（迷信を受けて）出生数が激減したんだよな。
まぁ，当時と違って今は少子化時代だし，迷信を知らない人も多そうだし，関係ないかな。

## ブックマーク

- [国民の祝日と休日 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2009_3.html)

- [冬至に関する与太話]({{< ref "/remark/2017/12/winter-solstice.md" >}})
- [カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない]({{< ref "/remark/2019/05/google-ephemeris.md" >}})
- [お彼岸]({{< ref "/remark/2022/03/ohigan.md" >}})
- [第五の季節：土用]({{< ref "/remark/2022/04/doyo-period.md" >}})
- [節分の狂躁]({{< ref "/remark/2025/02/setsubun.md" >}})

[国立天文台]: https://www.nao.ac.jp/ "国立天文台(NAOJ)"
[暦計算室]: https://eco.mtk.nao.ac.jp/koyomi/ "国立天文台 天文情報センター 暦計算室"
[`github.com/goark/koyomi`]: https://github.com/goark/koyomi "GitHub - goark/koyomi: 日本のこよみ"

## 参考

{{% review-paapi "4416723660" %}} <!-- 天文年鑑 2025年版 -->
{{% review-paapi "4295019313" %}} <!-- 天体写真カレンダー KAGAYA 2025年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->

## 作業中の BGV (メン限配信以外)

- [【#限界飯】節分の日には、恵方巻を食べると聞きました。【一条莉々華/hololive DEV_IS　ReGLOSS】 - YouTube](https://www.youtube.com/watch?v=PdTeJ6cl7WY)
- [【自然界の暗殺者！？】今年の干支・ヘビのヒミツを知ろう【ぴのらぼ】 - YouTube](https://www.youtube.com/watch?v=U3Lvcs3sh8A)
