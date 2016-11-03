package main

import (
	"fmt"
	"os"

	"./errdemo1"
)

func main() {
	if err := errdemo1.F(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	return
}
