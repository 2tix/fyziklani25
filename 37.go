package main

import (
	"fmt"
	"math"
)

func main() {
	c := 299792458.0
	Q := 200.0 / 1000.0

	t := 3 * math.Sqrt((c*math.Pow(math.Pi, 3))/math.Pow(Q, 3))

	fmt.Println(t)
}
