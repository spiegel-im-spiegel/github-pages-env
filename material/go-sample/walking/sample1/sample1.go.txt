package main

import (
	"fmt"
	"os"
	"sort"

	"walking"

	"github.com/saracen/walker"
)

func WalkWithMultiple(rootPath string) ([]string, error) {
	list := walking.New()
	err := walker.Walk(rootPath, func(path string, info os.FileInfo) error {
		if !info.IsDir() {
			list.Append(path)
		}
		return nil
	})
	return list.List(), err
}

func main() {
	rootPath := "/usr/local/go/src"
	list, err := WalkWithMultiple(rootPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("WalkWithMultiple ( Count =", len(list), "):")
	sort.Strings(list)
	for i, path := range list {
		fmt.Println("\t", path)
		if i > 10 {
			break
		}
	}
}
