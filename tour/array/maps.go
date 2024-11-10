package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

//typed keys

type Key string

const (
	First  Key = "First"
	Second Key = "Second"
)

var m2 map[Key]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["hei"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["hei"])

	// Assigning the string to consts
	m2 = make(map[Key]Vertex)
	m2[First] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m2)

	var m3 map[int]string
	m3 = make(map[int]string)
	m3[1] = "one"
	m3[1000] = "thousand"
	fmt.Println(len(m3))

}
