package main

import (
	"flag"
	"fmt"
	"os"
)

//Run returns error status in proess.
func Run(args []string) error {
	f := flag.NewFlagSet("fillzero", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: fillzero -n <number of digits> [digit string]")
	}
	num := f.Int("n", 0, "number of digits")

	if err := f.Parse(args); err != nil {
		return err
	}
	if *num == 0 {
		return os.ErrInvalid
	}
	if f.NArg() > 1 {
		return os.ErrInvalid
	}

	var str string
	if f.NArg() == 1 {
		str = f.Arg(0)
	}

	// Fill in digits...

	fmt.Println(str)
	return nil
}

func main() {
	if err := Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
