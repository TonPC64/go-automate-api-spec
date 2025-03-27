package main

import (
	"encoding/json"
	"net/http"
)

// handleGetExample godoc
//
//	@Summary		Get an example
//	@Description	Returns an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [get]
func handleGetExample(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "GET example"})
}
