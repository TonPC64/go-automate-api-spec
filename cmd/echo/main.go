package main

import (
	"github.com/labstack/echo/v4"
)

//go:generate swag i -o ./../../apispec/echo --outputTypes yaml --markdownFiles ./docs-markdown

//	@title			Swagger Example API echo
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8082
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	e := echo.New()
	setupRoutes(e)

	e.Start(":8082")
}
