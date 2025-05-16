package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type ErrorResponse struct {
	Code int `json:"code"`
}

// getExampleHandler godoc
//
//	@Summary				Get an example
//	@Description.markdown	example_get.md
//	@Description			Returns an example resource
//	@Tags					internal
//	@Accept					json
//	@Produce				json
//	@Success				200	{object}	Response
//	@Failure				500	{object}	ErrorResponse
//	@Router					/examples [get]
func getExampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "GET example"})
}
