package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io"
	"os"
)

func extract(zr io.Reader) (io.Reader, error) {
	return bzip2.NewReader(zr), nil
}

func main() {
	cdata := []byte{66, 90, 104, 57, 49, 65, 89, 38, 83, 89, 176, 48, 136, 246, 0, 0, 1, 85, 128, 0, 16, 64, 0, 0, 64, 6, 4, 144, 128, 32, 0, 34, 40, 246, 166, 244, 8, 6, 4, 105, 205, 172, 132, 162, 238, 72, 167, 10, 18, 22, 6, 17, 30, 192}
	//display compressed data
	fmt.Printf("%d bytes: %v\n", len(cdata), cdata)
	//extract from compressed data
	r, err := extract(bytes.NewReader(cdata))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display extracted data
	io.Copy(os.Stdout, r)
}
