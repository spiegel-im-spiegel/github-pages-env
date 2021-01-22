// +build run

package main

import (
	"fmt"
	"unicode"
)

func check(r rune) string {
	switch {
	case unicode.Is(unicode.Katakana, r):
		return "Katakana"
	case unicode.Is(unicode.Hiragana, r):
		return "Hiragana"
	case unicode.Is(unicode.Lm, r):
		return "Letter/modifier"
	case unicode.Is(unicode.Lo, r):
		return "Letter"
	case unicode.Is(unicode.Mc, r):
		return "Mark/spacing combining"
	case unicode.Is(unicode.Me, r):
		return "Mark/enclosing"
	case unicode.Is(unicode.Mn, r):
		return "Mark/nonspacing"
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
	text := "ペンギンペンギンﾍﾟﾝｷﾞﾝ"
	for _, c := range text {
		fmt.Printf("%#U (%v)\n", c, check(c))
	}
}
