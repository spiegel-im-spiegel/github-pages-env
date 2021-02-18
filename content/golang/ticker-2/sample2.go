// +build run

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func ticker(ctx context.Context) {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer func() {
		fmt.Println("Stopping ticker...")
		t.Stop()
	}()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("cancellation from context:", ctx.Err())
			return
		}
	}
}

func run() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	ticker(ctx)
}

func main() {
	run()
}
