package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// postExampleHandler godoc
//
//	@Summary		Create an example
//	@Description	Creates a new example resource
//	@Tags			examples
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/examples [post]
func postExampleHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{"message": "POST example"})
}
