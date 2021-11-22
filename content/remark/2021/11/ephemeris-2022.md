+++
title = "2021年の主な暦象"
date =  "2021-11-22T20:56:39+09:00"
description = "2022年は火星大接近の年。皆既月食もあるよ"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++


さて，年末と言えば『[天文年鑑](https://www.amazon.co.jp/dp/441662140X?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』ですよ。
今年は「[天文月報](https://www.asj.or.jp/jp/activities/geppou/ "天文月報 - 公益社団法人 日本天文学会")」と一緒に来ました（笑）

ホンマ，これが来ると年末って感じだよねー。

{{< fig-img src="./51697049566_9e2bddf5af_e.jpg" title="両方きてた。天文年鑑が来ると年末って感じ | Flickr" link="https://www.flickr.com/photos/spiegel/51697049566">}}

とりあえず流し読みしますかね。

## 日食・月食

2022年は部分日食が2回，皆既月食が2回ある。
日食・月食の情報は[国立天文台からも取得]({{< ref "/release/2020/05/koyomi.md" >}} "日本の暦情報を取得するパッケージを作ってみた")できる。
こんな感じ。

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/koyomi"
)

func main() {
    start, _ := koyomi.DateFrom("2022-01-01")
    end, _ := koyomi.DateFrom("2022-12-31")
    k, err := koyomi.NewSource(
        koyomi.WithCalendarID(koyomi.Eclipse),
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

これの実行結果から以下を取得してみた。

| 日付 | 内容 |
| ---- | ---- |
| 2022-05-01 | 部分日食 |
| 2022-05-16 | 皆既月食 |
| 2022-10-25 | 部分日食 |
| 2022-11-08 | 皆既月食 |

このうち日本で見られるのは 2022-11-08 の皆既月食のみ。
食の開始が18時過ぎなので，概ね全国で見れるだろう（晴れればね）。

## 惑星関連

火星は2年2ヶ月ごとに大接近するが，2022年12月がこれにあたる。
これから観測やら盛んになるであろう。

2022年に日本で観測可能な惑星食は 2022-05-27 の昼間に起こる金星食と 2021-07-21 から 2021-07-22 にかけて起こる火星食がある。
金星食は南西諸島でのみ観測可能だが火星食の方は山口県以東であれば見れるらしい。
ただし東京より西では出現のみ観測可能。

## 彗星

C/2021 O3 PANSTARRS 彗星が 2022-04-20 に近日点を通過するのだが，肉眼等級まで明るくなるのでは？ と期待されている。

## 流星群

2022年8月のペルセウス座流星群は満月の翌日でマジに条件が悪い。
2022年12月のふたご群も下弦の月よりも前なので微妙に条件が悪い。
2022年では正月早々のしぶんぎ群が新月の翌日で条件がいいが，ほろ酔い気分じゃ難しいか？

### その他

- 春分の日は3月21日（月），秋分の日は9月23日（金）である。更に夏至は6月21日，冬至は12月22日となる
- 立春は2月4日，立夏は5月5日，立秋は8月7日，立冬は11月7日
- 「[伝統的七夕](https://www.nao.ac.jp/faq/a0310.html "質問3-10）伝統的七夕について教えて | 国立天文台")」は8月4日
- 中秋の名月は9月10日で望（満月）と重なる

## 参考図書

- [立春も動き出す]({{< ref "/remark/2021/01/the-beginning-of-spring.md" >}}) : 2022年はふつうに 2月4日 です
- [2022年の暦]({{< ref "/remark/2021/02/ephemeris-in-2022.md" >}})

## 参考図書

{{% review-paapi "441662140X" %}} <!-- 天文年鑑 2022年版 -->
