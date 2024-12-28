package server

import (
	"fmt"
	"log"
	"mynew/src/api"
	"net/http"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware executed")

		clientInfo := r.Header.Get("User-Agent")
		fmt.Println(clientInfo)

		next.ServeHTTP(w, r)
	})
}

func StartServer() {

	api.RegisterEndpoints()

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", middleware(http.DefaultServeMux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
