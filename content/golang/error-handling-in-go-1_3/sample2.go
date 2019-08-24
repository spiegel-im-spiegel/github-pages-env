package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func checkFileOpen(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func main() {
	if err := checkFileOpen("not-exist.txt"); err != nil {
		var errno syscall.Errno
		if errors.As(err, &errno) {
			fmt.Fprintln(os.Stderr, errno)
			return
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
