package main

import "fmt"

type I2 interface {
	M2()
}

type T2 struct {
	S string
}

func main() {
	var i I2
	i = &T2{"Hello"}
	i.M2()
	describe2(i)
}

func (i *T2) M2() {
	fmt.Println(i)
}

func describe2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
