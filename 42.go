package main

import (
	"fmt"
	"math"
)

func main() {
	ni := 1.33
	nt := 1.00

	sinThetaC := nt / ni
	cosThetaC := math.Sqrt(1 - math.Pow(sinThetaC, 2))

	solidAngleFraction := (1 - cosThetaC) / 2.0

	R := math.Pow((ni-nt)/(ni+nt), 2)
	T := 1.0 - R

	percentage := solidAngleFraction * T * 100

	fmt.Printf("%.1f%%\n", percentage)
}
