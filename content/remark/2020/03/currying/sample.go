package main

import "fmt"

func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	fmt.Println(add(1)(2)) //Output: 3
	increment := add(1)
	fmt.Println(increment(2)) //Output: 3
}
