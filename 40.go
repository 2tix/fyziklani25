package main

import (
	"fmt"
	"math"
)

func main() {
	S := 42.0 * 1e-4
	s := 4.2 * 1e-6
	h := 0.05
	rhoW := 1000.0
	rhoO := 910.0
	rhoE := 789.0
	g := 9.81

	hW := h
	hO := h
	hE := h
	t := 0.0
	dt := 0.001

	for hW > 0 {
		P := g * (rhoW*hW + rhoO*hO + rhoE*hE)
		v := math.Sqrt(2 * P / rhoW)
		dh := (s / S) * v * dt
		if hW <= dh {
			t += hW / ((s / S) * v)
			hW = 0
			break
		}
		hW -= dh
		t += dt
	}

	for hO > 0 {
		P := g * (rhoO*hO + rhoE*hE)
		v := math.Sqrt(2 * P / rhoO)
		dh := (s / S) * v * dt
		if hO <= dh {
			t += hO / ((s / S) * v)
			hO = 0
			break
		}
		hO -= dh
		t += dt
	}

	fmt.Printf("%.0f s\n", t)
}
