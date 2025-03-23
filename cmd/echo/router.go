package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/example", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "GET example"})
	})

	e.POST("/example", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, map[string]string{"message": "POST example"})
	})

	e.PUT("/example", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "PUT example"})
	})

	e.DELETE("/example", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "DELETE example"})
	})
}
