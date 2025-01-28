package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AddLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{Format: "${status} - ${method} ${path}\n"}))
}
