package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/example", GetExample)
	r.POST("/example", PostExample)
	r.PUT("/example", PutExample)
	r.DELETE("/example", DeleteExample)
}
