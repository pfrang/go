package main

import "time"

// Constants can be primitive types like strings, integers, booleans and floats. They can not be more complex types like slices, maps and structs.
const someConst = 2

// That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:
// the current time can only be known when the program is running
// const currentTime = time.Now()
var currentTime = time.Now()

// BUt consts that can be computed at build is ok

const someMath = 25 * 35

func constExample() {

	const someConst = 2
	// someConst = 3 // This will throw an error
	println(someConst)
}
