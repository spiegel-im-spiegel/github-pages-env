//go:build run
// +build run

package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	msg := "'n"
	casers := []cases.Caser{
		cases.Title(language.English),
		cases.Title(language.Dutch),
	}
	for _, caser := range casers {
		fmt.Println(msg, "->", caser.String(msg))
	}
}
