package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// getExampleHandler godoc
//
//	@Summary		Get an example
//	@Description	Returns an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [get]
func getExampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "GET example"})
}
