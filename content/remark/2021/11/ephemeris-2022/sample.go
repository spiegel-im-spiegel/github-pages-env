//go:build run
// +build run

package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/koyomi"
)

func main() {
	start, _ := koyomi.DateFrom("2022-01-01")
	end, _ := koyomi.DateFrom("2022-12-31")
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Eclipse),
		koyomi.WithStartDate(start),
		koyomi.WithEndDate(end),
	).Get()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("| 日付 | 内容 |")
	fmt.Println("| ---- | ---- |")
	for _, e := range k.Events() {
		fmt.Printf("| %v | %v |\n", e.Date, e.Title)
	}
}
