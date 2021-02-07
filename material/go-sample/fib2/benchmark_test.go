package fib2

import "testing"

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

func BenchmarkFibonacciNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibonacciNumber(40)
	}
}

func BenchmarkFibonacciNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibonacciNumber2(40, make(map[int]int))
	}
}
