package api

import (
	"mynew/src/api/utils"
	"net/http"
)

func ListEndpoints(w http.ResponseWriter, r *http.Request) {
	utils.HandleResponse(w, r, Endpoints)
}
