package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice(a)

	// To specify the cap, pass 3rd argument
	b := make([]int, 0, 5)
	printSlice(b)
	b = b[:5]
	printSlice(b)
	b = b[:2]
	printSlice(b)
	b = b[:5]
	printSlice(b)
}

func printSlice(s []int) {
	if len(s) > 0 {
		fmt.Printf("len=%d cap=%d, address of slice header=%p, address of first element=%p %v\n", len(s), cap(s), &s, &s[0], s)
	} else {
		fmt.Printf("len=%d cap=%d, address of slice header=%p, no elements %v\n", len(s), cap(s), &s, s)
	}
}
