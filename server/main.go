package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const PORT = ":8080"
const SQLITE_DB = "recycletopia.db"

func main() {
	_, err := gorm.Open(sqlite.Open(SQLITE_DB), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't intialize GORM: ", err)
	}

	app := fiber.New()
	app.Get("/api/metrics", monitor.New())
	app.Static("/", "../client/dist")
	app.Static("*", "../client/dist/index.html")

	app.Use(logger.New(logger.Config{Format: "${status} - ${method} ${path}\n"}))
	log.Fatal(app.Listen(PORT))
}
