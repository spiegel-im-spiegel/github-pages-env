// +build run

package main

import (
	"fmt"
	"sync"
)

func generater(wg *sync.WaitGroup, ch chan<- int) {
	defer func() {
		close(ch)
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
}

func output(wg *sync.WaitGroup, num int, ch <-chan int) {
	for n := range ch {
		fmt.Printf("Worker %d: %v\n", num, n)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(1)
	go generater(&wg, ch)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go output(&wg, i+1, ch)
	}
	wg.Wait()
}
