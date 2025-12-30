package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	app := fiber.New(fiber.Config{
		BodyLimit:    200 * 1024 * 1024,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 2 * time.Minute,
	})

	db.InitDB()
	defer db.GetDB().Close()

	if err := helpers.CreateFolders(); err != nil {
		fmt.Println("Error:", err)
	}

	routes.SetupRoutes(app)

	_ = app.Listen(fmt.Sprintf(":%s", os.Getenv("PortServer")))

}
