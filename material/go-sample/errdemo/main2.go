package main

import (
	"fmt"
	"os"

	"./errdemo1"
)

func main() {
	if err := errdemo1.F(); err != nil {
		switch err.(type) {
		case *os.PathError:
			fmt.Fprintln(os.Stderr, "*os.PathError")
		default:
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}

	return
}
