// +build run

package main

import (
	"fmt"
	"unicode"
)

func check(r rune) string {
	switch {
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
	text := string([]byte{0xef, 0xbb, 0xbf, 0xe3, 0x82, 0x84, 0x09, 0xe3, 0x81, 0x82})
	fmt.Println(text)
	for _, c := range text {
		fmt.Printf("%#U (%v)\n", c, check(c))
	}
}
