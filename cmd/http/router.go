package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/example", exampleHandler).
		Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)

	http.Handle("/", router)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		handleGetExample(w, r)
	case http.MethodPost:
		handlePostExample(w, r)
	case http.MethodPut:
		handlePutExample(w, r)
	case http.MethodDelete:
		handleDeleteExample(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
