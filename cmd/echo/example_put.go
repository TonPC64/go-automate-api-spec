package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// putExampleHandler godoc
//
//	@Summary		Update an example
//	@Description	Updates an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [put]
func putExampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "PUT example"})
}
