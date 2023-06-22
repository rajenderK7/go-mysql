package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("This project uses Go, GORM & MySQL.")
	})
	api := app.Group("/api")
	userApi := api.Group("/user")
	userApi.Post("/create", controllers.CreateUser)
	userApi.Get("/users", controllers.GetUsers)
	userApi.Get("/user/{id}", controllers.GetUserById)

	movieApi := api.Group("/movie")
	movieApi.Post("/create", controllers.CreateMovie)
	movieApi.Get("/movies", controllers.GetMovies)
}
