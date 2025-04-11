package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/peliculas/handlers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// MOVIES
	v1.Get("/movies", handlers.Get_Movies)
	v1.Get("/movies/:id", handlers.Get_Movies)
	v1.Post("/movies", handlers.Post_Movie)
	v1.Put("/movies/:id", handlers.Put_Movies)
	v1.Delete("/movies/:id", handlers.Delete_Movies)

	// CATEGORIES
	v1.Get("/categories", handlers.Get_Contents)
	v1.Get("/categories/:id", handlers.Get_Content_ID)

	// SEASONS
	v1.Get("/season", handlers.Get_Seasons)
	v1.Get("/season/:id", handlers.Get_Season_ID)

	// EPISODIES
	v1.Get("/episodies/:title/:id", handlers.Get_Episodies)

}
