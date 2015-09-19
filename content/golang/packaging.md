+++
date = "2015-09-18T17:26:39+09:00"
description = "description"
draft = true
tags = ["golang", "package", "github"]
title = "機能のパッケージ化"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) その6 - Qiita](http://qiita.com/spiegel-im-spiegel/items/404871d2bafd22bdbb90)）

[前回]の続き。
なんだけど、今回はパッケージのお話。
ユリウス日の計算なんて簡単なので今まで `main()` 関数の中にゴリゴリ書いてましたが、今後のことを考えてパッケージ化を行うことにします。

まずは、[前回]のコードから計算処理部分をきちんと分離します。

```go:jd4.go
package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func ModifiedJulianDayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return mjdnUnix(t)
	} else {
		return mjdnGregorian(t)
	}
}

func mjdnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

func mjdnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			if enum, ok := err.(*strconv.NumError); ok {
				switch enum.Err {
				case strconv.ErrRange:
					fmt.Printf("Bad Range Error")
				case strconv.ErrSyntax:
					fmt.Printf("Syntax Error")
				}
			}
			return
		} else {
			args[i] = num
		}
	}
	tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	fmt.Printf("MJD = %d日\n", ModifiedJulianDayNumber(tm))
}
```

```shell
C:>go run jd4.go 1969 12 31
1969-12-31 00:00:00 +0000 UTC
MJD = 40586日

C:>go run jd4.go 1970 1 1
1970-01-01 00:00:00 +0000 UTC
MJD = 40587日

C:>go run jd4.go 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

ユリウス日の端数が鬱陶しいので修正ユリウス日の整数部分のみ計算しています。
あと1970年1月1日を境界として計算方法を変えています。
本当はユリウス暦の場合の計算も含めるべきなんでしょうけど、今回は割愛します。












## ブックマーク

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< ref "hugo/julian-day-number.md" >}} "「ユリウス日」で遊ぶ"
