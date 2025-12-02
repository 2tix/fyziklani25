package main

import (
	"fmt"
	"math"
)

func main() {
	kappa := 1.4
	expansionRatio := 2.0
	tempTargetRatio := 2.0

	tInterFactor := tempTargetRatio * math.Pow(expansionRatio, 1.0-kappa)
	qOptFactor := tInterFactor - 1.0
	qStdFactor := tempTargetRatio - 1.0

	percentage := (qOptFactor / qStdFactor) * 100
	fmt.Printf("%.0f%%\n", percentage)
}
