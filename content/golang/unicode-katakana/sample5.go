//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

func main() {
	i := 0
	for c := rune(0x3041); c <= rune(0x1f200); c++ {
		if unicode.In(c, unicode.Hiragana) {
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
