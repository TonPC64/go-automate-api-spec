package main

import (
	"github.com/gofiber/fiber/v2"
)

//go:generate swag i -o ./../../apispec/fiber --outputTypes yaml --markdownFiles ./docs-markdown

//	@title			Swagger Example API fiber
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8083
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":8083")
}
