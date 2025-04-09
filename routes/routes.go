package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/peliculas/handlers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/movies", handlers.GetMovies)
	// app.Post("/movies", handlers.CreateMovie)

	// app.Get("/series", handlers.GetSeries)
	// app.Post("/series", handlers.CreateSeries)
}
