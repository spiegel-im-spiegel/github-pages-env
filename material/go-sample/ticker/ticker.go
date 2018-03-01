package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ticker(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("Stop child")
			return ctx.Err()
		}
	}
}

func SignalContext(ctx context.Context) context.Context {
	parent, cancelParent := context.WithCancel(ctx)
	go func() {
		defer cancelParent()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		defer close(sig)

		select {
		case <-parent.Done():
			fmt.Println("Cancel from parent")
			return
		case s := <-sig:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Stop!")
				return
			}
		}
	}()

	return parent
}

func Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	defer close(errCh)

	parent := SignalContext(ctx)
	go func() {
		child, cancelChild := context.WithTimeout(parent, 10*time.Second)
		defer cancelChild()
		errCh <- ticker(child)
	}()

	err := <-errCh
	fmt.Println("Done parent")
	return err
}

func main() {
	ctx := context.Background()
	if err := Run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Done")
}
