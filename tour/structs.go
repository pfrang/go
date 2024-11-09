package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var Vertex1 Vertex
	Vertex1.X = 4
	Vertex1.Y = 5
	fmt.Println(Vertex1)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	// is equal to
	(*p).X = 1e9
	fmt.Println(v)
}
