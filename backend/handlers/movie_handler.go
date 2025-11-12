package handlers

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/models"
)

func GetMovie(c *fiber.Ctx) error {

	var dto []models.MovieDTO

	conn := db.GetDB()

	rows, err := conn.Query(`
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

func GetMovieId(c *fiber.Ctx) error {

	id := c.Params("id")

	var movie models.MovieDTO

	conn := db.GetDB()

	rows, err := conn.Query(`
	SELECT
		m.movie_id,
		m.movie_title,
		m.movie_year,
		m.movie_cover,
		m.movie_url,
		g.gender_name
	FROM movie AS m
	INNER JOIN gender AS g ON m.gender_id = g.gender_id
	WHERE m.movie_id = $1
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

func GetFindMovie(c *fiber.Ctx) error {

	value := helpers.QuitarGuiones(c.Params("value"))

	var dto []models.MovieDTO

	conn := db.GetDB()

	search := "%" + strings.ToUpper(value) + "%"

	rows, err := conn.Query(`
	SELECT
		m.movie_id,
		m.movie_title,
		m.movie_year,
		m.movie_cover,
		m.movie_url,
		g.gender_name
	FROM movie AS m
	INNER JOIN gender AS g ON m.gender_id = g.gender_id
	WHERE m.movie_title LIKE $1
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

func PostMovie(c *fiber.Ctx) error {

	var movie models.Movie

	conn := db.GetDB()

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := conn.QueryRow("SELECT movie_id FROM movie WHERE movie.movie_title = $1 ", movie.Movie_Title)

	var existingId int

	if err := row.Scan(&existingId); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar la existencia del registro",
		})
	}

	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "El registro ya existe",
		})
	}

	_, err := conn.Exec(`INSERT INTO movie (movie_title, movie_year, movie_cover, movie_url, gender_id) VALUES ($1, $2, $3, $4, $5)`,
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

func PutMovie(c *fiber.Ctx) error {

	id := c.Params("id")

	conn := db.GetDB()

	var movie models.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	row := conn.QueryRow("SELECT movie_id FROM movie WHERE movie_id = $1", id)

	var existingId int
	if err := row.Scan(&existingId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar la existencia del registro",
		})
	}

	_, err := conn.Exec(`UPDATE movie SET
		movie_title 	= $1, 
		movie_year 		= $2, 
		movie_cover 	= $3, 
		movie_url 		= $4, 
		gender_id 		= $5
		WHERE movie_id 	= $6`,
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

func DeleteMovie(c *fiber.Ctx) error {

	id := c.Params("id")

	conn := db.GetDB()

	row := conn.QueryRow("SELECT movie_id FROM movie WHERE movie.movie_id = $1", id)

	var existingId int

	if err := row.Scan(&existingId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No se encontró el registro",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar la existencia del registro",
		})
	}

	_, err := conn.Exec(`DELETE FROM movie WHERE movie.movie_id = $1`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo eliminar el registro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registro eliminado correctamente 🚀",
	})

}
