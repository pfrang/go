package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	// Address
	println(&s)
	s = append(s, 5, 6, 7)
	// Keeps the same address
	println(&s)

	println(s[4])

	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...)
	fmt.Println(a)

}
