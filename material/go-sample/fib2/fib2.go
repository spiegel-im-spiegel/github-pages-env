// +build run

package main

import "fmt"

func fibonacciNumber(n int, fibTable map[int]int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if fn, ok := fibTable[n]; ok {
			return fn
		}
		fn := fibonacciNumber(n-2, fibTable) + fibonacciNumber(n-1, fibTable)
		fibTable[n] = fn
		return fn
	}
}

func main() {
	fmt.Println(fibonacciNumber(40, make(map[int]int)))
}
