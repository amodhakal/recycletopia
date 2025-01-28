package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func AddMetrics(app *fiber.App) {
	app.Get("/api/metrics", monitor.New())
}
