package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) (string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Return the content as a string
	return string(content), nil
}

func main() {
	filename := "input.txt" // Replace with your file name
	content, err := readFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// fmt.Println("File content:")
	// fmt.Println(content)

	// Split the content by whitespace to get individual numbers
	parts := strings.Fields(content)
	fmt.Println(parts[0])

	var leftColumn []int
	var rightColumn []int

	for index, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		if index%2 == 0 {
			rightColumn = append(rightColumn, num)
		} else {
			leftColumn = append(leftColumn, num)
		}
	}

	// Sort the arrays
	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	var sum = 0

	for index, part := range leftColumn {
		num := part - rightColumn[index]
		sum += int(math.Abs(float64(num)))
	}

	fmt.Println(sum)

	// Convert the split strings to integers and add to an array
	// var numbers []int
	// for _, part := range parts {
	// 	num, err := strconv.Atoi(part)
	// 	if err != nil {
	// 		log.Fatalf("Error converting string to int: %v", err)
	// 	}
	// 	numbers = append(numbers, num)
	// }

	// Print the numbers array
	// fmt.Println("Numbers array:")
	// fmt.Println(numbers[0])
}
