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

	paths := []string{}
	for _, path := range os.Args[1:] {
		res := file.Glob(path, nil)
		if len(res) > 0 {
			paths = append(paths, res...)
		}
	}

	for _, path := range paths {
		fmt.Println(path)
	}
}
