package main

import "fmt"

var fibonacciNumbers = make(map[int64]int64)

func fibonacciNumber0(n int64) int64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if fn, ok := fibonacciNumbers[n]; ok {
			return fn
		}
		fn := fibonacciNumber0(n-2) + fibonacciNumber0(n-1)
		fibonacciNumbers[n] = fn
		return fn
	}
}

func main() {
	fmt.Println("| $n$ | $F_n$ |")
	fmt.Println("| ---:| ---:|")
	for n := int64(1); n <= 75; n++ {
		fmt.Printf("| %d | %d |\n", n, fibonacciNumber0(n))
	}
}
