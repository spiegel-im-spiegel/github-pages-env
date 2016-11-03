package main

import (
	"fmt"
	"os"

	"errdemo/errdemo2"

	"github.com/pkg/errors"
)

func main() {
	if err := errdemo2.F(); err != nil {
		switch errors.Cause(err).(type) {
		case *os.PathError:
			fmt.Fprintln(os.Stderr, "*os.PathError")
		default:
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}

	return
}
