// +build run

package main

import (
	_ "embed"
	"fmt"
)

//go:embed html/index.html
var hello string

func main() {
	fmt.Println(hello)
}
