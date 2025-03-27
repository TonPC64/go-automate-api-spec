package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteExample godoc
//
//	@Summary		Delete an example
//	@Description	Deletes an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [delete]
func DeleteExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DELETE example"})
}
