//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func main() {
	for c := rune(0x3041); c <= rune(0x1f200); c++ {
		if unicode.In(c, unicode.Hiragana) {
			n := norm.NFKC.String(string(c))
			if n != string(c) {
				fmt.Printf("%#U -> %s\n", c, n)
			}
		}
	}
}
