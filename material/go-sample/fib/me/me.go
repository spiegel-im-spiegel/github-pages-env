package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("| $n$ | $\\epsilon$ | $1+\\epsilon-1$ |")
	fmt.Println("| ---:| ---:| ---:|")
	for n := int64(1); n <= 55; n++ {
		e := math.Pow(2.0, float64(-n))
		fmt.Println(e)
		//me := 1.0 + e - 1.0
		//fmt.Printf("| %d | %v | %v |\n", n, e, me)
	}
}
