// +build run

package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	x := decimal.NewFromFloat(0)
	y, _ := decimal.NewFromString("1.0")

	for i := 0; i < 100000000; i++ {
		x = x.Add(y)
	}
	fmt.Printf("z = %s\n", x.StringFixed(20))
}
