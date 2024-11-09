package main

import "fmt"

func main() {
	// Is not called until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// This is called before the world
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
