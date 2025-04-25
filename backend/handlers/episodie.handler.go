package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GET_Episode(c *fiber.Ctx) error {

	var dto []models.EpisodieDTO

	rows, err := db.DB.Query(`SELECT
	episode.episode_id,
	episode.episode_name,
	episode.episode_number,
	episode.episode_url,
	FROM episodes AS episode
	INNER JOIN seasons AS season ON season.season_id = episode.season_id`)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		var episodie models.EpisodieDTO

		err := rows.Scan(
			&episodie.Episode_Id,
			&episodie.Episode_Name,
			&episodie.Episode_Number,
			&episodie.Episode_Url,
		)

		if err != nil {
			return err
		}

		dto = append(dto, episodie)

	}

	return c.JSON(dto)

}

func GET_Episode_ID(c *fiber.Ctx) error {
	return nil
}

func POST_Episode(c *fiber.Ctx) error {
	return nil
}

func PUT_Episode(c *fiber.Ctx) error {
	return nil
}

func DELETE_Episode(c *fiber.Ctx) error {
	return nil
}
