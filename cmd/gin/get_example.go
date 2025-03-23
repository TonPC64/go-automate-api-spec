package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET example"})
}
