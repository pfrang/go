package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create an API endpoint

	// resp, err := http.Get("http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Start the HTTP server

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
