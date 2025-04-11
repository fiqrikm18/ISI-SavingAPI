package routes

import (
	"fmt"

	"github.com/fiqrikm18/ISITransaction/docs"
	"github.com/fiqrikm18/ISITransaction/internal/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterWebRouter(route *fiber.App) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost%s", configs.AppPort)
	docs.SwaggerInfo.BasePath = "/api/v1/"
	docs.SwaggerInfo.Title = "ISI Saving Transaction API"
	docs.SwaggerInfo.Description = "This is a simple saving transaction API."
	docs.SwaggerInfo.Version = "1.0"

	route.Get("/docs/*", swagger.New(swagger.Config{})) // default
}
