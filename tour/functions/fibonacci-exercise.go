// Exercise: Fibonacci closure
// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci
// numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

func fibonacci() func() []int {
	fibonacciSequence := []int{}
	return func() []int {
		for i := 0; i < 10; i++ {
			if i == 0 || i == 1 {

				fibonacciSequence = append(fibonacciSequence, i)
				continue
			}
			add := fibonacciSequence[i-2] + fibonacciSequence[i-1]
			fibonacciSequence = append(fibonacciSequence, add)

		}

		return fibonacciSequence
	}
}
func main() {
	sequence := fibonacci()
	fmt.Println(sequence())
}
