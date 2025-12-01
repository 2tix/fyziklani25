package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		k  = 100.0
		m  = 0.40
		g  = 9.81
		x0 = 0.02
		h  = 1.0
	)

	Ep := m * g * h
	Ek := 0.5 * k * x0 * x0

	for Ek < Ep {
		Ek *= 2
	}

	x := math.Sqrt(2 * Ek / k)
	fmt.Print(x)
}
