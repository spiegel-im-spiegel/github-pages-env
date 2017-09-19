package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func makeTarGzip(w io.Writer) error {
	zw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer zw.Close()

	tw := tar.NewWriter(zw)
	defer tw.Close()

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)

		hd, e := tar.FileInfoHeader(info, "")
		if e != nil {
			return e
		}
		content, e := ioutil.ReadFile(path)
		if e != nil {
			return e
		}

		if e := tw.WriteHeader(hd); e != nil {
			return e
		}
		if _, e := tw.Write(content); e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	file, err := os.Create("test.tar.gz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	if err := makeTarGzip(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
