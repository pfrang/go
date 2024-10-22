package main

import (
	"fmt"
	"unicode/utf8"
)

//  The %v variant prints the Go syntax representation of a value, it's a nice default.

func strings() {

	s := fmt.Sprintf("I am %v years old", 10)
	// I am 10 years old

	s2 := fmt.Sprintf("I am %v years old", "way too many")
	// I am way too many years old

	fmt.Println(s)
	fmt.Println(s2)

	const name2 = "Saul Goodman"
	const openRate = 30.5

	// ?

	msg := fmt.Sprintf("Hi %v, your open rate is %v percent\n", name2, openRate)

	// fmt.Println automatically ends with newline, fmt.Print does not

	fmt.Print(msg)

	// Runes

	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
