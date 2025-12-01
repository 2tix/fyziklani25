package main

import (
	"fmt"
	"math"
)

func main() {
	c := 3.0e8
	f := 50.0e6
	h := 20000.0 * 0.3048
	tDelay := 0.10e-3
	df := 190.0

	d := c * tDelay / 2
	x := math.Sqrt(d*d - h*h)
	vr := (c * df) / (2 * f)
	v := vr * d / x
	result := x / v

	fmt.Printf("%.0f s\n", result)
}
