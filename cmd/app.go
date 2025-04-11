package cmd

import (
	"fmt"

	"github.com/fiqrikm18/ISITransaction/internal/configs"
	"github.com/fiqrikm18/ISITransaction/internal/routes"
	"github.com/gofiber/fiber/v2"
)

// @title ISI Saving Transation API
// @version 1.0
// @description This is a simpel saving transaction API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiqrikm18@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func Run() {
	initServer()
}

func initServer() {
	appPort := configs.AppPort

	// initialize the fiber app
	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: false,
		CaseSensitive: false,
	})

	// registering application routers
	routes.RegisterApiRouter(app)
	routes.RegisterWebRouter(app)

	// start the server
	fmt.Printf("Server is running on port %s\n", appPort)
	app.Listen(appPort)
}
