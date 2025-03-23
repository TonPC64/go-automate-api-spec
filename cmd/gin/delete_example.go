package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DELETE example"})
}
