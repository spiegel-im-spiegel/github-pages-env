package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"walking"

	"github.com/spiegel-im-spiegel/errs"
)

func main() {
	rootPath := "/usr/local/go/src/cmd/go/testdata/script"
	list := walking.NewList()

	start := time.Now()
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errs.Wrap(
				err,
				"prevent panic by handling failure accessing a path",
				errs.WithContext("root_path", rootPath),
				errs.WithContext("path", path),
			)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
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
