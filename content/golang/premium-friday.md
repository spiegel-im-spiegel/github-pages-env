+++
title = "「プレミアムフライデー」を求めるパッケージを作ってみた"
tags = ["golang", "programming", "time", "calendar"]
date = "2017-03-04T09:40:51+09:00"
description = "もちろん息抜きである。潤いは大事。でも実用性はないと思われ。"

[scripts]
  mathjax = false
  mermaidjs = false
+++

そういえば先月の「プレミアムフライデー」，皆様はいかがお過ごしでしたか。
私は3時間も残業してしまいましたよ（笑）

ちうわけで，以下を真似して「プレミアムフライデー」を求めるパッケージを考えてみる。
もちろん息抜きである。
潤いは大事。
でも実用性はないと思われ。

- [プレミアムフライデーを求めるメソッドを作った - Qiita](http://qiita.com/neko_the_shadow/items/4ebf94a8a6d9282e7207)
- [プレミアムフライデーを求めるメソッドを作った（Java8版） - Qiita](http://qiita.com/deaf_tadashi/items/963a62072338f09f12a5)

まずはパッケージ分割しないでベタに書いてみる[^rf1]。

[^rf1]: 元記事のコードがループさせてたんでこっちもついループさせちゃったけど，考えてみれば（いや考えるまでもなく）ループを回す必要はなかった。

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"
)

//GetPremiumFriday returns day of premium friday
func GetPremiumFriday(y int, m time.Month) (int, error) {
    //引数のチェック
    if y < 2017 || m < time.January || m > time.December {
        return 0, os.ErrInvalid
    }
    if y == 2017 && m < time.February { //2017年1月は実施前なのでエラー
        return 0, os.ErrInvalid
    }

    //指定月末（翌月0日）で初期化する
    tm := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で

    w := tm.Weekday() - time.Friday
    if w < 0 {
        w += 7
    }
    return tm.Day() - (int)(w), nil
}

func main() {
    flag.Parse()
    argsStr := flag.Args()
    if len(argsStr) < 2 {
        fmt.Fprintln(os.Stderr, "年月を指定してください")
        return
    }
    args := make([]int, 2)
    for i := 0; i < 2; i++ {
        num, err := strconv.Atoi(argsStr[i])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
        args[i] = num
    }
    d, err := GetPremiumFriday(args[0], time.Month(args[1]))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(d)
}
```

指定月末を求めるのに「翌月0日」で初期化するのがポイント（つか，ここしかポイントになるものがない`w`）。
実行結果は以下の通り。

```text
$ go run pf.go 2017 2
24
```

期待通りの値が得られた。

関数1個だけなんでパッケージにするのもどうかと思うけど折角なのでパッケージ化してみた。

- [spiegel-im-spiegel/pf: Premium Friday](https://github.com/spiegel-im-spiegel/pf)

内容は `GetPremiumFriday()` 関数を切り出しただけ。

```go
package pf

import (
    "os"
    "time"
)

//GetPremiumFriday returns day of premium friday
func GetPremiumFriday(y int, m time.Month) (int, error) {
    //引数のチェック
    if y < 2017 || m < time.January || m > time.December {
        return 0, os.ErrInvalid
    }
    if y == 2017 && m < time.February { //2017年1月は実施前なのでエラー
        return 0, os.ErrInvalid
    }

    //指定月末（翌月0日）で初期化する
    tm := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で

    w := tm.Weekday() - time.Friday
    if w < 0 {
        w += 7
    }
    return tm.Day() - (int)(w), nil
}
```

したがって `main()` 関数はこうなる。

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/spiegel-im-spiegel/pf"
)

func main() {
    flag.Parse()
    argsStr := flag.Args()
    if len(argsStr) < 2 {
        fmt.Fprintln(os.Stderr, "年月を指定してください")
        return
    }
    args := make([]int, 2)
    for i := 0; i < 2; i++ {
        num, err := strconv.Atoi(argsStr[i])
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
        args[i] = num
    }
    d, err := pf.GetPremiumFriday(args[0], time.Month(args[1]))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(d)
}
```

まぁこんなもんかな。
遊んだ遊んだ。

## ブックマーク

- [Golangでプレミアムフライデーかどうか判定する - Qiita](http://qiita.com/qube81/items/1e93c837c0a7e3d99a10)
- [PowerShell Gallery | PremiumFriday](https://www.powershellgallery.com/packages/PremiumFriday/)

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
