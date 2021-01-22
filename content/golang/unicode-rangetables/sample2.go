// +build run

package main

import (
	"fmt"
	"unicode"
)

func check(r rune) string {
	switch {
	case unicode.Is(unicode.Sc, r):
		return "Symbol/currency"
	case unicode.Is(unicode.Sk, r):
		return "Symbol/modifier"
	case unicode.Is(unicode.Sm, r):
		return "Symbol/math"
	case unicode.Is(unicode.So, r):
		return "Symbol/other"
	case unicode.Is(unicode.Variation_Selector, r):
		return "Variation Selector"
	case unicode.Is(unicode.Join_Control, r):
		return "Join Control"
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
	text := "🙇‍♂️"
	for _, c := range text {
		fmt.Printf("%#U (%v)\n", c, check(c))
	}
}
