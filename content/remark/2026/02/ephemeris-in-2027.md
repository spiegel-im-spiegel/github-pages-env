+++
title = "2027年の暦"
date =  "2026-02-02T17:03:15+09:00"
description = "2027年は月食はなし。"
image = "/images/attention/kitten.jpg"
tags = [ "calendar", "astronomy", "ephemeris" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++
今年も[国立天文台]より翌年（2027年）の{{% ruby "れきようこう" %}}暦要項{{% /ruby %}}が発表された[^na1]。

[^na1]: [国立天文台]では毎年2月最初の官報で翌年の暦を発表する。

- [令和9（2027）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2026/20260202-rekiyoko.html)
- [暦要項 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/yoko/index.html)
- {{< pdf-file title="令和9年 (2027) 暦要項" link="https://eco.mtk.nao.ac.jp/koyomi/yoko/pdf/yoko2027.pdf" >}}

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
    start, _ := value.DateFrom("2027-01-01")
    end, _ := value.DateFrom("2027-12-31")
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

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2027年1月1日 | 金 | 元日 |
| 2027年1月11日 | 月 | 成人の日 |
| 2027年2月11日 | 木 | 建国記念の日 |
| 2027年2月23日 | 火 | 天皇誕生日 |
| 2027年3月21日 | 日 | 春分の日 |
| 2027年3月22日 | 月 | 休日 |
| 2027年4月29日 | 木 | 昭和の日 |
| 2027年5月3日 | 月 | 憲法記念日 |
| 2027年5月4日 | 火 | みどりの日 |
| 2027年5月5日 | 水 | こどもの日 |
| 2027年7月19日 | 月 | 海の日 |
| 2027年8月11日 | 水 | 山の日 |
| 2027年9月20日 | 月 | 敬老の日 |
| 2027年9月23日 | 木 | 秋分の日 |
| 2027年10月11日 | 月 | スポーツの日 |
| 2027年11月3日 | 水 | 文化の日 |
| 2027年11月23日 | 火 | 勤労感謝の日 |

「休日」のルールは以下の通り。

{{< fig-quote type="markdown" title="国民の祝日に関する法律" link="https://laws.e-gov.go.jp/law/323AC1000000178" >}}
第三条　「国民の祝日」は、休日とする。<br>
２　「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする。<br>
３　その前日及び翌日が「国民の祝日」である日（「国民の祝日」でない日に限る。）は、休日とする。
{{< /fig-quote >}}

当然ながら「祭日」については国家は規定しない。

## 2027年の暦象

2027年は日食が2回ある。
月食はなし。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2027年2月7日 | 日 | 金環日食 |
| 2027年8月2日 | 月 | 皆既日食 |

{{< fig-quote type="markdown" title="令和9（2027）年暦要項の発表" link="https://www.nao.ac.jp/news/topics/2026/20260202-rekiyoko.html" >}}
- 2月6日から7日にかけて金環日食がありますが、日本では見ることができません。
- 8月2日には皆既日食がありますが、日本では見ることができません。
{{< /fig-quote >}}

他（二十四節気・雑節）はこんな感じ。

| 日付 | 曜日 | 内容 |
| ---- |:----:| ---- |
| 2027年1月5日 | 火 | 小寒 |
| 2027年1月17日 | 日 | 土用の入り |
| 2027年1月20日 | 水 | 大寒 |
| 2027年2月3日 | 水 | 節分 |
| 2027年2月4日 | 木 | 立春 |
| 2027年2月19日 | 金 | 雨水 |
| 2027年3月6日 | 土 | 啓蟄 |
| 2027年3月18日 | 木 | 彼岸の入り |
| 2027年3月21日 | 日 | 春分 |
| 2027年4月5日 | 月 | 清明 |
| 2027年4月17日 | 土 | 土用の入り |
| 2027年4月20日 | 火 | 穀雨 |
| 2027年5月2日 | 日 | 八十八夜 |
| 2027年5月6日 | 木 | 立夏 |
| 2027年5月21日 | 金 | 小満 |
| 2027年6月6日 | 日 | 芒種 |
| 2027年6月11日 | 金 | 入梅 |
| 2027年6月21日 | 月 | 夏至 |
| 2027年7月2日 | 金 | 半夏生 |
| 2027年7月7日 | 水 | 小暑 |
| 2027年7月20日 | 火 | 土用の入り |
| 2027年7月23日 | 金 | 大暑 |
| 2027年8月8日 | 日 | 立秋 |
| 2027年8月23日 | 月 | 処暑 |
| 2027年9月1日 | 水 | 二百十日 |
| 2027年9月8日 | 水 | 白露 |
| 2027年9月20日 | 月 | 彼岸の入り |
| 2027年9月23日 | 木 | 秋分 |
| 2027年10月8日 | 金 | 寒露 |
| 2027年10月21日 | 木 | 土用の入り |
| 2027年10月24日 | 日 | 霜降 |
| 2027年11月8日 | 月 | 立冬 |
| 2027年11月22日 | 月 | 小雪 |
| 2027年12月7日 | 火 | 大雪 |
| 2027年12月22日 | 水 | 冬至 |

## その他

2027年の「[伝統的七夕](https://www.nao.ac.jp/faq/a0310.html "質問3-10）伝統的七夕について教えて | 国立天文台")」は8月8日である。
立秋と同じ日。

## ブックマーク

- [国民の祝日と休日 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2009_3.html)

- [カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない]({{< ref "/remark/2019/05/google-ephemeris.md" >}})
- [暦の改訂（DE430 から DE440 へ）]({{< ref "/remark/2025/05/de440.md" >}})

- [冬至に関する与太話]({{< ref "/remark/2017/12/winter-solstice.md" >}})
- [「新暦七夕」なるものは存在しない]({{< ref "/remark/2019/06/traditional-tanabata.md" >}})
- [立春も動き出す]({{< ref "/remark/2021/01/the-beginning-of-spring.md" >}})
- [お彼岸]({{< ref "/remark/2022/03/ohigan.md" >}})
- [第五の季節：土用]({{< ref "/remark/2022/04/doyo-period.md" >}})
- [春夏秋冬は「四季」ではない？]({{< ref "/remark/2025/03/four-seasons.md" >}})
- [地球が遠日点を通過する話]({{< ref "/remark/2025/06/passage-through-aphelion.md" >}})
- [旧暦の閏月と旧暦2033年問題]({{< ref "/remark/2025/07/leap-month.md" >}})

[国立天文台]: https://www.nao.ac.jp/ "国立天文台(NAOJ)"
[暦計算室]: https://eco.mtk.nao.ac.jp/koyomi/ "国立天文台 天文情報センター 暦計算室"
[`github.com/goark/koyomi`]: https://github.com/goark/koyomi "GitHub - goark/koyomi: 日本のこよみ"

## 参考

{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
