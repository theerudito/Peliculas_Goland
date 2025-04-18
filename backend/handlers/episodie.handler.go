package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func Get_Episodies(c *fiber.Ctx) error {

	var dto []models.Episodie

	pTitle := c.Params("title")

	pID := c.Params("id")

	rows, err := db.DB.Query(`SELECT 
	e.episode_id,
	e.episode_number,
	e.title,
	e.url
	FROM episodes e
	LEFT JOIN seasons s ON e.season_id  = s.season_id
	LEFT JOIN content_types ct ON s.content_type_id  = ct.content_type_id
	WHERE ct.title = ? AND ct.content_type_id  = ?`, pTitle, pID)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

		var episodie models.Episodie

		err := rows.Scan(&episodie.Episode_ID, &episodie.Episode_Number, &episodie.Episode_Name, &episodie.Episode_Url)

		if err != nil {
			return err
		}

		dto = append(dto, episodie)

	}

	return c.JSON(dto)

}
