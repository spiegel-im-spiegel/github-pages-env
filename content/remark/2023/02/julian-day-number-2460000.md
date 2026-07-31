+++
title = "ユリウス日が2,460,000日を超えた日"
date =  "2023-02-25T11:18:01+09:00"
description = "2023-02-24T21:00:00+09:00 をもちまして，ユリウス日が2,460,000.0日を超えた。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "julian-day-number", "golang", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

以前に[自分で書いて][前の記事]てすっかり忘れていたのだが，昨日 2023-02-24T21:00:00+09:00 をもちまして，ユリウス日が2,460,000.0日を超えた。
単純に1年を365.25日（ユリウス年）と考えても1万日経過するのに27年ほどかかる。

{{< fig-img-quote src="./calculator.png" title="DuckDuckGo" link="https://duckduckgo.com/?q=10000%2F365.25%3D&ia=calculator" lang="en" width="865" >}}

[前の記事]で書いたコードを少し変えて，コマンドライン引数から任意のユリウス日の値を現行暦に変換できるようにしてみた。

```go
package main

import (
    "flag"
    "fmt"

    "github.com/goark/koyomi/jdn"
    "github.com/goark/koyomi/value"
)

func main() {
    day := flag.Float64("d", 0, "Julian day")
    flag.Parse()

    dt := jdn.FromJD(*day)
    fmt.Printf("Julian Day: %.3f\n", *day)
    fmt.Printf("Gregorian: %v (%v)\n", dt, dt.In(value.JST))
}
```

これを使ってユリウス日 2,460,000.0 を計算してみる。

```text
$ go run sample.go -d 2460000.0
Julian Day: 2460000.000
Gregorian: 2023-02-24 12:00:00 +0000 UTC (2023-02-24 21:00:00 +0900 JST)
```

よしよし。
ちゃんと動いてるな。

じゃあ，1万日前の 2,450,000.0 日はいつだったか調べてみよう。

```text
$ go run sample.go -d 2450000.0
Julian Day: 2450000.000
Gregorian: 1995-10-09 12:00:00 +0000 UTC (1995-10-09 21:00:00 +0900 JST)
```

おー。
1995-10-09 かぁ。
次の1万日後はどうかな。

```text
$ go run sample.go -d 2470000.0
Julian Day: 2470000.000
Gregorian: 2050-07-12 12:00:00 +0000 UTC (2050-07-12 21:00:00 +0900 JST)
```

ふむむ。
2050-07-12 か。
その頃は80代半ばだし，流石に生きちゃおらんだろう（笑）

というわけで，小ネタでした。

[前の記事]: {{< ref "/remark/2022/05/julian-day-number.md" >}} "ユリウス日が2,460,000日を超える日"

## 参考図書

{{% review-paapi "4416522940" %}} <!-- 天文年鑑 2023年版 -->
{{% review-paapi "4805202254" %}} <!-- 天体の位置計算 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
