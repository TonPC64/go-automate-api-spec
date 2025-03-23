package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostExample(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "POST example"})
}
