package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const PORT = ":8080"

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{Format: "${status} - ${method} ${path}\n"}))
	log.Fatal(app.Listen(PORT))
}
