package main

import (
	"fmt"
	"os"

	"errdemo/errdemo2"
)

func main() {
	if err := errdemo2.F(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	return
}
