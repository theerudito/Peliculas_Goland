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
	v1.Get("/movies", handlers.GetMovies)
	v1.Get("/movies/:id", handlers.GetMoviebyid)
	v1.Get("/movies/find/:value", handlers.FindMovie)
	v1.Post("/movies", handlers.PostMovie)
	v1.Put("/movies/:id", handlers.PutMovies)
	v1.Delete("/movies/:id", handlers.DeleteMovies)

	// CONTENT
	v1.Get("/contents", handlers.GetContent)
	v1.Get("/contents/:id", handlers.GetContentID)
	v1.Get("/contents/find/:value", handlers.FindContent)
	v1.Get("/contents/content-type/:id", handlers.GetContenType)
	v1.Get("/contents/season/:value", handlers.GetContentSeason)
	v1.Get("/contents/episodes/:id", handlers.GetContentEpisode)
	v1.Get("/contents/full/:value", handlers.GetContentFull)
	v1.Post("/contents", handlers.PostContent)
	v1.Put("/contents", handlers.PutContent)
	v1.Delete("/contents/:id", handlers.DeleteContent)

	// SEASONS
	v1.Get("/seasons", handlers.Get_Seasons)

	// GENDER
	v1.Get("/genders", handlers.Get_Gender)

	// EPISODIES
	v1.Get("/episodes", handlers.Get_Episodies)

}
