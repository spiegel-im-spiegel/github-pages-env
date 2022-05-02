package main

import (
	"fmt"

	"github.com/goark/koyomi"
	"github.com/goark/koyomi/jdn"
)

func main() {
	var num int64 = 2460000
	dt := jdn.FromJD(num)
	mjd := jdn.GetMJD(dt)
	fmt.Printf("Julian Day Number: %v (%v)\n", num, mjd.FloatString(3))
	fmt.Printf("Gregorian: %v (%v)\n", dt, dt.In(koyomi.JST))
}
