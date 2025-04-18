package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func Get_Contents(c *fiber.Ctx) error {

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`SELECT 
	ct.content_type_id,
	ct.title,
	ct.descripcion,
	ct.cover,
	ct.year
	FROM content_types ct`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.ContentDTO
		err := rows.Scan(&content.Content_Id, &content.Content_Type, &content.Content_Cover, &content.Content_Url, &content.Content_Year)
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

	rows, err := db.DB.Query(`
		SELECT 
			ct.content_type_id,
			ct.title,
			ct.descripcion,
			ct.cover,
			ct.year
		FROM content_types ct
		WHERE ct.title = ?`, pID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var content models.ContentDTO
		if err := rows.Scan(&content.Content_Id, &content.Content_Type, &content.Content_Cover, &content.Content_Url, &content.Content_Year); err != nil {
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
