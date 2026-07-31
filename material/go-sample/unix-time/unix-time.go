package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Date(2016, time.December, 31, 23, 59, 59, 0, time.UTC)
	t2 := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("t1 = %v (%v)\n", t1, t1.Unix())
	fmt.Printf("t2 = %v (%v)\n", t2, t2.Unix())
	fmt.Printf("t2 - t1 = %v\n", t2.Sub(t1))
	fmt.Printf("t2 - t1 = %v sec\n", t2.Unix()-t1.Unix())
}
