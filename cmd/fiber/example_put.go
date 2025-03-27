package main

import (
	"github.com/gofiber/fiber/v2"
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
func putExampleHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "PUT example"})
}
