package main

import (
	"fmt"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/jdn"
)

func main() {
	num := 2460000.0
	dt := jdn.FromJD(num)
	mjd := jdn.GetMJD(dt)
	fmt.Printf("Julian Day Number: %.3f (%v)\n", num, mjd.FloatString(3))
	fmt.Printf("Gregorian: %v (%v)\n", dt, dt.In(koyomi.JST))
}
