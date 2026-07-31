// +build run

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ticker(ctx context.Context, num int) {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer func() {
		fmt.Printf("Stopping ticker (%d) ...\n", num)
		t.Stop()
	}()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Printf("cancellation from context (%d) : %v\n", num, ctx.Err())
			return
		}
	}
	return
}

func run() {
	parent, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// parent, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer func() {
	// 	fmt.Println("cancellation by WithTimeout 0")
	// 	cancel()
	// }()

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		n := i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			child, _ := context.WithTimeout(parent, time.Duration(n)*5*time.Second)
			// defer func() {
			// 	fmt.Println("cancellation by WithTimeout", n)
			// 	cancelChild()
			// }()
			ticker(child, n)
		}()
	}
	wg.Wait()
}

func main() {
	run()
}
