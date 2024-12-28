package api

import (
	"net/http"
)

// Define a struct with some fields
type Endpoint struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

// Define a slice of endpoints
var Endpoints = []Endpoint{
	{Path: "/", Description: "Root endpoint"},
	{Path: "/create-user", Description: "Create a new user"},
	{Path: "/list-endpoints", Description: "List all available endpoints"},
}

func RegisterEndpoints() {

	for _, endpoint := range Endpoints {
		switch endpoint.Path {
		case "/":
			http.HandleFunc(endpoint.Path, Root)
		case "/create-user":
			http.HandleFunc(endpoint.Path, CreateUser)
		}
	}
}
