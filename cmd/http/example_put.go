package main

import (
	"encoding/json"
	"net/http"
)

// handlePutExample godoc
//
//	@Summary		Update an example
//	@Description	Updates an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [put]
func handlePutExample(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "PUT example"})
}
