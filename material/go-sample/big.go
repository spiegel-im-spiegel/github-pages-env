package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x, y, z big.Float //zero initialize
	y.SetInt64(1)         //64bit precision

	for i := 0; i < 100000000; i++ {
		z.Add(&x, &y)
		x.Copy(&z)
	}
	fmt.Printf("x = %.10g (%s, prec = %d, acc = %s)\n", &x, x.Text('p', 0), x.Prec(), x.Acc())
}
