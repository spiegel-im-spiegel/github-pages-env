// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		n := i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			ticker(ctx, n)
		}()
	}
	wg.Wait()
}

func main() {
	run()
}
