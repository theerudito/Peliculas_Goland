package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theerudito/peliculas/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// MOVIES
	v1.Get("/movies", handlers.Get_Movies)
	v1.Get("/movies/:id", handlers.Get_Movies)
	v1.Post("/movies", handlers.Post_Movie)
	v1.Put("/movies/:id", handlers.Put_Movies)
	v1.Delete("/movies/:id", handlers.Delete_Movies)

	// CONTENT
	v1.Get("/contents", handlers.Get_Contents)
	v1.Get("/contents/:id", handlers.Get_Content_ID)
	v1.Post("/contents", handlers.Post_Movie)
	v1.Put("/contents/:id", handlers.Put_Movies)
	v1.Delete("/contents/:id", handlers.Delete_Movies)

	// SEASONS
	v1.Get("/seasons", handlers.Get_Seasons)

	// GENDER
	v1.Get("/genders", handlers.Get_Gender)

	// EPISODIES
	v1.Get("/episodies", handlers.Get_Episodies)

}
