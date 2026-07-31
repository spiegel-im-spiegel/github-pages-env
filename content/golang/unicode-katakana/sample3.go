//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

var halfWidthKatakana = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0xff61, 0xff9f, 1},
	},
	R32: []unicode.Range32{},
}

func main() {
	i := 0
	for c := rune(0x30a1); c <= rune(0x1b167); c++ {
		if unicode.In(c, unicode.Katakana) && unicode.In(c, halfWidthKatakana) {
			fmt.Printf("%#U ", c)
			i++
			i &= 0x0f
			if i == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}
