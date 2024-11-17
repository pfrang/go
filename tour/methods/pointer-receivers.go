package main

import (
	"fmt"
)

type VertexPointer struct {
	X, Y float64
}

func (v VertexPointer) Abs() float64 {
	// 30 * 40 = 1200
	// or 3 * 4 = 12
	// depending on pointer vs value receiver
	return v.X * v.Y
}

// Pointer receiver instead of value receiver
// modifies original value
// try removing * and see how program behaves
func (v *VertexPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexPointer{3, 4}
	// even though v isnt a pointer, it is converted to a pointer
	// when calling Scale due to the pointer receiver
	v.Scale(10)
	fmt.Println(v.Abs())
}
