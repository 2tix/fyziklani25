package main

import (
	"fmt"
	"math"
)

func main() {
	E := 5.12e-10
	R := 1.11
	c := 2.99792458e8
	m := 1.6726219e-27
	q := 1.60217663e-19

	restE := m * c * c
	p := math.Sqrt(E*E-restE*restE) / c
	B := p / (q * R)

	fmt.Printf("%.2f T\n", B)
}
