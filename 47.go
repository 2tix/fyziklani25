package main

import (
	"fmt"
	"math"
)

func main() {
	ni := 1.33
	nt := 1.00
	steps := 1000000

	criticalAngle := math.Asin(nt / ni)
	dTheta := criticalAngle / float64(steps)
	integral := 0.0

	for i := 0; i < steps; i++ {
		theta := (float64(i) + 0.5) * dTheta
		sinTheta := math.Sin(theta)
		cosTheta := math.Cos(theta)

		sinThetaT := (ni / nt) * sinTheta
		cosThetaT := math.Sqrt(1 - sinThetaT*sinThetaT)

		rs := math.Pow((ni*cosTheta-nt*cosThetaT)/(ni*cosTheta+nt*cosThetaT), 2)
		rp := math.Pow((ni*cosThetaT-nt*cosTheta)/(ni*cosThetaT+nt*cosTheta), 2)

		transmission := 1.0 - 0.5*(rs+rp)
		integral += transmission * sinTheta
	}

	result := 0.5 * integral * dTheta * 100
	fmt.Printf("%.1f%%\n", result)
}
