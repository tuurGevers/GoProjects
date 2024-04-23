package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("the sqrt of %f will result in an irrational number", *e)
}

func Sqrt(x float64) interface{} {
	if x < 0 {
		error := ErrNegativeSqrt(x)
		return &error
	}
	z := 1.0
	prev := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(prev-z) < 0.001 {
			break
		}
		prev = z
	}
	return z
}

func mainSqrt() {
	result := Sqrt(2)
	fmt.Println(result)

	result = Sqrt(-2)
	fmt.Println(result)
}
