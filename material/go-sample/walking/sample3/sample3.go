package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/saracen/walker"
)

func WalkWithSingle(rootPath string) (int64, time.Duration, error) {
	count := int64(0)
	start := time.Now()
	lastErr := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return count, time.Since(start), lastErr
}

func WalkWithMultiple(rootPath string) (int64, time.Duration, error) {
	count := int64(0)
	start := time.Now()
	err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
		if !info.IsDir() {
			atomic.AddInt64(&count, 1)
		}
		return nil
	})
	return count, time.Since(start), err
}

func main() {
	rootPath := "/usr/local/go/src"
	ct, dt, err := WalkWithSingle(rootPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("WalkWithSingle")
	fmt.Println("\tDuration:", dt)
	fmt.Println("\t   Count:", ct)

	ct, dt, err = WalkWithMultiple(rootPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("WalkWithMultiple")
	fmt.Println("\tDuration:", dt)
	fmt.Println("\t   Count:", ct)
}
