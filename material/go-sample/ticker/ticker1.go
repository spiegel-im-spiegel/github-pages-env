package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ticker() {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case s := <-sig:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Stop!")
				return
			}
		}
	}
}

func main() {
	ticker()
}
