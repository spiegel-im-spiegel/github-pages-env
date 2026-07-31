// +build run

package main

import "fmt"

func fibonacciNumber(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacciNumber(n-2) + fibonacciNumber(n-1)
	}
}

func main() {
	fmt.Println(fibonacciNumber(40))
}
