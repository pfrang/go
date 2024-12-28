package api

import (
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

func RegisterEndpoints() {
	root, createUser := handlers.Root, handlers.HandlerCreateUser

	for _, endpoint := range endpoints {
		switch endpoint.Path {
		case "/":
			http.HandleFunc(endpoint.Path, root)
		case "/create-user":
			http.HandleFunc(endpoint.Path, createUser)
		}
	}
}
