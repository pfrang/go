package main

import (
	"basics/mylib"
	"basics/packages/importexport"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	_ = mylib.Add(1, 2)

	if len(args) == 0 {
		fmt.Println("Usage: go run index.go <arg1> <arg2> ...")
		return
	}

	input := args[0]

	fmt.Println("Input:", input)

	switch input {
	case "const":
		constFunc()
		return
	case "strings":
		strings()
		return
	case "import":
		importexport.SomeImportFunc()
		fmt.Printf("%d arguments provided\n", len(args))
	}

	// A switch statement based on cli arguments
}
