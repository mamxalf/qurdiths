package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"hadithgo/app/config"
	"hadithgo/app/exceptions"
	"hadithgo/app/routers"
)

func main() {
	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(logger.New(config.LoggerConfig()))

	// Fiber Monitor
	app.Get("/monitor", monitor.New())

	// Setup Routing API
	api := app.Group("/api")

	// Setup V1 Router
	v1 := api.Group("/v1")
	routers.SetupRoutesV1(v1)

	// Start App
	err := app.Listen(":3000")
	exceptions.PanicIfNeeded(err)
}
