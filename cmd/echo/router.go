package main

import (
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/example", getExampleHandler)
	e.POST("/example", postExampleHandler)
	e.PUT("/example", putExampleHandler)
	e.DELETE("/example", deleteExampleHandler)
}
