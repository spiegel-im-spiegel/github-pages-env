//go:build run
// +build run

package main

import (
	"fmt"
	"unicode"
)

func main() {
	i := 0
	for c := rune(0x30a1); c <= rune(0x1b167); c++ {
		if unicode.In(c, unicode.Katakana) {
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
