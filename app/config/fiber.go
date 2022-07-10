package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"hadithgo/app/exceptions"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
		BodyLimit:    100 * 1024 * 1024,
	}
}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}
}

func LoggerConfig() logger.Config {
	return logger.Config{
		Format:   "[${time}] - ${method} ${path} ~ [${status}][${latency}]\n",
		TimeZone: "Local",
	}
}
