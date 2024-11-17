package main

import (
	"fmt"
)

type VertexValuePointer struct {
	X, Y float64
}

func (v VertexValuePointer) ValueAbs() float64 {
	return v.X * v.Y
}

func (v VertexValuePointer) ValueScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ValueFunc(v VertexValuePointer) float64 {
	return v.X * v.Y
}

func main() {
	v := VertexValuePointer{3, 4}
	// even though v isnt a pointer, it is converted to a pointer
	// when calling Scale due to the pointer receiver
	v.ValueScale(10)
	p := &v
	fmt.Println(v.ValueAbs())
	fmt.Println(p.ValueAbs())
	ValueFunc(v)
	ValueFunc(p) // Not allowed for raw functions, but for method receivers both are ok
}
