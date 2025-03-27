package main

import (
	"encoding/json"
	"net/http"
)

// handleDeleteExample godoc
//
//	@Summary		Delete an example
//	@Description	Deletes an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [delete]
func handleDeleteExample(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "DELETE example"})
}
