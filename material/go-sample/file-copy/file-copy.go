package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Println(os.ErrInvalid)
		return
	}
	if err := copyFile(flag.Arg(0), flag.Arg(1)); err != nil {
		fmt.Println(err)
		return
	}
	return
}
