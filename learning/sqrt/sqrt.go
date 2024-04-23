package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(prev-z) < 0.001 {
			break
		}
		prev = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
