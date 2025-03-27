package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// deleteExampleHandler godoc
//
//	@Summary		Delete an example
//	@Description	Deletes an example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [delete]
func deleteExampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "DELETE example"})
}
