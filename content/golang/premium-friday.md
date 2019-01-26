+++
draft = false
tags = ["golang", "programming", "time", "calendar"]
date = "2017-03-04T09:40:51+09:00"
update =  "2018-12-30T16:08:46+09:00"
title = "「プレミアムフライデー」を求めるパッケージを作ってみた"
description = "もちろん息抜きである。潤いは大事。でも実用性はないと思われ。"

[author]
  flickr = "spiegel"
  license = "by-sa"
  name = "Spiegel"
  flattr = ""
  linkedin = "spiegelimspiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  tumblr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
