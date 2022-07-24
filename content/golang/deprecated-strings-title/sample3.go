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
	tags := []language.Tag{
		language.English,
		language.Dutch,
	}
	for _, tag := range tags {
		fmt.Println(tag, ":", msg, "->", cases.Title(tag).String(msg))
	}
}
