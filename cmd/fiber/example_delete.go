package main

import (
	"github.com/gofiber/fiber/v2"
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
func deleteExampleHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "DELETE example"})
}
