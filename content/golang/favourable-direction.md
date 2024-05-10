+++
title = "今年の恵方を取得する"
date =  "2023-01-14T16:21:26+09:00"
description = "koyomi パッケージに追加した"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

正月気分も終わり，近所のスーパーとかで節分グッズが並ぶようになって「そういや今年の恵方はどっちだっけ？」と思って調べてみた。

どうやら，その年の十干で恵方が決定するらしい。

{{< fig-img-quote src="./Ehou-direction.png" title="File:Ehou-direction.png - Wikimedia Commons" link="https://commons.wikimedia.org/wiki/File:Ehou-direction.png" lang="en" >}}

って，4方向を十干で分けるのか。
中途半端やなぁ。

まぁ，でも，その年の十干が分かれば恵方も分かるわけだ。
拙作の [`github.com/goark/koyomi`](https://github.com/goark/koyomi "goark/koyomi: 日本のこよみ") パッケージにオマケで十干十二支を数え上げる機能を付けているのだが，これに恵方を取得する機能を追加してみた。

こんな感じ。

```go
package main

import (
    "fmt"

    "github.com/goark/koyomi/zodiac"
)

func main() {
    year := 2023
    干, 支 := zodiac.ZodiacYearNumber(year)
    fmt.Printf("%d年は%v%v，恵方は%v (%v°)", year, 干, 支, 干.DirectionJp(), 干.Direction())
}
```

これを[実行](https://go.dev/play/p/PEKVng6jwFc)すると

```text
2023年は癸卯，恵方は南南東微南 (165°)
```

と出力される。

というわけで，2023年はだいたい南南東を向いてモグモグすればいいらしい。

この「恵方」というのはその年の歳神様のおられる方位で，居住地から見て恵方にあたる社寺に詣ることを「恵方詣り」と言うらしいのだが，明治以降の鉄道の発達で長距離移動が容易になり，方位に依存する「恵方詣り」が廃れ，代わりに「初詣」が主流になっていったとのこと。
その一方で，節分の「恵方巻き」みたいな文化が平成以降に台頭してくるのは面白い。

[Go]: https://go.dev/

## ブックマーク

- {{< pdf-file title="食卓の縁起に関する研究 I　－恵方巻の受容とその背景－" link="http://libro.do-bunkyodai.ac.jp/research/pdf/journal32/12.pdf" >}}
- [恵方の方角！2023年はどっち？恵方の決め方も具体的に解説！ | いい日本再発見](https://ii-nippon.net/%e6%97%a5%e6%9c%ac%e3%81%ae%e9%a2%a8%e7%bf%92/1183.html)
- [歳徳神 - Wikipedia](https://ja.wikipedia.org/wiki/%E6%AD%B3%E5%BE%B3%E7%A5%9E)

- [節分どうでしょう]({{< ref "/remark/2018/02/folklore.md" >}})

## 参考図書

{{% review-paapi "B0191845R0" %}} <!-- 鉄道が変えた社寺参詣 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
