package main

import (
	"fmt"
	"math"
)

func main() {
	sunAngleDeg := 32.0
	launchHeight := 0.70
	launchAngleDeg := 38.0
	distance := 25.0
	planeTiltDeg := 23.0
	g := 9.81

	sunAngle := sunAngleDeg * math.Pi / 180
	launchAngle := launchAngleDeg * math.Pi / 180
	planeTilt := planeTiltDeg * math.Pi / 180

	cosAlpha := math.Cos(launchAngle)
	sinAlpha := math.Sin(launchAngle)
	tanAlpha := math.Tan(launchAngle)

	numerator := g * distance * distance
	denominator := 2 * (launchHeight + distance*tanAlpha) * cosAlpha * cosAlpha
	v0Squared := numerator / denominator
	v0 := math.Sqrt(v0Squared)

	vx := v0 * cosAlpha
	vy0 := v0 * sinAlpha

	vzLanding := vy0 - g*(distance/vx)

	calcShadowSpeed := func(vz float64) float64 {
		cotSigma := 1.0 / math.Tan(sunAngle)
		vsx := vx - vz*cotSigma*math.Cos(planeTilt)
		vsy := -vz * cotSigma * math.Sin(planeTilt)
		return math.Sqrt(vsx*vsx + vsy*vsy)
	}

	speedLaunch := calcShadowSpeed(vy0)
	speedLand := calcShadowSpeed(vzLanding)

	maxSpeed := math.Max(speedLaunch, speedLand)

	fmt.Printf("%.0f m/s\n", maxSpeed)
}
