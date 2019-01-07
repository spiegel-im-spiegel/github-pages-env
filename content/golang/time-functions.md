+++
date = "2016-01-22T20:56:21+09:00"
update = "2016-05-21T10:27:41+09:00"
description = "今回は日付処理の話。特にフォーマットの定義の仕方はよく忘れるので覚え書きとして記しておく。"
draft = false
tags = ["golang", "time"]
title = "Go 言語の日付処理"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

今回は日付処理の話。
特にフォーマットの定義の仕方はよく忘れるので覚え書きとして記しておく。

[Go 言語]で日付処理を行うには [`time`] パッケージを使う。
よく使う型としては

- [`time`].`Time`
- [`time`].`Duration`
- [`time`].`Location`

が挙げられるだろう。
[`time`].`Time` は時刻を， [`time`].`Duration` は2時点間の時間を，そして [`time`].`Location` は地球上の時差を表す型である。
たとえば

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    tz, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    tm := time.Date(2006, 1, 2, 15, 4, 5, 0, tz)
    fmt.Println(tm)
    now := time.Now()
    fmt.Println(now)
    d := now.Sub(tm)
    fmt.Println(d)
}
```

と書くと

```
2006-01-02 15:04:05 +0900 JST
2009-11-10 23:00:00 +0000 UTC
33808h55m55s
```

てな感じになる[^m]。

[^m]: 厳密にいうと月の値は [`time`].`Month` 型である。サンプル・コードのようにリテラルな数値なら気にする必要はないが，変数を [`time`].`Date()` 関数にセットする場合は注意が必要である。

時刻を任意のフォーマットで表示する場合は少し特殊な方法を使う。
たとえば [RFC 3339](https://tools.ietf.org/html/rfc3339) フォーマットに出力するなら

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    tz, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    tm := time.Date(2015, 12, 31, 23, 59, 59, 0, tz)
    fmt.Println(tm.Format("2006-01-02T15:04:05Z07:00"))
}
```

とすれば

```
2015-12-31T23:59:59+09:00
```

と出力される。
テンプレート文字列が `%Y` とか `YYYY` のような形式ではないのだ。
これは適当な文字列ではなく一応法則があって

- 月は `1` （または `01`, `Jan`, `January`）
- 日は `2` （または `02`, `_2`）
- 時は `3` （または `03`, `15`） ※ 午後3時を指す
- 分は `4` （または `04`）
- 秒は `5` （または `05`）
- 年は `06` （または `2006`）
- 時差は `-07` （または `-0700`, `-07:00`, `Z07:00`, `MST`[^tz] など）
- 曜日は `Mon` （または `Monday`）
- AM/PM は `PM` （または `pm`）

という感じに 1 からの連番になっている（曜日等は例外だけど）ので，まぁ覚えられるかな？ でもよく忘れるんだよなぁ。
`%Y` みたいなのとどちらがいいかは微妙な気がするが，慣れの問題かもしれない。

[^tz]: 時差の MST は米国の山岳部時間（Mountain Standad Time）を指すらしい。ソルトレイクシティとかデンバーとかかな。

典型的なフォーマットは定数化されている。

```go
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
```

なので，先ほどのコードも出力部分を

```go
fmt.Println(tm.Format(time.RFC3339))
```

とすれば同じ結果が得られる。
時刻フォーマットは，いったんシステムの中で決めてしまえば同じものを使い回すことになると思うので，定数化してしまえば「フォーマットどうだっけ？」と煩わされることも少ないだろう。

ところでバージョン 1.5 系の [`time`].`Parse()` 関数は日付の解釈が寛容で，各月の末日を31日まで許容している。
たとえば閏年でない2月29日でも

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    tm, err := time.Parse(time.RFC3339, "2015-02-29T23:59:59+09:00")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Println(tm)
}
```

```
2015-03-01 23:59:59 +0900 +0900
```

となり，エラーとならずいい感じ（？）に加減してくれるのだが，バージョン 1.6 からは少し解釈が厳密になりエラーを返すようだ。

{{< fig-quote title="Go 1.6 Release Notes - The Go Programming Language" link="https://tip.golang.org/doc/go1.6" lang="en" >}}
<q>The <a href="https://tip.golang.org/pkg/time/"><code>time</code></a> package's <a href="https://tip.golang.org/pkg/time/#Parse"><code>Parse</code></a> function has always rejected any day of month larger than 31, such as January 32. In Go 1.6, Parse now also rejects February 29 in non-leap years, February 30, February 31, April 31, June 31, September 31, and November 31.</q>
{{< /fig-quote >}}

実際に 1.6 で上のコードを実行すると

```
parsing time "2015-02-29T23:59:59+09:00": day out of range
```

とエラーが返ってくる。

ちなみに [`time`].`Date()` 関数は更に寛容である。

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    tz, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    tm := time.Date(2015, 13, 32, 25, 60, 00, 0, tz)
    fmt.Println(tm)
}
```

```
2016-02-02 02:00:00 +0900 JST
```

## ブックマーク

- [Goで良い感じに日時をパースするライブラリdatemakiの話とGo 1.6 - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2015/12/22/000011)
- [golangのtime.Timeの当日00:00:00を取得する方法とベンチマーク - Qiita](http://qiita.com/ushio_s/items/3e270933641710bbd88e)
- [Golang 日付のフォーマットでハマった話 - Qiita](http://qiita.com/masa23/items/e781124a7e0305bc40c4)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`time`]: http://golang.org/pkg/time/
