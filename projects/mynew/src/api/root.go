package api

import (
	"net/http"
)

func Root() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleResponse(w, r, "Hello, World!")
	})
}
