package routes

import (
	"github.com/gofiber/fiber/v2"
)

// TODO: Complete post creationg of controllers
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/", nil)
}
