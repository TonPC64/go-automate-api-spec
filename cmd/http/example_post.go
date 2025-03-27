package main

import (
	"encoding/json"
	"net/http"
)

// handlePostExample godoc
//
//	@Summary		Create an example
//	@Description	Creates a new example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [post]
func handlePostExample(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "POST example"})
}
