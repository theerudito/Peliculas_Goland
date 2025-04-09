package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/peliculas/models"
)

func GetMovies(c *fiber.Ctx) error {
	var movies []models.Movie
	//var movieDTOs []model.MovieDTO

	return c.JSON(movies)

}
