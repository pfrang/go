package api2

import (
	"fmt"
	"log"
	"mynew/src/api2/user"
	"net/http"
)

type MyServer http.Server

func StartServer() {

	s := &MyServer{Addr: ":8080"}

	user.RegisterEndpoints()

	fmt.Println("Server started at localhost:8080...")

	if err := http.ListenAndServe(s.Addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
