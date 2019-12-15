package main

import (
	"fmt"
	"strings"
)

func main() {
	lefts := []string{"go", "ｇｏ"}
	rights := []string{"Go", "GO", "go", "Ｇｏ", "ＧＯ", "ｇｏ"}

	for _, left := range lefts {
		for _, right := range rights {
			fmt.Printf("%s == %s : %v\n", left, right, (left == strings.ToLower(right)))
		}
	}
}
