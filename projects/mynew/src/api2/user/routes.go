package user

import "net/http"

func RegisterEndpoints() {
	http.HandleFunc("/create-user", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}
