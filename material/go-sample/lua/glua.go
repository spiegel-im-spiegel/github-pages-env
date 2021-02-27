// +build run

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/milochristiansen/lua"
	"github.com/milochristiansen/lua/lmodbase"
	"github.com/milochristiansen/lua/lmodmath"
	"github.com/milochristiansen/lua/lmodpackage"
	"github.com/milochristiansen/lua/lmodstring"
	"github.com/milochristiansen/lua/lmodtable"
)

func Run(r io.Reader) error {
	l := lua.NewState()

	err := l.Protect(func() {
		l.Push(lmodbase.Open)
		l.Call(0, 0)

		l.Push(lmodmath.Open)
		l.Call(0, 0)

		l.Push(lmodpackage.Open)
		l.Call(0, 0)

		l.Push(lmodstring.Open)
		l.Call(0, 0)

		l.Push(lmodtable.Open)
		l.Call(0, 0)
	})
	if err != nil {
		return err
	}

	if err := l.LoadText(r, "", 0); err != nil {
		return err
	}
	if err := l.PCall(0, 0); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	if err := Run(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
