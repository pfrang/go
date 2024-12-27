package handlers

import (
	"mynew/src/api/utils"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func Root(w http.ResponseWriter, r *http.Request) {
	utils.HandleResponse(w, r, "Hello, World!")
}
