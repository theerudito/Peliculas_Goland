package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GetMovies(c *fiber.Ctx) error {
	var movieDTOs []models.MovieDTO

	query, err := db.DB.Query(`
		SELECT 
			m.id, m.title, m.year, g.descripcion 
		FROM 
			movies m 
		JOIN 
			genders g 
		ON 
			m.gender_id = g.id
	`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching movies",
		})
	}
	defer query.Close()

	for query.Next() {
		var dto models.MovieDTO
		if err := query.Scan(&dto.ID, &dto.Title, &dto.Year, &dto.Gender); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error scanning movie",
			})
		}
		movieDTOs = append(movieDTOs, dto)
	}

	return c.JSON(movieDTOs)

}
