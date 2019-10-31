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
