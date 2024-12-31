package api

import (
	"net/http"
)

func listEndpoints(w http.ResponseWriter, r *http.Request) {
	handleResponse(w, r, endpoints)
}
