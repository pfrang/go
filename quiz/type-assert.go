package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

// THis function is a factory function that returns a Developer struct
func GetDeveloper(name interface{}, age interface{}) Developer {
	// Implement this function
	return Developer{Name: name.(string), Age: age.(int)}
}

func main() {
	fmt.Println("Hello World")

	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
