package main

import (
	"fmt"
)

func gen(max int) <-chan int {
	ch := make(chan int)
	go func(max int) {
		for i := 1; i <= max; i++ {
			ch <- i
		}
		close(ch)
	}(max)

	return ch
}

func pipe(ch <-chan int, max int) <-chan int {
	p := make(chan int)
	go func(ch <-chan int) {
		for i := 0; i < max; i++ {
			p <- <-ch
		}
		close(p)
	}(ch)

	return p
}

func main() {
	ch := gen(10)
	p := pipe(ch, 10)
	for i := range p {
		fmt.Println(i)
	}
}
