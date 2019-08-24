package main

import (
	"fmt"
	"os"
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
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
