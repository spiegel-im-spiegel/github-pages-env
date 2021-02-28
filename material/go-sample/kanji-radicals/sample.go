package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	for r := rune(0x2e80); r <= 0x2eff; r++ {
		rr := []rune(norm.NFKC.String(string([]rune{r})))
		if r != rr[0] {
			fmt.Printf("%#U -(NFKC)-> %#U\n", r, rr[0])
		}
	}
}
