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
	v1.Get("/movie", handlers.GET_Movie)
	v1.Get("/movie/:id", handlers.GET_Movie_ID)
	v1.Get("/movie/find/:value", handlers.GET_Find_Movie)
	v1.Post("/movie", handlers.POST_Movie)
	v1.Put("/movie/:id", handlers.PUT_Movie)
	v1.Delete("/movie/:id", handlers.DELETE_Movie)

	// CONTENT
	v1.Get("/content", handlers.GET_Content)
	v1.Get("/content/:id", handlers.GET_Content_ID)
	v1.Get("/content/find/:value", handlers.GET_Find_Content)
	v1.Get("/content/type_content/:id", handlers.GET_Content_Type)
	v1.Get("/content/full_content/:id", handlers.GET_Full_Content)
	v1.Post("/content", handlers.POST_Content)
	v1.Put("/content/:id", handlers.PUT_Content)
	v1.Delete("/content/:id", handlers.DELETE_Content)

	// SEASONS
	v1.Get("/season", handlers.GET_Season)
	v1.Get("/season/:id", handlers.GET_Season_ID)
	v1.Post("/season", handlers.POST_Season)
	v1.Put("/season/:id", handlers.PUT_Season)
	v1.Delete("/season/:id", handlers.DELETE_Season)

	// GENDER
	v1.Get("/gender", handlers.GET_Gender)
	v1.Get("/gender/:id", handlers.GET_Gender_ID)
	v1.Post("/gender", handlers.POST_Gender)
	v1.Put("/gender/:id", handlers.PUT_Gender)
	v1.Delete("/gender/:id", handlers.DELETE_Gender)

	// EPISODES
	v1.Get("/episode", handlers.GET_Episode)
	v1.Get("/episode/:id", handlers.GET_Episode_ID)
	v1.Post("/episode", handlers.POST_Episode)
	v1.Put("/episode/:id", handlers.PUT_Episode)
	v1.Delete("/episode/id", handlers.DELETE_Episode)

	// LOGIN
	v1.Post("/login", handlers.GET_Episode)

}
