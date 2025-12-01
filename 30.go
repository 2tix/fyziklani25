package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		R       = 6371.0
		DayMins = 1440.0
		OmegaE  = 2 * math.Pi / DayMins
		V       = 300.0 / 60.0
		Alpha   = 30.0 * math.Pi / 180.0
		TGap    = 28.0
	)

	Vh := V * math.Cos(Alpha)
	Vv := V * math.Sin(Alpha)

	f := func(t float64) float64 {
		h := Vv * t
		delta := math.Acos(R / (R + h))
		thetaTravel := (Vh * t) / R
		thetaSun := (OmegaE * TGap) - (OmegaE * t) - thetaTravel
		return thetaSun - delta
	}

	low := 0.0
	high := 28.0
	epsilon := 1e-9

	for i := 0; i < 1000; i++ {
		mid := (low + high) / 2
		val := f(mid)

		if math.Abs(val) < epsilon {
			fmt.Printf("%.4f\n", mid)
			return
		}

		if f(low)*val < 0 {
			high = mid
		} else {
			low = mid
		}
	}

	fmt.Printf("%.4f\n", (low+high)/2)
}
