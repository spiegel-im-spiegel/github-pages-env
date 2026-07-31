package main

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func compress(dst io.Writer, src io.Reader) error {
	zw := zlib.NewWriter(dst)
	defer zw.Close()

	if _, err := io.Copy(zw, src); err != nil {
		return err
	}
	return nil
}

const maxSize = 1024 * 1024 * 1024 //1GB

var ErrTooLarge = errors.New("too laege decompressed data")

func extract(dst io.Writer, src io.Reader) error {
	r, err := zlib.NewReader(src)
	if err != nil {
		return err
	}
	if _, err := io.CopyN(dst, r, maxSize); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	return ErrTooLarge
}

func main() {
	content := "Hello world\n" //raw data
	zbuf := &bytes.Buffer{}
	//compress raw data
	if err := compress(zbuf, strings.NewReader(content)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display compressed data
	b := zbuf.Bytes()
	fmt.Printf("%d bytes: %v\n", len(b), b)
	//extract from compressed data
	buf := &bytes.Buffer{}
	err := extract(buf, bytes.NewReader(b))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//display extracted data
	if _, err := io.Copy(os.Stdout, buf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
