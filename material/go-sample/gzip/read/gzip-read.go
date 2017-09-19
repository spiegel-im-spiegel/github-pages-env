package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt.gz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	zr, err := gzip.NewReader(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer zr.Close()

	io.Copy(os.Stdout, zr)
}
