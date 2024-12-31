package api

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	handleResponse(w, r, "Hello, World!")
}
