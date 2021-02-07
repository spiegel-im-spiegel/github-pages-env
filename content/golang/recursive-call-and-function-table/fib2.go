// +build run

package main

import "fmt"

func fibonacciNumber2(n int, fibMap map[int]int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if fn, ok := fibMap[n]; ok {
			return fn
		}
		fn := fibonacciNumber2(n-2, fibMap) + fibonacciNumber2(n-1, fibMap)
		fibMap[n] = fn
		return fn
	}
}

func main() {
	fmt.Println(fibonacciNumber2(40, make(map[int]int)))
}
