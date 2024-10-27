package main

import "fmt"

func main() {

	x, _ := getPoint()
	fmt.Println(x)
}

func getPoint() (x int, y int) {
	return 3, 4
}

// ignore y value
