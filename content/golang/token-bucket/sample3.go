// +build run

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func generater(wg *sync.WaitGroup, ch chan<- int, max int) {
	defer func() {
		close(ch)
		wg.Done()
	}()
	l := rate.NewLimiter(rate.Every(time.Second*2), max)
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
	max := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	ch := make(chan int, max)
	wg.Add(1)
	go generater(&wg, ch, max)
	for i := 0; i < max; i++ {
		wg.Add(1)
		go output(&wg, i+1, ch)
	}
	wg.Wait()
}
