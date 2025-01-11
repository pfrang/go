package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

func GetWeatherEndpoint() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

		params := r.URL.Query()
		resp, err := getWeather(params)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			return
		}

		handleResponse(w, r, string(body))

	})
}

func getWeather(params url.Values) (*http.Response, error) {
	var err error

	lat, lon := params.Get("lat"), params.Get("lon")

	if lat == "" || lon == "" {
		return nil, errors.New("missing lat or lon parameter")
	}

	url := "https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=" + lat + "&lon=" + lon

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "MyUniqueIdentifier/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err

}
