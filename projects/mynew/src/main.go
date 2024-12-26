package main

import (
	"mynew/src/api" // Import the api package
)

func main() {
	// create an API endpoint

	// resp, err := http.Get("http://localhost:8080")

	// Start the HTTP server

	// Register the route from the api package
	api.StartServer()
	// handler := &MyHandler{}

}
