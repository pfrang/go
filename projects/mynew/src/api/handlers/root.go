package handlers

import (
	"net/http"
	"src/api/utils"
)

func Root(w http.ResponseWriter, r *http.Request) {
	utils.HandleResponse(w, r, "Hello, World!")
}
