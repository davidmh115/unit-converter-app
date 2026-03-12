package main

import (
	"encoding/json"
	"net/http"
)

type ConvertRequest struct {
	Kilometers float64 `json:"kilometers"`
}

type ConvertResponse struct {
	Kilometers float64 `json:"kilometers"`
	Miles      float64 `json:"miles"`
}

func kilometersToMiles(km float64) float64 {
	return km * 0.621371
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid JSON"}`))
		return
	}

	resp := ConvertResponse{
		Kilometers: req.Kilometers,
		Miles:      kilometersToMiles(req.Kilometers),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/convert", convertHandler)

	http.ListenAndServe(":8080", nil)
}
