package main

import (
	"fmt"
	"math"
)

type Abs2Type func(int) int
type Vertex struct {
	X, Y float64
	Abs2 Abs2Type
}

// a receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var Abs2 Abs2Type = func(n int) int {
	return n * n
}

func main() {
	v := Vertex{3, 4, Abs2}
	fmt.Println(v.Abs())
	// or with regular function
	fmt.Println(Abs(v))

	fmt.Println((v.Abs2(6)))
}
