package utils

import "github.com/gofiber/fiber/v2"

func AddStaticFiles(app *fiber.App) {
	app.Static("/", "../client/dist")
	app.Static("*", "../client/dist/index.html")
}
