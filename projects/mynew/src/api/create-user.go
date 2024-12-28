package api

import (
	"encoding/json"
	"mynew/src/api/utils"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON structure", http.StatusBadRequest)
		return
	}

	utils.HandleResponse(w, r, user)

}
