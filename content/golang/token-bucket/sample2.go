// +build run

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func generater(wg *sync.WaitGroup, ch chan<- int) {
	defer func() {
		close(ch)
		wg.Done()
	}()
	l := rate.NewLimiter(rate.Every(time.Second*2), 1)
	for i := 0; i < 10; i++ {
		if err := l.Wait(context.Background()); err != nil {
			fmt.Printf("generater: %v\n", err)
			return
		}
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
