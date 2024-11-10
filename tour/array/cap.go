package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	// Cap always counts from the first element of the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//default value

	var h []int
	fmt.Printf("Default is %v\n", h)
	if h == nil {
		fmt.Println("nil!")
	}
}
