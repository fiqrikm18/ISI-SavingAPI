package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/fiqrikm18/ISITransaction/internal/configs"
	"github.com/fiqrikm18/ISITransaction/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
	slogmulti "github.com/samber/slog-multi"
)

// @title ISI Saving Transaction API
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
	// Create log file
	file, err := os.OpenFile("logs/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// Set up two handlers: terminal and file
	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo})

	// Combine both using a multi-logger (samber/slog-multi)
	loggerHanlers := slog.New(slogmulti.Fanout(fileHandler))
	slog.SetDefault(loggerHanlers)

	appConfig := configs.NewConfig()

	// initialize the fiber app
	app := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: false,
		CaseSensitive: false,
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(slogfiber.New(slog.Default()))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	// registering application routers
	routes.RegisterApiRouter(app)
	routes.RegisterWebRouter(app)

	// start the server
	if err := app.Listen(appConfig.AppPort); err != nil {
		fmt.Println("Error starting server:", err.Error())
		return
	}
}
