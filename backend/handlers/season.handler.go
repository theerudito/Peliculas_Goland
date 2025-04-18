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
	s.season_id,
	s.title,
	ct.descripcion,
	ct."year"
	FROM seasons s
	`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var season models.Season

		err := rows.Scan(&season.Season_ID, &season.Season_Name)

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
	s.season_id,
	s.title,
	ct.descripcion,
	ct."year"
	FROM seasons s
	LEFT JOIN content_types ct ON s.content_type_id  = ct.content_type_id
	WHERE ct.title = 'Anime' AND ct.content_type_id  = ?`, pID)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var season models.Season

		err := rows.Scan(&season.Season_ID, &season.Season_Name)

		if err != nil {
			return err
		}

		dto = append(dto, season)

	}

	return c.JSON(dto)

}
