// +build run

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y, z big.Float //zero initialize
	y.SetFloat64(0.1)     //53bit precision
	x.SetPrec(64)
	y.SetPrec(64)

	for i := 0; i < 10; i++ {
		z.Add(&x, &y)
		x.Set(&z)
		fmt.Printf("x = %v (prec = %d bits)\n", &x, x.Prec())
	}
}
