package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/peliculas/handlers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/movies", handlers.Get_Movies)
	v1.Get("/movies/:id", handlers.Get_Movies)
	v1.Post("/movies", handlers.Post_Movie)
	v1.Put("/movies/:id", handlers.Put_Movies)
	v1.Delete("/movies/:id", handlers.Delete_Movies)
}
