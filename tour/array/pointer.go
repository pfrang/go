package main

import "fmt"

func main() {
	var s = [4]int{1, 2, 3, 4}
	var pS = &s
	fmt.Printf("Address of array: %p, and it points to %v\n", &s, (*pS))

	var a = []int{1, 2, 3, 4}
	a = a[:5]
	fmt.Printf("Address of slice: %p, and it points to %v\n", &a, a)
}
