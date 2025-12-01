package main

import (
	"fmt"
	"math"
)

func main() {
	h := 2.44
	m := 140.0
	c := 1.05
	rho := 1.225
	g := 9.81

	v := math.Sqrt((2 * m * g) / (rho * c * math.Pow(h, 2)))

	fmt.Printf("%.1f m/s\n", v)
}
