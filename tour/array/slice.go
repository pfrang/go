package main

import "fmt"

func main() {

	primesSlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primesSlice)

	slicedPrime := primesSlice[1:4]
	fmt.Println(slicedPrime)

	// Slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	// Other slices that share the same underlying array will see those changes.

	slicedPrime[0] = 100

	// 3 from original array has now bee nset to 100
	fmt.Println(primesSlice)

	primesSliceFirstTwo := primesSlice[:2]
	fmt.Println(primesSliceFirstTwo)
	primesSliceLastTwo := primesSlice[len(primesSlice)-2:]
	fmt.Println(primesSliceLastTwo)
}
