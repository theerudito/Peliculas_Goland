package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/models"
)

func GET_Movie(c *fiber.Ctx) error {

	var dto []models.MovieDTO

	rows, err := db.DB.Query(`
		SELECT
		m.movie_id,
		m.movie_title,
		m.movie_year,
		m.movie_cover,
		m.movie_url,
		g.gender_name
		FROM movie AS m
		INNER JOIN gender AS g ON m.gender_id = g.gender_id
	`)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var movie models.MovieDTO

		err := rows.Scan(
			&movie.Movie_Id,
			&movie.Movie_Title,
			&movie.Movie_Year,
			&movie.Movie_Cover,
			&movie.Movie_Url,
			&movie.Gender,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		dto = append(dto, movie)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(dto)

}

func GET_Movie_ID(c *fiber.Ctx) error {

	id := c.Params("id")

	var movie models.MovieDTO

	rows, err := db.DB.Query(`
	SELECT
		m.movie_id,
		m.movie_title,
		m.movie_year,
		m.movie_cover,
		m.movie_url,
		g.gender_name
	FROM movie AS m
	INNER JOIN gender AS g ON m.gender_id = g.gender_id
	WHERE m.movie_id = ?
`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}
	defer rows.Close()

	found := false

	for rows.Next() {
		err := rows.Scan(
			&movie.Movie_Id,
			&movie.Movie_Title,
			&movie.Movie_Year,
			&movie.Movie_Cover,
			&movie.Movie_Url,
			&movie.Gender,
		)

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

	return c.JSON(movie)

}

func GET_Find_Movie(c *fiber.Ctx) error {

	value := helpers.QuitarGuiones(c.Params("value"))

	var dto []models.MovieDTO

	search := "%" + strings.ToUpper(value) + "%"

	rows, err := db.DB.Query(`
	SELECT
		m.movie_id,
		m.movie_title,
		m.movie_year,
		m.movie_cover,
		m.movie_url,
		g.gender_name
	FROM movie AS m
	INNER JOIN gender AS g ON m.gender_id = g.gender_id
	WHERE m.movie_title LIKE ?
`, search)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var movie models.MovieDTO

		err := rows.Scan(
			&movie.Movie_Id,
			&movie.Movie_Title,
			&movie.Movie_Year,
			&movie.Movie_Cover,
			&movie.Movie_Url,
			&movie.Gender,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		dto = append(dto, movie)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(dto)

}

func POST_Movie(c *fiber.Ctx) error {

	var movie models.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT movie_id FROM movie WHERE movie.movie_title = ? ", movie.Movie_Title)

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

	_, err := db.DB.Exec(`INSERT INTO movie (movie_title, movie_year, movie_cover, movie_url, gender_id) VALUES (?, ?, ?, ?, ?)`,
		strings.ToUpper(movie.Movie_Title),
		movie.Movie_Year,
		movie.Movie_Cover,
		movie.Movie_Url,
		movie.Gender_Id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al insertar el registro",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registro creado correctamente 🚀",
	})

}

func PUT_Movie(c *fiber.Ctx) error {

	id := c.Params("id")

	var movie models.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := db.DB.QueryRow("SELECT movie_id FROM movie WHERE movie_id = ?", id)

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

	_, err := db.DB.Exec(`UPDATE movie SET
		movie_title = ?, 
		movie_year = ?, 
		movie_cover = ?, 
		movie_url = ?, 
		gender_id = ?
		WHERE movie_id = ?`,
		strings.ToUpper(movie.Movie_Title),
		movie.Movie_Year,
		movie.Movie_Cover,
		movie.Movie_Url,
		movie.Gender_Id,
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

func DELETE_Movie(c *fiber.Ctx) error {

	id := c.Params("id")

	row := db.DB.QueryRow("SELECT movie_id FROM movie WHERE movie.movie_id = ?", id)

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

	_, err := db.DB.Exec(`DELETE FROM movie WHERE movie.movie_id = ?`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo eliminar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro eliminado correctamente 🚀",
	})

}
