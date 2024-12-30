package server

import (
	"fmt"
	"log"
	"mynew/src/api"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.Middleware(w, r)

	http.DefaultServeMux.ServeHTTP(w, r)
}

func StartServer() {

	s := &server{addr: ":8080"}
	api.RegisterEndpoints()

	fmt.Println("Server running on port 8080...")

	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
