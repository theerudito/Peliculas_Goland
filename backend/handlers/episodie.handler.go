package handlers

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GetEpisode(c *fiber.Ctx) error {

	var dto []models.EpisodieDTO

	conn := db.GetDB()

	rows, err := conn.Query(`
		SELECT
			e.episode_id,
			e.episode_name,
			e.episode_number,
			e.episode_url,
			s.season_name
			FROM episode AS e
			INNER JOIN season AS s ON s.season_id = e.season_id
	`)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al consultar los registros",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var episodie models.EpisodieDTO

		err := rows.Scan(
			&episodie.Episode_Id,
			&episodie.Episode_Name,
			&episodie.Episode_Number,
			&episodie.Episode_Url,
			&episodie.Season,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		dto = append(dto, episodie)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(dto)

}

func GetEpisodeId(c *fiber.Ctx) error {

	id := c.Params("id")

	var episodie models.EpisodieDTO

	conn := db.GetDB()

	row := conn.QueryRow(`
	SELECT
	e.episode_id,
	e.episode_name,
	e.episode_number,
	e.episode_url,
	s.season_name
	FROM episode AS e
	INNER JOIN season AS s ON s.season_id = e.season_id
	WHERE e.episode_id = $1`, id)

	err := row.Scan(
		&episodie.Episode_Id,
		&episodie.Episode_Name,
		&episodie.Episode_Number,
		&episodie.Episode_Url,
		&episodie.Season,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al consultar el registro",
		})
	}

	return c.JSON(episodie)

}
