package main

import (
	"fmt"
	"math"
)

func main() {
	S := 16.0
	s := 0.40
	H := 8.0
	h := 5.0
	g := 981.0

	val := (S / s) * math.Sqrt((2*(H-h)*(1-math.Pow(s/S, 2)))/g)

	fmt.Printf("%.4f s\n", val)
}
