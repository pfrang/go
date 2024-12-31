package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
