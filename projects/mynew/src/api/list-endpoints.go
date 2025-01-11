package api

import (
	"net/http"
)

type EndpointSummary struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

func listEndpoints(endpoints []Endpoint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var summaries []EndpointSummary
		for _, endpoint := range endpoints {
			summaries = append(summaries, EndpointSummary{
				Path:        endpoint.Path,
				Description: endpoint.Description,
			})
		}
		handleResponse(w, r, summaries)
	})
}
