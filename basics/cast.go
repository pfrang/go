package main

import (
	"fmt"
	"strconv"
)

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)

	fmt.Println("Your account has existed for", accountAgeInt, "years")

	var myAge uint16 = 25

	// inputFunc(myAge)
	fmt.Println("Your account has existed for", myAge, "years")
	concatenation()
}

func inputFunc(input int) {
	fmt.Println("Your account has existed for", input, "years")
}

func concatenation() {
	username := "Pjaff"
	password := string(1245)

	fmt.Println("Authorization: Basic", username+":"+password)

	password2 := strconv.Itoa(1245) // Convert integer to string

	fmt.Println("Authorization: Basic", username+":"+password2)
}
