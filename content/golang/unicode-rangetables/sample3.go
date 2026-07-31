// +build run

package main

import (
	"fmt"
	"unicode"
)

func check(r rune) string {
	switch {
	case unicode.Is(unicode.Radical, r):
		return "Radical"
	case unicode.Is(unicode.Ideographic, r):
		return "Ideographic"
	case unicode.IsGraphic(r):
		return "Graphic"
	case unicode.IsControl(r):
		return "Latin1 Control"
	case unicode.Is(unicode.C, r):
		return "Unicode Control"
	}
	return "Unknown"
}

func main() {
	text := "⽟玉"
	for _, c := range text {
		fmt.Printf("%#U (%v)\n", c, check(c))
	}
}
