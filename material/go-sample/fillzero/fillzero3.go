package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Identity is digit string class
type Identity struct {
	id uint64 // identity
	nd int    // number of digits
}

//IdentityNew returns Identity instance.
func IdentityNew(n int, s string) (*Identity, error) {
	id := &Identity{}
	if n <= 0 {
		return id, errors.New("Number of digits is greater than zero.")
	}
	if len(s) > 0 {
		i, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return id, err
		}
		id.id = i
	}
	id.nd = n
	return id, nil
}

//String is Stringer method
func (id *Identity) String() string {
	if id.nd <= 0 {
		return ""
	}
	s := strconv.FormatUint(id.id, 10)
	l := len(s)
	if id.nd == l {
		return s
	} else if id.nd < l {
		return s[l-id.nd:]
	}
	return strings.Repeat("0", id.nd-l) + s
}

//Run returns error status in proess.
func Run(args []string) error {
	//initialize flag options
	f := flag.NewFlagSet("fillzero", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: fillzero -n <number of digits> [digit string]")
	}
	n := f.Int("n", 0, "number of digits")

	//parse arguments
	if err := f.Parse(args); err != nil {
		return os.ErrInvalid
	}
	if f.NArg() > 1 {
		f.Usage()
		return errors.New("Too many arguments.")
	}

	// get digit string
	var s string
	if f.NArg() == 1 {
		s = f.Arg(0)
	}

	d, err := IdentityNew(*n, s)
	if err != nil {
		return err
	}
	fmt.Println(d)

	return nil
}

func main() {
	if err := Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
