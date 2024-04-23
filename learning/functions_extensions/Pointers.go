package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// mainPointers is a function that demonstrates the usage of pointers in Go.
func mainPointers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
