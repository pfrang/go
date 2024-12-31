package api

import (
	"fmt"
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

func (e *Handler) LogLoL() {
	fmt.Println("LOL")
}

// Define a struct with some fields
type Endpoint struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	Handler     Handler
}

// Define a slice of endpoints
var endpoints = []Endpoint{
	{Path: "/", Description: "Root endpoint", Handler: Root},
	{Path: "/create-user", Description: "Create a new user", Handler: CreateUser},
	// {Path: "/list-endpoints", Description: "List all available endpoints", Handler: listEndpoints},
}

func registerEndpoints() {
	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint.Path, endpoint.Handler)
	}
}

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if found := middleware(w, r); !found {
		return
	}

	http.DefaultServeMux.ServeHTTP(w, r)
}

func StartServer() {

	s := &server{addr: ":8080"}
	registerEndpoints()

	fmt.Println("Server running on port 8080...")

	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
