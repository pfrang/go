package main

import "fmt"

type I interface {
	M()
}

type G interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
	var i2 G = T{"hello"}
	i2.M()
}
