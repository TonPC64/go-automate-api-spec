package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/example", getExampleHandler)
	app.Post("/example", postExampleHandler)
	app.Put("/example", putExampleHandler)
	app.Delete("/example", deleteExampleHandler)
}
