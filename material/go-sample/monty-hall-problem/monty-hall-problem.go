package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectDoor(limit, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < max; i++ {
			ch <- rand.Intn(limit)
		}
	}()
	return ch
}

func challenges(shuldChange bool, dct, max int) int {
	doors := make([]bool, dct, dct) //initilized by false
	doors[0] = true
	ch := selectDoor(len(doors), max)
	count := 0
	for n := range ch {
		result := doors[n]
		if shuldChange {
			result = !result
		}
		if result {
			count++
		}
	}
	return count
}

func main() {
	dct := 100
	max := 100000
	fmt.Println("nochange:", float64(challenges(false, dct, max))/float64(max))
	fmt.Println("  change:", float64(challenges(true, dct, max))/float64(max))
}
