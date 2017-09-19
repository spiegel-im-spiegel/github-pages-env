package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func makeGzip(dst io.Writer, content []byte) error {
	zw, err := gzip.NewWriterLevel(dst, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer zw.Close()

	if _, err := zw.Write(content); err != nil {
		return err
	}
	return nil
}

func main() {
	content := []byte("Hello world\n")

	file, err := os.Create("test.txt.gz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	if err := makeGzip(file, content); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
