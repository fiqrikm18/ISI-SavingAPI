package routes

import (
	"github.com/fiqrikm18/ISITransaction/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

var (
	pingController controllers.IPingController
)

func init() {
	pingController = controllers.NewPingController()
}

func RegisterApiRouter(route *fiber.App) {
	api := route.Group("/api/v1")
	api.Get("/ping", pingController.Ping)
}
