package main

import (
	"log"

	"github.com/amodhakal/recycletopia/repo"
	"github.com/amodhakal/recycletopia/utils"
	"github.com/gofiber/fiber/v2"
)

const PORT = ":8080"

func main() {
	app := fiber.New()

	repo.InitDB()
	utils.AddLogger(app)
	utils.AddMetrics(app)
	utils.AddStaticFiles(app)

	log.Fatal(app.Listen(PORT))
}
