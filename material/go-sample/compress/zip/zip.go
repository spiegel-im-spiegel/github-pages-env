package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zw, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		return buf, err
	}
	if _, err := io.Copy(zw, r); err != nil {
		return buf, err
	}
	return buf, zw.Close()
}

func extract(zr io.Reader) (io.Reader, error) {
	return flate.NewReader(zr), nil
}

func main() {
	content := "Hello world\n" //raw data
	//compress raw data
	zr, err := compress(bytes.NewBufferString(content))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display compressed data
	b := zr.Bytes()
	fmt.Printf("%d bytes: %v\n", len(b), b)
	//extract from compressed data
	r, err := extract(zr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display extracted data
	io.Copy(os.Stdout, r)
}
