package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y, z big.Rat //zero initialize
	var a, b big.Int
	a.SetInt64(1)
	b.SetInt64(10)
	y.SetFrac(&a, &b)

	for i := 0; i < 10; i++ {
		z.Add(&x, &y)
		x.Set(&z)
		fmt.Printf("x = %s (%v)\n", x.FloatString(20), &x)
	}
}
