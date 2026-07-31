// +build run

package main

import (
	"fmt"
	"time"
)

func ticker() {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer func() {
		fmt.Println("Stopping ticker...")
		t.Stop()
	}()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		}
	}
}

func main() {
	ticker()
}
