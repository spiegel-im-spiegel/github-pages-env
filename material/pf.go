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
	if y < 2017 || (m < time.January && m > time.December) {
		return 0, os.ErrInvalid
	}
	if y == 2017 && m < time.February { //2017年1月は実施前なのでエラー
		return 0, os.ErrInvalid
	}

	//指定月末（翌月0日）で初期化する
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return 0, err
	}
	tm := time.Date(y, m+1, 0, 0, 0, 0, 0, tz)

	//月末尾から1日ずつ減じて最終金曜日を探す
	for {
		if tm.Weekday() == time.Friday {
			break
		}
		tm = tm.AddDate(0, 0, -1)
	}
	return tm.Day(), nil
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
