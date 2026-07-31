// +build run

package main

import "fmt"

func main() {
	for d := 0.0; d < 1.0; d += 0.1 {
		fmt.Println(d)
	}
}
