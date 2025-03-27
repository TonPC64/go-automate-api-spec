package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostExample godoc
//
//	@Summary		Create an example
//	@Description	Creates a new example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [post]
func PostExample(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "POST example"})
}
