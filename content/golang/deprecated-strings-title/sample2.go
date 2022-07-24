//go:build run
// +build run

package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	msg := "hello, world!"
	fmt.Println(msg, "->", cases.Title(language.Und).String(msg))
}
