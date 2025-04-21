package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func Get_Gender(c *fiber.Ctx) error {

	var dto []models.Gender

	rows, err := db.DB.Query(`
	SELECT
	gender.gender_id,
	gender.gender_name
	FROM genders AS gender
	`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var gender models.Gender

		err := rows.Scan(&gender.Gender_Id,
			&gender.Gender_Name)

		if err != nil {
			return err
		}

		dto = append(dto, gender)

	}

	return c.JSON(dto)
}

func Get_Gender_ID(c *fiber.Ctx) error {

	var dto []models.Gender

	pID := c.Params("id")

	rows, err := db.DB.Query(`SELECT
	gender.gender_id,
	gender.gender_name
	FROM genders AS gender
	WHERE gender.gender_id = ?`, pID)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var gender models.Gender

		err := rows.Scan(&gender.Gender_Id,
			&gender.Gender_Name)

		if err != nil {
			return err
		}

		dto = append(dto, gender)

	}

	return c.JSON(dto)

}
