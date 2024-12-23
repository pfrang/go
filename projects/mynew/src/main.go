package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
	h.ServeHTTP(w, r)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware executed")

		// some inf oiabout client

		clientInfo := r.Header.Get("User-Agent")
		fmt.Println(clientInfo)

		next.ServeHTTP(w, r)
	})
}

func main() {
	// create an API endpoint

	// resp, err := http.Get("http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Start the HTTP server

	// handler := &MyHandler{}

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", middleware(http.DefaultServeMux))

}
