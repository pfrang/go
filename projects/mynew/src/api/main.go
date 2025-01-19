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

func someFunc() {
	var userSaver UserSaver = func(user string) error {
		fmt.Println("User saved")
		return nil
	}

	userSaver.SaveUser("user")
}

func registerEndpoints(endpoints []Endpoint) {

	for _, endpoint := range endpoints {
		http.Handle(endpoint.Path, endpoint.Handler)
	}
}

type test struct {
	Name string
}

type server struct {
	Name      int
	endpoints []Endpoint
	test
	http.Server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if found := middleware(w, r, s.endpoints); !found {
		return
	}

	http.DefaultServeMux.ServeHTTP(w, r)
}

func StartServer() {
	endpoints := []Endpoint{
		{Path: "/", Description: "Root endpoint", Handler: Root()},
		{Path: "/create-user", Description: "Create a new user", Handler: CreateUserEndpoint()},
		{Path: "/get-weather", Description: "Get the weather", Handler: GetWeatherEndpoint()},
	}

	endpoints = append(endpoints, Endpoint{Path: "/list-endpoints", Description: "List all available endpoints", Handler: listEndpoints(endpoints)})

	registerEndpoints(endpoints)
	s := &server{Server: http.Server{Addr: ":8080"}, endpoints: endpoints}

	fmt.Println("Server running on port 8080...")

	if err := http.ListenAndServe(s.Addr, s); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
