package main

import (
	"fmt"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/zodiac"
)

func main() {
	start, _ := koyomi.DateFrom("2022-04-17")
	end, _ := koyomi.DateFrom("2022-05-05")
	for d := start; d.Before(end); d = d.AddDay(1) {
		干, 支 := zodiac.ZodiacDayNumber(d)
		fmt.Printf("Day %v is %v%v\n", d, 干, 支)
	}
}
