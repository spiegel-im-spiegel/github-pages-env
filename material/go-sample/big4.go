// +build run

package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	x := decimal.NewFromFloat(0)
	y, _ := decimal.NewFromString("0.1")

	for i := 0; i < 10; i++ {
		x = x.Add(y)
		fmt.Printf("x = %s\n", x.StringFixed(20))
	}
}
