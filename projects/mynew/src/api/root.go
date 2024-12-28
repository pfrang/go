package api

import (
	"mynew/src/api/utils"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	utils.HandleResponse(w, r, "Hello, World!")
}
