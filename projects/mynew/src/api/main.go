package api

import (
	"fmt"
	"log"
	"net/http"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type UserSaver func(string) error

func (u UserSaver) SaveUser(user string) error {
	fmt.Printf("Saving user: %s\n", user)
	return u(user)
}

type UserData string

func (u UserData) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user")
}

// Define a struct with some fields
type Endpoint struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	Handler     http.Handler
}

// Define a slice of endpoints

func registerEndpoints() {
	var userSaver UserSaver = func(user string) error {
		fmt.Println("User saved")
		return nil
	}

	userSaver.SaveUser("user")

	endpoints := []Endpoint{
		{Path: "/", Description: "Root endpoint", Handler: http.HandlerFunc(Root)},
		{Path: "/create-user", Description: "Create a new user", Handler: CreateUserEndpoint()},
		{Path: "/list-endpoints", Description: "List all available endpoints", Handler: http.HandlerFunc(listEndpoints)},
	}

	for _, endpoint := range endpoints {
		http.Handle(endpoint.Path, endpoint.Handler)
	}
}

type test struct {
	Name string
}

type server struct {
	Name int
	test
	http.Server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if found := middleware(w, r); !found {
		return
	}

	http.DefaultServeMux.ServeHTTP(w, r)
}

func StartServer() {

	s := &server{Addr: ":8080"}

	fmt.Println("Server running on port 8080...")

	if err := http.ListenAndServe(s.Addr, s); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
