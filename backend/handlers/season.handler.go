package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func Get_Seasons(c *fiber.Ctx) error {

	var dto []models.Season

	rows, err := db.DB.Query(`
	SELECT
	season.season_id,
	season.season_name
	FROM seasons AS season
	`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var season models.Season

		err := rows.Scan(&season.Season_Id,
			&season.Season_Name)

		if err != nil {
			return err
		}

		dto = append(dto, season)

	}

	return c.JSON(dto)

}

func Get_Season_ID(c *fiber.Ctx) error {

	var dto []models.Season

	pID := c.Params("id")

	rows, err := db.DB.Query(`SELECT
	season.season_id,
	season.season_title
	FROM seasons AS season
	WHERE season.season_id = ?`, pID)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var season models.Season

		err := rows.Scan(&season.Season_Id,
			&season.Season_Name)

		if err != nil {
			return err
		}

		dto = append(dto, season)

	}

	return c.JSON(dto)

}
