package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GET_Season(c *fiber.Ctx) error {

	var seasons []models.Season

	rows, err := db.DB.Query(`
		SELECT
		s.season_id,
		s.season_name
		FROM season AS s`)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var season models.Season

		err := rows.Scan(&season.Season_Id, &season.Season_Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		seasons = append(seasons, season)
	}

	if len(seasons) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(seasons)

}

func GET_Season_ID(c *fiber.Ctx) error {

	id := c.Params("id")

	var season models.Season

	rows, err := db.DB.Query(`
		SELECT
		s.season_id,
		s.season_name
		FROM season AS s
		WHERE s.season_id = ?`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	found := false

	for rows.Next() {
		err := rows.Scan(&season.Season_Id, &season.Season_Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}
		found = true
	}

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(season)
}

func POST_Season(c *fiber.Ctx) error {

	var season models.Season

	if err := c.BodyParser(&season); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT season_id FROM season WHERE season_name = ?", season.Season_Name)

	var existingId int

	if err := row.Scan(&existingId); err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar la existencia del registro",
		})
	}

	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "El registro ya existe",
		})
	}

	_, err := db.DB.Exec("INSERT INTO season (season_title) VALUES (?)", season.Season_Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al insertar el registro",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registro creado correctamente 🚀",
	})

}

func PUT_Season(c *fiber.Ctx) error {

	id := c.Params("id")

	var season models.Season
	if err := c.BodyParser(&season); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT season_id FROM season WHERE season_id = ?", id)

	var existingId int
	if err := row.Scan(&existingId); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar la existencia del registro",
		})
	}

	_, err := db.DB.Exec("UPDATE season SET season_name = ? WHERE season_id = ?", season.Season_Name, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo actualizar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro actualizado correctamente 🚀",
	})

}

func DELETE_Season(c *fiber.Ctx) error {

	id := c.Params("id")

	row := db.DB.QueryRow("SELECT season_id FROM season WHERE season_id = ?", id)
	var existingId int
	if err := row.Scan(&existingId); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar la existencia del registro",
		})
	}

	_, err := db.DB.Exec("DELETE FROM season WHERE season_id = ?", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo eliminar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro eliminado correctamente 🚀",
	})

}
