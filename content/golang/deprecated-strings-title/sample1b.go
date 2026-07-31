//go:build run
// +build run

package main

import (
	"fmt"
	"strings"
)

func main() {
	msgs := []string{
		"hello, world!",
		"HELLO, WORLD!",
	}
	for _, msg := range msgs {
		fmt.Println(msg, "->", strings.ToTitle(msg))
	}
}
