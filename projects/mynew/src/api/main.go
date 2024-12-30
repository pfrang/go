package api

import (
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

// func listEndpoints(w http.ResponseWriter, r *http.Request) {
// 	utils.HandleResponse(w, r, Endpoints)
// }

// Define a struct with some fields
type Endpoint struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	Handler     Handler
}

// Define a slice of endpoints
var Endpoints = []Endpoint{
	{Path: "/", Description: "Root endpoint", Handler: Root},
	{Path: "/create-user", Description: "Create a new user", Handler: CreateUser},
	// {Path: "/list-endpoints", Description: "List all available endpoints", Handler: listEndpoints},
}

func RegisterEndpoints() {
	for _, endpoint := range Endpoints {
		http.HandleFunc(endpoint.Path, endpoint.Handler)
	}
}
