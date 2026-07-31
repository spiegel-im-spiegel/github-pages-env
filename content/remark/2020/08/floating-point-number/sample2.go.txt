package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewRat(1, 10)
	b := big.NewRat(2, 10)
	c := (&big.Rat{}).Add(a, b)
	fmt.Println(big.NewRat(3, 10).Cmp(c) == 0)
}
