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
        koyomi.WithCalendarID(koyomi.Holiday),
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
        fmt.Printf("| %v | %v | %v |\n", e.Date.StringJp(), e.Date.WeekdayJp().ShortStringJp(), e.Title)
    }
}
```

これを実行すると以下の出力が得られる。

```text
$ go run main.go
| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2026年1月1日 | 木 | 元日 |
| 2026年1月12日 | 月 | 成人の日 |
| 2026年2月11日 | 水 | 建国記念の日 |
| 2026年2月23日 | 月 | 天皇誕生日 |
| 2026年3月20日 | 金 | 春分の日 |
| 2026年4月29日 | 水 | 昭和の日 |
| 2026年5月3日 | 日 | 憲法記念日 |
| 2026年5月4日 | 月 | みどりの日 |
| 2026年5月5日 | 火 | こどもの日 |
| 2026年5月6日 | 水 | 休日 |
| 2026年7月20日 | 月 | 海の日 |
| 2026年8月11日 | 火 | 山の日 |
| 2026年9月21日 | 月 | 敬老の日 |
| 2026年9月22日 | 火 | 休日 |
| 2026年9月23日 | 水 | 秋分の日 |
| 2026年10月12日 | 月 | スポーツの日 |
| 2026年11月3日 | 火 | 文化の日 |
| 2026年11月23日 | 月 | 勤労感謝の日 |
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
| 2026年2月17日 | 火 | 金環日食 |
| 2026年3月3日 | 火 | 皆既月食 |
| 2026年8月13日 | 木 | 皆既日食 |
| 2026年8月28日 | 金 | 部分月食 |

{{< fig-quote type="markdown" title="令和8（2026）年暦要項の発表" link="https://www.nao.ac.jp/news/topics/2025/20250203-rekiyoko.html" >}}
- 2月17日には金環日食がありますが、日本では見ることができません。
- 3月3日には皆既月食があり、日本では全国で皆既食を見ることができます。
- 8月13日には皆既日食がありますが、日本では見ることができません。
- 8月28日には部分月食がありますが、日本では見ることができません。
{{< /fig-quote >}}

他（二十四節気・雑節）はこんな感じ。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2026年1月5日 | 月 | 小寒 |
| 2026年1月17日 | 土 | 土用の入り |
| 2026年1月20日 | 火 | 大寒 |
| 2026年2月3日 | 火 | 節分 |
| 2026年2月4日 | 水 | 立春 |
| 2026年2月19日 | 木 | 雨水 |
| 2026年3月5日 | 木 | 啓蟄 |
| 2026年3月17日 | 火 | 彼岸の入り |
| 2026年3月20日 | 金 | 春分 |
| 2026年4月5日 | 日 | 清明 |
| 2026年4月17日 | 金 | 土用の入り |
| 2026年4月20日 | 月 | 穀雨 |
| 2026年5月2日 | 土 | 八十八夜 |
| 2026年5月5日 | 火 | 立夏 |
| 2026年5月21日 | 木 | 小満 |
| 2026年6月6日 | 土 | 芒種 |
| 2026年6月11日 | 木 | 入梅 |
| 2026年6月21日 | 日 | 夏至 |
| 2026年7月2日 | 木 | 半夏生 |
| 2026年7月7日 | 火 | 小暑 |
| 2026年7月20日 | 月 | 土用の入り |
| 2026年7月23日 | 木 | 大暑 |
| 2026年8月7日 | 金 | 立秋 |
| 2026年8月23日 | 日 | 処暑 |
| 2026年9月1日 | 火 | 二百十日 |
| 2026年9月7日 | 月 | 白露 |
| 2026年9月20日 | 日 | 彼岸の入り |
| 2026年9月23日 | 水 | 秋分 |
| 2026年10月8日 | 木 | 寒露 |
| 2026年10月20日 | 火 | 土用の入り |
| 2026年10月23日 | 金 | 霜降 |
| 2026年11月7日 | 土 | 立冬 |
| 2026年11月22日 | 日 | 小雪 |
| 2026年12月7日 | 月 | 大雪 |
| 2026年12月22日 | 火 | 冬至 |

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
