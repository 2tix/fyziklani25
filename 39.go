package main

import (
	"fmt"
	"math"
)

func main() {
	const G = 6.67430e-11
	const M = 1.989e30
	const R = 6.957e8
	const AU = 1.495978707e11
	const v0 = 70.0

	var hits int

	rowHeight := math.Sqrt(12*12 - 6*6)

	rows := []struct {
		count   int
		yOffset float64
		xStart  float64
	}{
		{4, 0.0, 13.0},
		{3, -rowHeight, 19.0},     // 13 + 6
		{2, -2 * rowHeight, 25.0}, // 13 + 12
		{1, -3 * rowHeight, 31.0}, // 13 + 18
	}

	for _, row := range rows {
		y := (730.0 + row.yOffset) * AU
		for i := 0; i < row.count; i++ {
			x := (row.xStart + float64(i)*12.0) * AU

			r0 := math.Sqrt(x*x + y*y)

			E := 0.5*v0*v0 - G*M/r0

			h := x * v0

			term := math.Sqrt(G*G*M*M + 2*E*h*h)
			rMin := (-G*M + term) / (2 * E)

			if rMin <= R {
				hits++
			}
		}
	}

	fmt.Println(hits)
}
