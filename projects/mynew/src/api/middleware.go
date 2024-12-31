package api

import (
	"fmt"
	"net/http"
)

func middleware(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("Middleware executed")

	clientInfo := r.Header.Get("User-Agent")
	fmt.Println(clientInfo)

	found := false
	for _, endpoint := range endpoints {
		if r.URL.Path == endpoint.Path {
			found = true
			break
		}
	}

	if !found {
		http.NotFound(w, r)
	}

	return found
}
