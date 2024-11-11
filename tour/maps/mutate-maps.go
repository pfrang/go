package main

import "fmt"

func main() {
	a := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	a["four"] = 4
	delete(a, "four")
	fmt.Println(a)

	elem, ok := a["five"]
	fmt.Printf("Elem %d is present? %t ", elem, ok)
}
