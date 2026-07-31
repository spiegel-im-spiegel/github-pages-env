+++
title = "時刻の「時」以下を切り捨てる処理"
date =  "2019-10-31T23:06:37+09:00"
description = "無責任なこと言って，マジすんません。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "time" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昔 Qiita で

- [golangのtime.Timeの当日00:00:00を取得する方法とベンチマーク - Qiita](https://qiita.com/ushio_s/items/3e270933641710bbd88e)

という記事があって，無責任にも「Unix-Time を使えばいいんじゃないの？」みたいなコメントを残した後にすっかり忘れていたのだが，ありがたくも[バグの報告](https://qiita.com/go_sagawa/items/836398020100df486184)を頂いた。

最初に提案したコードは

```go
now := time.Now()
ut := now.Unix()
_, offset := now.Zone()
day := time.Unix((ut/86400)*86400-int64(offset), 0)
```

だったのだが，見ての通り，時差の考慮が不完全である。
したがって最後の行を

```go
day := time.Unix(((ut+int64(offset))/86400)*86400-int64(offset), 0).In(now.Location())
```

とする必要がある。
いや，マジすんません `m(_ _)m`

ついでなので，この記事でもベンチマークテストをしておこう。
対象コードは以下の5つ。

```go
package zerotime

import (
	"fmt"
	"time"
)

//fmt パッケージを使って文字列整形する
func TrancateByFmt(t time.Time) time.Time {
	day, _ := time.Parse(
		"2006-01-02 15:04:05 -0700",
		fmt.Sprintf("%s 00:00:00 %s", t.Format("2006-01-02"), t.Format("-0700")),
	)
	return day
}

//time.Formatを使って文字列整形する
func TrancateByFormat(t time.Time) time.Time {
	day, _ := time.Parse(
		"2006-01-02 15:04:05 -0700 MST",
		t.Format("2006-01-02 00:00:00 -0700 MST"),
	)
	return day
}

//差分を計算してtime.Time.Add関数を使う
func TrancateByCalc(t time.Time) time.Time {
	nanosecond := time.Duration(t.Nanosecond())
	second := time.Duration(t.Second())
	minute := time.Duration(t.Minute())
	hour := time.Duration(t.Hour())
	dur := -1 * (nanosecond + second*time.Second + minute*time.Minute + hour*time.Hour)
	return t.Add(dur)
}

//time.Date関数で0時0分を再設定する
func TrancateByDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

//Unix-Timeを使う
func TrancateByUnixTime(t time.Time) time.Time {
	_, offset := t.Zone()
	return time.Unix(((t.Unix()+int64(offset))/86400)*86400-int64(offset), 0).In(t.Location())
}
```

これらに対応するベンチマークテストは以下の通り。

```go
func BenchmarkTrancateByFmt(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByFmt(t)
	}
}

func BenchmarkTrancateByFormat(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByFormat(t)
	}
}

func BenchmarkTrancateByCalc(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByCalc(t)
	}
}

func BenchmarkTrancateBydate(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByDate(t)
	}
}

func BenchmarkTrancateByUnixTime(b *testing.B) {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zerotime.TrancateByUnixTime(t)
	}
}
```

ほんで，実行結果は以下の通り。

```text
$ go test -bench Trancate -benchmem
goos: linux
goarch: amd64
pkg: zerotime
BenchmarkTrancateByFmt-4        	  987878	      1096 ns/op	     224 B/op	       8 allocs/op
BenchmarkTrancateByFormat-4     	 1233117	       953 ns/op	     176 B/op	       4 allocs/op
BenchmarkTrancateByCalc-4       	33584716	        31.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrancateBydate-4       	11734460	        99.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrancateByUnixTime-4   	77463438	        15.5 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	zerotime	7.601s
```

分かりやすく表にしておこう。

| 関数名               | 実行時間 | Alloc サイズ | Alloc 回数 |
| -------------------- | --------:| ------------:| ----------:|
| `TrancateByFmt`      | 1,096 ns |    224 bytes |          8 |
| `TrancateByFormat`   |   953 ns |    176 bytes |          4 |
| `TrancateByCalc`     |  31.0 ns |      0 bytes |          0 |
| `TrancateByDate`     |  99.7 ns |      0 bytes |          0 |
| `TrancateByUnixTime` |  15.5 ns |      0 bytes |          0 |

まぁ，文字列を介した処理が遅いのは当然として， [`time`]`.Date()` を使った処理が意外に遅いな。
いや，内部で暦計算をしてるならこんなもんか？

改めて

{{< fig-quote title="sudo の特権昇格バグはなぜ起こったのか" link="https://mattn.kaoriya.net/software/lang/c/20191016143950.htm" >}}
<q>境界値チェックを行わないと、死ぬ</q>
{{< /fig-quote >}}

お後がよろしいようで。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`time`]: https://golang.org/pkg/time/ "time - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
