package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/gocli/file"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
	}

	for _, path := range os.Args[1:] {
		paths := file.Glob(path, file.NewGlobOption())
		if len(paths) > 0 {
			for _, path := range paths {
				fmt.Println(path)
			}
		}
	}
}
