package api

import (
	"encoding/json"
	"fmt"
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

func handleResponse(w http.ResponseWriter, r *http.Request, response interface{}) {
	switch v := response.(type) {
	case string:
		w.Write([]byte(v))
	case int, float64:
		w.Write([]byte(fmt.Sprintf("%v", v)))
	default:
		jsonResponse, err := json.Marshal(v)
		if err != nil {
			http.Error(w, "Error processing response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

func StartServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Example usage of handleResponse with different types
		handleResponse(w, r, "Hello, World!")
		// handleResponse(w, r, 123)
		// handleResponse(w, r, map[string]string{"message": "Hello, World!"})
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", middleware(http.DefaultServeMux))
}
