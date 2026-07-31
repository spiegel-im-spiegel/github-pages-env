package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		s, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(s)
	}
}
