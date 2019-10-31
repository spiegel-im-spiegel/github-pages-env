package main

import (
	"fmt"
	"time"

	"zerotime"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Date(2019, time.October, 31, 23, 34, 56, 0, loc)
	fmt.Println(zerotime.TrancateByFmt(t))
	fmt.Println(zerotime.TrancateByFormat(t))
	fmt.Println(zerotime.TrancateByCalc(t))
	fmt.Println(zerotime.TrancateByDate(t))
	fmt.Println(zerotime.TrancateByUnixTime(t))
}
