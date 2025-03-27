package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetExample godoc
//
//	@Summary		Get an example
//	@Description	Returns an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [get]
func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET example"})
}
