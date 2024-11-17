package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (n MyFloat) Hei(n2 int) int {
	return int(n) * n2
}

// Methods can be attached to any type, here a float64
func main() {
	a := int(42314)
	fmt.Println(a)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())
	fmt.Println(f.Hei(5))
}
