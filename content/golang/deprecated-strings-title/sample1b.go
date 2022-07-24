//go:build run
// +build run

package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "hello, world!"
	fmt.Println(msg, "->", strings.ToTitle(msg))
}
