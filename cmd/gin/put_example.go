package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PUT example"})
}
