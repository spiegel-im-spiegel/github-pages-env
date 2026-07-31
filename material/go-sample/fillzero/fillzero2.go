// +build run

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//regular expression object for Identity class
var re = regexp.MustCompile(`^[0-9]+$`)

//Identity is digit string class
type Identity struct {
	ds string // digit string
	nd int    // number of digits
}

//IdentityNew returns Identity instance.
func IdentityNew(n int, s string) (*Identity, error) {
	d := &Identity{}
	if n <= 0 {
		return d, errors.New("Number of digits is greater than zero.")
	}
	if len(s) > 0 {
		if !re.Copy().MatchString(s) {
			return d, errors.New(s + " is not digit string.")
		}
	}
	d.nd = n
	d.ds = s
	return d, nil
}

//String is Stringer method
func (s *Identity) String() string {
	if s.nd <= 0 {
		return ""
	}
	l := len(s.ds)
	if s.nd == l {
		return s.ds
	} else if s.nd < l {
		return s.ds[l-s.nd:]
	}
	return strings.Repeat("0", s.nd-l) + s.ds
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
