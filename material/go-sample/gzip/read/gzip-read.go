package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func readGzip(dst io.Writer, src io.Reader) error {
	zr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer zr.Close()

	io.Copy(dst, zr)

	return nil
}

func main() {
	file, err := os.Open("test.txt.gz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	if err := readGzip(os.Stdout, file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
