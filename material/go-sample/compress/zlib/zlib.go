package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zw := zlib.NewWriter(buf)
	defer zw.Close()

	if _, err := io.Copy(zw, r); err != nil {
		return buf, err
	}
	return buf, nil
}

func extract(zr io.Reader) (io.Reader, error) {
	return zlib.NewReader(zr)
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
