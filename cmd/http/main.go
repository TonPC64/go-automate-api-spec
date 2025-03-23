package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	setupRoutes() // Add this line to setup routes

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	http.ListenAndServe(":8084", nil)
}
