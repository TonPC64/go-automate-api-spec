package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PutExample godoc
//
//	@Summary		Update an example
//	@Description	Updates an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [put]
func PutExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PUT example"})
}
