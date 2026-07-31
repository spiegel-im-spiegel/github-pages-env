// +build run

package main

import "fmt"

func main() {
	var d float64 = 0.0
	for i := 0; i < 100000000; i++ {
		d += 1.0
	}
	fmt.Println(d)
}
