package main

import (
	"encoding/json"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(map[string]string{"message": "GET example"})
		case http.MethodPost:
			json.NewEncoder(w).Encode(map[string]string{"message": "POST example"})
		case http.MethodPut:
			json.NewEncoder(w).Encode(map[string]string{"message": "PUT example"})
		case http.MethodDelete:
			json.NewEncoder(w).Encode(map[string]string{"message": "DELETE example"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
