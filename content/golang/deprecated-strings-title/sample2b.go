//go:build run
// +build run

package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	msgs := []string{
		"hello, world!",
		"HELLO, WORLD!",
	}
	for _, msg := range msgs {
		fmt.Println(msg, "->", cases.Title(language.Und, cases.NoLower).String(msg))
	}
}
