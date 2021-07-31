+++
title = "十干十二支を数え上げるパッケージを作ってみた"
date =  "2021-07-31T16:46:36+09:00"
description = "検算には，若い頃に買った『新こよみ便利帳』に載っていた表を使った。 何でも取っておくものである（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

夏の終りの土用の丑の日に

- [土用の丑の日なので...](https://zenn.dev/spiegel/articles/20210728-zodiac-day)

という記事を Zenn に上げたが，調子に乗って [Go] パッケージ化してみた。

- [spiegel-im-spiegel/jzodiac: Japanese Zodiac](https://github.com/spiegel-im-spiegel/jzodiac)

これを使って，たとえば

```go
// +build run

package main

import (
    "flag"
    "fmt"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/jzodiac"
)

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) < 1 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
        return
    }
    for _, s := range args {
        t, err := time.Parse("2006-01-02", s)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }
        kan, shi := jzodiac.ZodiacYearNumber(t.Year())
        fmt.Printf("Year %v is %v%v\n", t.Year(), kan, shi)
        kan, shi = jzodiac.ZodiacDayNumber(t)
        fmt.Printf("Day %v is %v%v\n", t.Format("2006-01-02"), kan, shi)
    }
}
```

というコードを組めば

```text
$ go run sample.go 2021-07-28
Year 2021 is 辛丑
Day 2021-07-28 is 丁丑
```

という感じに年と日の干支を取得できる。
ちなみに月の干支は旧暦と連動しているので，今回は実装していない。

検算には，若い頃に買った『[新こよみ便利帳](https://www.amazon.co.jp/dp/4769907001?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』に載っていた表を使った。
何でも取っておくものである（笑）

[Go]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4769907001" %}} <!-- 新こよみ便利帳 -->
