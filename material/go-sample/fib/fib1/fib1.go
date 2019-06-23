package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("| $n$ | $\\varphi^n$ | $\\psi^n$ | $F_n$ |")
	fmt.Println("| ---:| ---:| ---:| ---:|")
	r5 := math.Sqrt(5)
	phi := (1 + math.Sqrt(5)) / 2
	psi := 1 - phi
	for n := int64(1); n <= 75; n++ {
		phin := math.Pow(phi, float64(n))
		psin := math.Pow(psi, float64(n))
		f := (phin - psin) / r5
		fmt.Printf("| %d | %.15f | %.15f | %.15f |\n", n, phin, psin, f)
	}
}
