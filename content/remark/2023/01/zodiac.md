+++
title = "今年は卯年らしい"
date =  "2023-01-02T12:57:11+09:00"
description = "今年は色々と仕切り直しの年になりそうです。"
image = "/remark/2023/01/zodiac/16631382170_7f4baf9a2b_e.jpg"
tags = [ "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

改めまして，明けましておめでとうございます。
[昨年末のゴタゴタ]({{< ref "/remark/2022/12/heart-attack.md" >}} "ハライタだと思った？ 残念！ 心筋梗塞でした")のおかげで今年は色々と仕切り直しの年になりそうです。

## 今年は卯年らしい

えーっと，2023年の干支は...

```go
package main

import (
    "fmt"

    "github.com/goark/koyomi/zodiac"
)

func main() {
    fmt.Println(zodiac.ZodiacYearNumber(2023))
}
```

これを[実行](https://go.dev/play/p/dlV4qiTFMjR "Go Playground - The Go Programming Language")すると

```text
癸 卯
```

ということで，今年は卯年らしい。
まぁ，正月過ぎたら次の年末年始まで忘れるんだけどね（笑）

## ユリウス日が2,460,000日を超える日（再掲載）

[昨年やったネタ]({{< ref "/remark/2022/05/julian-day-number.md" >}})だが，2023年はユリウス日が2,460,000日を超える。

```go
package main

import (
    "fmt"

    "github.com/goark/koyomi"
    "github.com/goark/koyomi/jdn"
)

func main() {
    dt := jdn.FromJDN(2460000)
    fmt.Printf("%v (%v)\n", dt, dt.In(koyomi.JST))
}
```

これを[実行](https://go.dev/play/p/PqTiW1UWGls "Go Playground - The Go Programming Language")すると

```text
2023-02-24 12:00:00 +0000 UTC (2023-02-24 21:00:00 +0900 JST)
```

というわけで，日本時間の 2023-02-24 の21時で2,460,000日を超える[^jd1]。

[^jd1]: ユリウス日は UT または UTC で計算し（地球時（TT）による拡張もある），正午（12時）が1日の起点となる。

## 今年は明治の改暦から150年

明治の改暦が1873年なので，2023年の今年はちょうど150周年ということになる。
イベントとかあるのかな？

ちなみに1973年の100周年のときはこんな本が出ていた。

{{< fig-img src="./52600868047_0c4e183009_e.jpg" title="1973年、明治の改暦100周年記念で出版された本 | Flickr" link="https://www.flickr.com/photos/spiegel/52600868047/">}}

この本をネタに記事を書いたことがある。

- [「暦」日本史 （再掲載）]({{< ref "/remark/2015/japanese-koyomi.md" >}})

### 「明治の改暦」関連ブックマーク

- [暦Wiki/歴史/日本の暦/6.明治維新と太陽暦 - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/wiki/CEF2BBCB2FC6FCCBDCA4CECEF12F6.CCC0BCA3B0DDBFB7A4C8C2C0CDDBCEF1.html)

## 今年は卯年らしいので

2015年に出雲大社で撮った写真など。

{{< fig-img src="./16631382170_7f4baf9a2b_e.jpg" title="ご注文はうさぎですか | Flickr" link="https://www.flickr.com/photos/spiegel/16631382170/">}}

これを見たときは「なんじゃこら」と思ったものだが，なんでも撮っておくものである（笑）

Facebook で “flying rabbit snowman morning in Studio Ghibli style” で AI に絵を描かせた人がいて，私も便乗して同じキーワードで [DALL·E](https://labs.openai.com/) に描かせてみたのがこんな感じ。

{{< fig-img src="./flying-rabbit-snowman.png" title="flying rabbit snowman morning in Studio Ghibli style" link="./flying-rabbit-snowman.png" width="1024" >}}

ジブリっぽくないな（笑）

## ...というわけで

今年も「頑張ること」を頑張らずに，できる範囲でほどほどに頑張ることにします。
しかし，昨年末から年末年始休暇も含めて3週間近く仕事を休んでいるのだが，ちゃんと社会復帰できるのだろうか。
ふむむむ。

## 参考

{{% review-paapi "4416522940" %}} <!-- 天文年鑑 2023年版 -->
{{% review-paapi "B01FI3DQ9K" %}} <!-- ノーポイッ! -->
{{% review-paapi "B01FI3BM00" %}} <!-- ときめきポポロン♪ -->
