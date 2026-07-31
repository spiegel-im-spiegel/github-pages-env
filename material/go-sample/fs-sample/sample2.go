package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func output(fileSystem fs.FS) error {
	paths, err := fs.Glob(fileSystem, "*.txt")
	if err != nil {
		return err
	}
	for _, path := range paths {
		if _, err := func(path string) (int64, error) {
			f, err := fileSystem.Open(path)
			if err != nil {
				return 0, err
			}
			defer f.Close()
			return io.Copy(os.Stdout, f)
		}(path); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	rc, err := zip.OpenReader("./testdata.zip")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer rc.Close()

	fileSystem, err := fs.Sub(rc, "testdata")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := output(fileSystem); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
