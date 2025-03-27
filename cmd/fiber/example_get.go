package main

import (
	"github.com/gofiber/fiber/v2"
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
func getExampleHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "GET example"})
}
