package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/routes"
)

func main() {
	app := fiber.New()

	db.Connect()

	routes.SetupRoutes(app)

	_ = app.Listen(":1000")
}
