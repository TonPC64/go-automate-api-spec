package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/example", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "GET example"})
	})

	app.Post("/example", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "POST example"})
	})

	app.Put("/example", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "PUT example"})
	})

	app.Delete("/example", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "DELETE example"})
	})
}
