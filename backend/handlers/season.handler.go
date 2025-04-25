package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GET_Season(c *fiber.Ctx) error {

	var dto []models.SeasonDTO

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

		var season models.SeasonDTO

		err := rows.Scan(&season.Season_Id,
			&season.Season_Name)

		if err != nil {
			return err
		}

		dto = append(dto, season)

	}

	return c.JSON(dto)

}

func GET_Season_ID(c *fiber.Ctx) error {

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

func POST_Season(c *fiber.Ctx) error {
	return nil
}

func PUT_Season(c *fiber.Ctx) error {
	return nil
}

func DELETE_Season(c *fiber.Ctx) error {
	return nil
}
