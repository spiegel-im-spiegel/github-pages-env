package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/koyomi"
)

var weekShortNames = []string{"日", "月", "火", "水", "木", "金", "土"}

func WeekShortNameJp(dt koyomi.DateJp) string {
	return weekShortNames[dt.Weekday()%7]
}

func main() {
	start, _ := koyomi.DateFrom("2021-01-01")
	end, _ := koyomi.DateFrom("2021-12-31")
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("| 日付 | 曜日 | 内容 | 備考 |")
	fmt.Println("| ---- |:----:| ---- | ---- |")
	for _, e := range k.Events() {
		fmt.Printf("| %v | %v | %v |  |\n", e.Date, WeekShortNameJp(e.Date), e.Title)
	}
}
