package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"walking"

	"github.com/saracen/walker"
)

func main() {
	rootPath := "/usr/local/go/src/cmd/go/testdata/script"
	list := walking.NewList()
	fmt.Println("CPUs:", runtime.NumCPU())

	start := time.Now()
	err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
		if !info.IsDir() {
			list.Append(path)
		}
		return nil
	})
	fmt.Println("Time:", time.Now().Sub(start))
	if err != nil {
		fmt.Printf("%+v\n", err) //for debug
		return
	}
	fmt.Println("Count:", list.Count())
	for _, path := range list.List() {
		fmt.Println(path)
	}
}
