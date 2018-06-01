package main

import (
	"fmt"
	"strconv"
)

func newCounter() func() int {
	ct := 0
	return func() int {
		ct++
		return ct
	}
}

func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "Fizz Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func main() {
	max := 30
	c := newCounter()
	for n := c(); ; n = c() {
		if n < max {
			fmt.Print(fizzBuzz(n), ", ")
		} else {
			fmt.Println(fizzBuzz(n))
			break
		}
	}
}
