package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/dsnet/compress/bzip2"
)

func compress(r io.Reader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	conf := new(bzip2.WriterConfig)
	conf.Level = bzip2.BestCompression
	zw, err := bzip2.NewWriter(buf, conf)
	if err != nil {
		return buf, err
	}
	if _, err := io.Copy(zw, r); err != nil {
		return buf, err
	}
	return buf, zw.Close()
}

func extract(zr io.Reader) (io.Reader, error) {
	return bzip2.NewReader(zr, new(bzip2.ReaderConfig))
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
