package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GET_Episode(c *fiber.Ctx) error {

	var dto []models.Episodie

	rows, err := db.DB.Query(`
		SELECT
			e.episode_id,
			e.episode_name,
			e.episode_number,
			e.episode_url
		FROM episode AS e
	`)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al consultar los registros",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var episodie models.Episodie

		err := rows.Scan(
			&episodie.Episode_Id,
			&episodie.Episode_Name,
			&episodie.Episode_Number,
			&episodie.Episode_Url,
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

func GET_Episode_ID(c *fiber.Ctx) error {

	id := c.Params("id")

	var episodie models.Episodie

	row := db.DB.QueryRow(`
		SELECT 
			e.episode_id, 
			e.episode_name, 
			e.episode_number, 
			e.episode_url
		FROM episode AS e
		WHERE e.episode_id = ?`, id)

	err := row.Scan(
		&episodie.Episode_Id,
		&episodie.Episode_Name,
		&episodie.Episode_Number,
		&episodie.Episode_Url,
	)

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

	return c.JSON(episodie)

}

func POST_Episode(c *fiber.Ctx) error {

	var episodie models.Episodie

	if err := c.BodyParser(&episodie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT episode_id FROM episode WHERE episode_name = ?", episodie.Episode_Name)

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

	_, err := db.DB.Exec(
		"INSERT INTO episode (episode_name, episode_number, episode_url) VALUES (?, ?, ?)",
		episodie.Episode_Name,
		episodie.Episode_Number,
		episodie.Episode_Url,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo crear el registro",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registro creado correctamente 🚀",
	})

}

func PUT_Episode(c *fiber.Ctx) error {

	id := c.Params("id")

	var episodie models.EpisodieDTO

	if err := c.BodyParser(&episodie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT episode_id FROM episode WHERE episode_id = ?", id)

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

	_, err := db.DB.Exec(
		"UPDATE episode SET episode_name = ?, episode_number = ?, episode_url = ? WHERE episode_id = ?",
		episodie.Episode_Name,
		episodie.Episode_Number,
		episodie.Episode_Url,
		id,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro actualizado correctamente 🚀",
	})

}

func DELETE_Episode(c *fiber.Ctx) error {

	id := c.Params("id")

	row := db.DB.QueryRow("SELECT episode_id FROM episode WHERE episode_id = ?", id)

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

	_, err := db.DB.Exec("DELETE FROM episode WHERE episode_id = ?", id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo eliminar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro eliminado correctamente 🚀",
	})

}
