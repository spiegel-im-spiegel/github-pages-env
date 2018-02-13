package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	f, err := Assets.Open("/data/hello.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	io.Copy(os.Stdout, f)
}
