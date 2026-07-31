package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

func makeGzip(body []byte) ([]byte, error) {
	var b bytes.Buffer
	err := func() error {
		gw := gzip.NewWriter(&b)
		defer gw.Close()

		if _, err := gw.Write(body); err != nil {
			return err
		}
		return nil
	}()
	return b.Bytes(), err
}

func main() {
	content := []byte("Hello world\n")

	file, err := os.Create("test.txt.gz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	z, err := makeGzip(content)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if _, err := file.Write(z); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
