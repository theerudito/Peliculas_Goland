package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func Get_Contents(c *fiber.Ctx) error {

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`SELECT
	content.content_id,
	content.content_cover,
	content.content_url,
	content.content_year,
	gender.gender_name,
	CASE
			WHEN content.content_type = 1 THEN 'ANIME'
			ELSE 'SERIE'
	END AS type
	FROM content_types content
	INNER JOIN genders AS gender ON gender.gender_id = content.gender_id`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.ContentDTO
		err := rows.Scan(&content.Content_Id, &content.Content_Cover, &content.Content_Url, &content.Content_Year, &content.Gender, &content.Content_Type)
		if err != nil {
			return err
		}
		dto = append(dto, content)
	}
	return c.JSON(dto)

}

func Get_Content_ID(c *fiber.Ctx) error {

	var dto []models.ContentDTO

	pID := c.Params("id")

	rows, err := db.DB.Query(`SELECT
	content.content_id,
	content.content_cover,
	content.content_url,
	content.content_year,
	gender.gender_name,
	CASE
			WHEN content.content_type = 1 THEN 'ANIME'
			ELSE 'SERIE'
	END AS type
	FROM content_types content
	INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
	WHERE content.content_type = 1`, pID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.ContentDTO
		if err := rows.Scan(&content.Content_Id, &content.Content_Cover, &content.Content_Url, &content.Content_Year, &content.Gender, &content.Content_Type); err != nil {
			return err
		}
		dto = append(dto, content)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No content found for the given category.",
		})
	} else {
		return c.JSON(dto)
	}

}
