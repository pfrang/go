package main

// take in cli arguments
import (
	"fmt"
	"os"
)

// main function
func main() {

	// get the first argument
	// if there is no argument, print usage
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run get-monthly-price.go <tier>")
		return
	}

	// get the first argument
	tier := os.Args[1]

	fmt.Println("Tier:", tier)

	output := getMonthlyPrice(tier)

	fmt.Println(output)

}

func getMonthlyPrice(tier string) int {
	const dollarToPenny = 100

	switch tier {
	case "basic":
		return 100 * dollarToPenny
	case "premium":
		return 150 * dollarToPenny
	case "enterprise":
		return 500 * dollarToPenny
	default:
		return 0
	}
}
