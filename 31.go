package main

import (
	"fmt"
	"math"
)

func main() {
	r := 3.00e-3
	h := 5.00e-3
	f := 65.0
	torqueRes := 2.00e-9
	conc := 2.00e-4
	molarMass := 39.948e-3

	omega := 2 * math.Pi * f

	A := 0.5 * molarMass * conc * h * r * r

	B := molarMass * conc * h * r * r * r * omega

	C := -torqueRes

	discriminant := B*B - 4*A*C
	if discriminant < 0 {
		fmt.Println("Physics doesn't work.")
		return
	}

	vGas := (B + math.Sqrt(discriminant)) / (2 * A)

	nDot := conc * r * h * vGas

	fmt.Printf("%.1e\n", nDot)
}
