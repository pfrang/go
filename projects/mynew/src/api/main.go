package api

import (
	"fmt"
	"log"
	"mynew/src/api/handlers"
	"net/http"
)

// Define a struct with some fields
type Endpoint struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

// Define a slice of endpoints
var endpoints = []Endpoint{
	{Path: "/", Description: "Root endpoint"},
	{Path: "/create-user", Description: "Create a new user"},
	{Path: "/list-endpoints", Description: "List all available endpoints"},
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware executed")

		clientInfo := r.Header.Get("User-Agent")
		fmt.Println(clientInfo)

		next.ServeHTTP(w, r)
	})
}

func StartServer() {

	root, createUser := handlers.Root, handlers.HandlerCreateUser

	for _, endpoint := range endpoints {
		switch endpoint.Path {
		case "/":
			http.HandleFunc(endpoint.Path, root)
		case "/create-user":
			http.HandleFunc(endpoint.Path, createUser)
		}
	}

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", middleware(http.DefaultServeMux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
