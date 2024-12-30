package api

import (
	"fmt"
	"net/http"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Middleware executed")

	clientInfo := r.Header.Get("User-Agent")
	fmt.Println(clientInfo)
}
