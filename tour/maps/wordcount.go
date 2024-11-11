package main

import (
	"fmt"
	"strings"
)

// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	stringMap := map[string]int{}

	for _, elem := range fields {
		stringMap[elem] += 1
	}
	return stringMap
}

func main() {
	s := WordCount("I am learning am Go!")
	fmt.Println(s)
	// wc.Test(WordCount)
}
