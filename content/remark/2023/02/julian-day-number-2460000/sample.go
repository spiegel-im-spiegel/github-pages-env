package main

import (
	"flag"
	"fmt"

	"github.com/goark/koyomi/jdn"
	"github.com/goark/koyomi/value"
)

func main() {
	day := flag.Float64("d", 0, "Julian day")
	flag.Parse()

	dt := jdn.FromJD(*day)
	fmt.Printf("Julian Day: %.3f\n", *day)
	fmt.Printf("Gregorian: %v (%v)\n", dt, dt.In(value.JST))
}
