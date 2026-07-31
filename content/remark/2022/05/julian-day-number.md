+++
title = "ユリウス日が2,460,000日を超える日"
date =  "2022-05-02T20:38:59+09:00"
description = "2023年にはユリウス日が2,460,000日を超えるそうな。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "julian-day-number", "golang", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

国立天文台から面白い記事が出ている。

- [ユリウス日について - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2023_1.html)

この記事によると

{{< fig-quote type="markdown" title="ユリウス日について" link="https://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2023_1.html" >}}
{{% quote %}}令和5年(2023)にはついに246 ****日となる{{% /quote %}}
{{< /fig-quote >}}

のだそうだ。
また，この記事にグレゴリオ暦とユリウス日を相互変換するアルゴリズムも書かれている（ちなみにこのアルゴリズムは『天文年鑑』にも載っている）。

実は Go 言語を覚え始めた頃に練習で[ユリウス日を求める処理を書いたことがある]({{< ref "/golang/julian-day-number.md" >}})のだが，アレとは若干アルゴリズムが違うので，拙作 [`github.com/goark/koyomi`][goark/koyomi] のサブパッケージとして追加してみた。
洒落で有理数を扱える [`big`][math/big]`.Rat` 型を使っていることを除けば大したアルゴリズムではないので，中身については割愛する[^jd1]。
興味のある方は[リポジトリ][goark/koyomi]をご覧ください。

[^jd1]: ユリウス日の計算程度であれば `float64` を使っても全く問題ない。ユリウス日の起点は紀元前4713年1月1日正午（12時）とされているが，もちろんその時代に（グレゴリオ暦はおろか）ユリウス暦があったわけではないし時刻系も今とは異なる。あくまでも現在の暦・時刻系の（過去方向への）延長線上にある仮想的起点と捉えるべきだろう。ユリウス日は日付（整数部）のみを扱う場合は Julian Day Number，時刻（小数部）を含める場合は Julian Date と呼び分けることが多い。今はあまり聞かないが Julian Day Number を日本語でユリウス通日と言う人もいる。時刻を含める場合は UT または UTC で計算する。ちなみに座表時系のひとつである TT (Terrestrial Time; 地球時) への拡張もあるらしい。

というわけで，記事に書かれているアルゴリズムを使ってユリウス日が2,460,000日を超えるのはいつか調べてみる。

```go
package main

import (
    "fmt"

    "github.com/goark/koyomi"
    "github.com/goark/koyomi/jdn"
)

func main() {
    num := 2460000.0
    dt := jdn.FromJD(num)
    mjd := jdn.GetMJD(dt)
    fmt.Printf("Julian Day Number: %.3f (%v)\n", num, mjd.FloatString(3))
    fmt.Printf("Gregorian: %v (%v)\n", dt, dt.In(koyomi.JST))
}
```

これを実行すると

```text
$ go run sample.go 
Julian Day Number: 2460000.000 (59999.500)
Gregorian: 2023-02-24 12:00:00 +0000 UTC (2023-02-24 21:00:00 +0900 JST)
```

となる。
国立天文台のページで検算してみると

{{< fig-img-quote src="./julian-date.png" title="ユリウス日 - 国立天文台暦計算室" link="http://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/date2jd.cgi" width="1337" >}}

と同じ値が出たので，たぶん問題ないだろう[^jd2]。

[^jd2]: 拙作の [`github.com/goark/koyomi`][goark/koyomi]`/jdn` パッケージではグレゴリオ暦しか対応していない。これは [`time`][time]`.Time` 型がグレゴリオ暦を前提に実装されているため。グレゴリオ暦以外の暦に対応するなら [`time`][time]`.Time` 型は使えないだろう。グレゴリオ暦の導入時期は国によってかなり違うが，欧州では17世紀には概ねグレゴリオ暦に移行したと言われている。日本は1873年（明治6年）のいわゆる「明治の改暦」でグレゴリオ暦互換の暦に移行した（[厳密にはグレゴリオ暦ではない]({{< ref "/remark/2015/japanese-koyomi.md" >}} "「暦」日本史 （再掲載）")）。

## 参考図書

{{% review-paapi "441662140X" %}} <!-- 天文年鑑 2022年版 -->
{{% review-paapi "4805202254" %}} <!-- 天体の位置計算 -->
{{% review-paapi "4769907001" %}} <!-- 新こよみ便利帳 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->

[goark/koyomi]: https://github.com/goark/koyomi "goark/koyomi: 日本のこよみ"
[math/big]: https://pkg.go.dev/math/big "big package - math/big - pkg.go.dev"
[time]: https://pkg.go.dev/time "time package - time - pkg.go.dev"
