package main

import (
	"fmt"
	"math"
)

func calculateDiscountRate(messagesSent int) float64 {
	// ?

	if messagesSent > 2500 {
		return 0.2
	} else if messagesSent > 1000 {
		return 0.1
	} else {
		return 0
	}
}

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}
