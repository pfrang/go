package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var genericArray [2]interface{}
	genericArray[0] = "Hello"
	genericArray[1] = 1
	fmt.Println(genericArray)

	//Slices
	primesSlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primesSlice)

	slicedPrime := primesSlice[1:4]
	fmt.Println(slicedPrime)

}
