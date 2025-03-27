package main

import (
	"github.com/gofiber/fiber/v2"
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
func postExampleHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "POST example"})
}
