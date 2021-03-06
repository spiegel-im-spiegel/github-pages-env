package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	dt, err := time.Parse("2006-01-02", "2022-03-31")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%v is %d days from today\n", dt.Format("2006-01-02"), time.Until(dt)/(time.Hour*24))
}
