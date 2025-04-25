package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GET_Gender(c *fiber.Ctx) error {

	var dto []models.Gender

	rows, err := db.DB.Query(`
		SELECT g.gender_id, g.gender_name
		FROM gender AS g
	`)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al consultar los registros",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var gender models.Gender
		err := rows.Scan(&gender.Gender_Id, &gender.Gender_Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}
		dto = append(dto, gender)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(dto)

}

func GET_Gender_ID(c *fiber.Ctx) error {

	id := c.Params("id")

	var gender models.Gender

	row := db.DB.QueryRow(`
		SELECT g.gender_id, g.gender_name
		FROM gender AS g
		WHERE g.gender_id = ?`, id)

	err := row.Scan(&gender.Gender_Id, &gender.Gender_Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al consultar el registro",
		})
	}

	return c.JSON(gender)

}

func POST_Gender(c *fiber.Ctx) error {

	var gender models.Gender

	if err := c.BodyParser(&gender); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT gender_id FROM gender WHERE gender_name = ?", gender.Gender_Name)

	var existingId int
	if err := row.Scan(&existingId); err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar el registro",
		})
	}
	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "El registro ya existe",
		})
	}

	_, err := db.DB.Exec("INSERT INTO gender (gender_name) VALUES (?)", gender.Gender_Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear el registro",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registro creado correctamente 🚀",
	})

}

func PUT_Gender(c *fiber.Ctx) error {

	id := c.Params("id")

	var gender models.Gender

	if err := c.BodyParser(&gender); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT gender_id FROM gender WHERE gender_id = ?", id)

	var existingId int
	if err := row.Scan(&existingId); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar el registro",
		})
	}

	_, err := db.DB.Exec("UPDATE gender SET gender_name = ? WHERE gender_id = ?", gender.Gender_Name, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro actualizado correctamente 🚀",
	})

}

func DELETE_Gender(c *fiber.Ctx) error {

	id := c.Params("id")

	row := db.DB.QueryRow("SELECT gender_id FROM gender WHERE gender_id = ?", id)

	var existingId int

	if err := row.Scan(&existingId); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar el registro",
		})
	}

	_, err := db.DB.Exec("DELETE FROM gender WHERE gender_id = ?", id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo eliminar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro eliminado correctamente 🚀",
	})

}
