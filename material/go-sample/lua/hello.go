package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/milochristiansen/lua"
	"github.com/milochristiansen/lua/lmodbase"
)

func main() {
	l := lua.NewState()

	err := l.Protect(func() {
		l.Push(lmodbase.Open)
		l.Call(0, 0)
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := l.LoadText(strings.NewReader("print(\"Hello Lua!\")"), "", 0); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := l.PCall(0, 0); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
