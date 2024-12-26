package main

import "fmt"

type Vertec struct {
	X, Y float64
}

func (v *Vertec) ModifyX() {
	v.X = v.X * v.Y
}

func main() {

	var v Vertec
	v = Vertec{3, 4}

	v.ModifyX()

	fmt.Println(v.X)

}
