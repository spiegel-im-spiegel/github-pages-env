package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spiegel-im-spiegel/gocli/signal"
)

func ticker(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second) // 1 second cycle
	defer t.Stop()

	for {
		select {
		case now := <-t.C: // ticker event
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done(): // cancel event from context
			fmt.Println("Stop ticker")
			return ctx.Err()
		}
	}
}

func Run() error {
	errCh := make(chan error, 1)
	defer close(errCh)

	go func() {
		child, cancelChild := context.WithTimeout(
			signal.Context(context.Background(), os.Interrupt), // cancel event by SIGNAL
			10*time.Second, // timeout after 10 seconds
		)
		defer cancelChild()
		errCh <- ticker(child)
	}()

	err := <-errCh
	fmt.Println("Done")
	return err
}

func main() {
	if err := Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
