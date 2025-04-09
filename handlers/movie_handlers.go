package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GetMovies(c *fiber.Ctx) error {
	var movies []models.Movies

	db.DB.Find(&movies)

	return c.Status(fiber.StatusOK).JSON(movies)
}

func CreateMovie(c *fiber.Ctx) {
	var movie models.Movies
	if err := c.BodyParser(&movie); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		return
	}

	if err := db.DB.Create(&movie).Error; err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create movie"})
		return
	}
	c.Status(http.StatusCreated).JSON(movie)
}
