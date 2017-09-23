package main

import (
	"fmt"
	"math/big"

	"github.com/robpike/filter"
)

func gcd(m, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(m)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func main() {
	values := []uint64{290021904, 927964716, 826824516, 817140688}

	fmt.Println(filter.Reduce(values, gcd, 1).(uint64)) // 92
}
