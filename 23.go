package main

import (
	"fmt"
	"math"
)

func main() {
	h := 1.00
	d := 20.0

	coefficient := (2 * math.Pi * h) / d

	fmt.Println(coefficient)
}
