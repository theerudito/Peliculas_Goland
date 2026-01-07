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

func GetContent(c *fiber.Ctx) error {

	var contents []models.ContentDTO

	conn := db.GetDB()

	rows, err := conn.Query(`
	SELECT
		c.content_id,
		c.content_title,
		c.content_year,
		g.gender_id,
		g.gender_name,
		COALESCE(s.url, '') AS cover,
		CASE
		WHEN c.content_type = 1 THEN 'ANIME'
		ELSE 'SERIE'
		END AS type
	FROM content_type c
		LEFT JOIN gender AS g ON g.gender_id = c.gender_id
		LEFT JOIN storage AS s ON s.storage_id = c.cover_id`)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	for rows.Next() {

		var content models.ContentDTO

		err := rows.Scan(
			&content.Content_Id,
			&content.Content_Title,
			&content.Content_Year,
			&content.Content_Gender_Id,
			&content.Content_Gender,
			&content.Content_Cover,
			&content.Content_Type)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		contents = append(contents, content)
	}

	if len(contents) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(contents)

}

func GetContentId(c *fiber.Ctx) error {

	var content models.ContentDTO

	conn := db.GetDB()

	id := c.Params("id")

	rows, err := conn.Query(`
	SELECT
		c.content_id,
		c.content_title,
		c.content_year,
		g.gender_id,
		g.gender_name,
		COALESCE(s.url, '') AS cover,
		CASE
		WHEN c.content_type = 1 THEN 'ANIME'
		ELSE 'SERIE'
		END AS type
	FROM content_type c
		LEFT JOIN gender AS g ON g.gender_id = c.gender_id
		LEFT JOIN storage AS s ON s.storage_id = c.cover_id
	WHERE c.content_id = $1`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	found := false

	for rows.Next() {
		err := rows.Scan(
			&content.Content_Id,
			&content.Content_Title,
			&content.Content_Year,
			&content.Content_Gender_Id,
			&content.Content_Gender,
			&content.Content_Cover,
			&content.Content_Type,
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

	return c.JSON(content)

}

func GetFindContent(c *fiber.Ctx) error {

	value := helpers.QuitarGuiones(c.Params("value"))

	search := "%" + strings.ToUpper(value) + "%"

	id := c.Params("id")

	var contents []models.ContentDTO

	conn := db.GetDB()

	rows, err := conn.Query(`
	SELECT
		c.content_id,
		c.content_title,
		c.content_year,
		g.gender_id,
		g.gender_name,
		COALESCE(s.url, '') AS cover,
		CASE
		WHEN c.content_type = 1 THEN 'ANIME'
		ELSE 'SERIE'
		END AS type
	FROM content_type c
		LEFT JOIN gender AS g ON g.gender_id = c.gender_id
		LEFT JOIN storage AS s ON s.storage_id = c.cover_id
	WHERE c.content_title LIKE $1 AND c.content_type = $2
	`, search, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var content models.ContentDTO

		err := rows.Scan(
			&content.Content_Id,
			&content.Content_Title,
			&content.Content_Year,
			&content.Content_Gender_Id,
			&content.Content_Gender,
			&content.Content_Cover,
			&content.Content_Type,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		contents = append(contents, content)
	}

	if len(contents) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(contents)
}

func GetContentType(c *fiber.Ctx) error {

	id := c.Params("id")

	var dto []models.ContentDTO

	conn := db.GetDB()

	rows, err := conn.Query(`
	SELECT
		c.content_id,
		c.content_title,
		c.content_year,
		g.gender_name,
		COALESCE(s.url, '') AS cover,
		CASE
		WHEN c.content_type = 1 THEN 'ANIME'
		ELSE 'SERIE'
		END AS type
	FROM content_type c
		LEFT JOIN gender AS g ON g.gender_id = c.gender_id
		LEFT JOIN storage AS s ON s.storage_id = c.cover_id
	WHERE c.content_type = $1
	GROUP BY 
	    c.content_id, 
	    c.content_title, 
	    c.content_year,
	    g.gender_name,
	    cover,
	    type
	`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}

	defer rows.Close()

	for rows.Next() {
		var content models.ContentDTO
		err := rows.Scan(
			&content.Content_Id,
			&content.Content_Title,
			&content.Content_Year,
			&content.Content_Gender,
			&content.Content_Cover,
			&content.Content_Type,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al leer los registros",
			})
		}

		dto = append(dto, content)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	return c.JSON(dto)
}

func GetFullContent(c *fiber.Ctx) error {

	id := c.Params("id")

	var content models.ContentDTO
	seasonsMap := make(map[uint]*models.SeasonDTO)

	conn := db.GetDB()

	rows, err := conn.Query(`
		SELECT
			c.content_id,
			c.content_title,
			COALESCE(i.url, '') AS cover,
			c.content_year,
			CASE
				WHEN c.content_type = 1 THEN 'ANIME'
				ELSE 'SERIE'
			END AS type,
			e.episode_id,
			e.episode_number,
			e.episode_name,
			COALESCE(v.url, '') AS video,
			s.season_id,
			s.season_name,
			g.gender_id,
			g.gender_name
		FROM episode AS e
			LEFT JOIN season AS s ON s.season_id = e.season_id
			LEFT JOIN content_type AS c ON c.content_id = e.content_id
			LEFT JOIN gender g ON c.gender_id = g.gender_id
			LEFT JOIN storage AS i ON i.storage_id = e.video_id
			LEFT JOIN storage AS v ON v.storage_id = c.cover_id
		WHERE c.content_id = $1
		ORDER BY s.season_id, e.episode_number
	`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}
	defer rows.Close()

	found := false

	for rows.Next() {
		var (
			contentID     uint
			contentTitle  string
			contentCover  string
			contentYear   int
			contentType   string
			episodeID     uint
			episodeNumber int
			episodeName   string
			episodeURL    string
			seasonID      uint
			seasonName    string
			genderID      uint
			genderName    string
		)

		err := rows.Scan(
			&contentID,
			&contentTitle,
			&contentCover,
			&contentYear,
			&contentType,
			&episodeID,
			&episodeNumber,
			&episodeName,
			&episodeURL,
			&seasonID,
			&seasonName,
			&genderID,
			&genderName,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error leyendo registros",
			})
		}

		if content.Content_Id == 0 {
			content = models.ContentDTO{
				Content_Id:        contentID,
				Content_Title:     contentTitle,
				Content_Cover:     contentCover,
				Content_Year:      contentYear,
				Content_Type:      contentType,
				Content_Gender_Id: genderID,
				Content_Gender:    genderName,
			}
		}

		season, exists := seasonsMap[seasonID]
		if !exists {
			season = &models.SeasonDTO{
				Season_Id:   seasonID,
				Season_Name: seasonName,
			}
			seasonsMap[seasonID] = season
		}

		season.Episodes = append(season.Episodes, models.EpisodieDTO{
			Episode_Id:     episodeID,
			Episode_Number: episodeNumber,
			Episode_Name:   episodeName,
			Episode_Video:  episodeURL,
		})

		found = true
	}

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	var seasons []models.SeasonDTO
	for _, season := range seasonsMap {
		seasons = append(seasons, *season)
	}

	contentData := models.ContentDataDTO{
		Content: content,
		Seasons: seasons,
	}

	return c.JSON(contentData)

}

func PostContent(c *fiber.Ctx) error {

	var (
		content    models.Content
		existingId int
	)

	conn := db.GetDB()

	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	err := conn.QueryRow(`SELECT content_id FROM content_type WHERE UPPER(content_title) = $1`, strings.ToUpper(content.Content_Title)).Scan(&existingId)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar existencia del contenido",
		})
	}

	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Ya existe un contenido con ese título",
		})
	}

	_, err = conn.Exec(`
		INSERT INTO content_type (content_title, content_type, content_cover, content_year, gender_id)
		VALUES ($1, $2, $3, $4, $5)
	`,
		strings.ToUpper(content.Content_Title),
		content.Content_Type,
		content.Content_Cover,
		content.Content_Year,
		content.Content_Gender_Id,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al insertar el contenido",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Contenido creado correctamente 🚀",
	})

}

func PutContent(c *fiber.Ctx) error {

	id := c.Params("id")

	var content models.Content

	conn := db.GetDB()

	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	var existingId int
	err := conn.QueryRow(`
		SELECT content_id 
		FROM content_type 
		WHERE UPPER(content_title) = $1
		AND content_id != $2
	`,
		strings.ToUpper(content.Content_Title),
		id,
	).Scan(&existingId)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar existencia del contenido",
		})
	}

	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Ya existe un contenido con ese título",
		})
	}

	_, err = conn.Exec(`
		UPDATE content_type SET
			content_title 	= $1,
			content_type 	= $2,
			content_cover 	= $3,
			content_year 	= $4,
			gender_id 		= $5
		WHERE content_id 	= $6
	`,
		strings.ToUpper(content.Content_Title),
		content.Content_Type,
		content.Content_Cover,
		content.Content_Year,
		content.Content_Gender_Id,
		id,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el contenido",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Contenido actualizado correctamente 🚀",
	})
}

func PostContentSeason(c *fiber.Ctx) error {

	var data models.ContentSeason

	conn := db.GetDB()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	tx, err := conn.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al iniciar la transacción",
		})
	}

	for i, episode := range data.Episodes {

		episodeNumber := i + 1

		_, err := tx.Exec(`
			INSERT INTO episode (episode_number, episode_name, episode_url, season_id, content_id)
			VALUES ($1, $2, $3, $4, $5)
		`,
			episodeNumber,
			strings.ToUpper(episode.Episode_Name),
			episode.Episode_Video,
			data.Season_Id,
			data.Content_Id,
		)

		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al insertar episodio",
			})
		}
	}

	if err := tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al confirmar la transacción",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Episodios creados correctamente 🚀",
	})
}

func PutContentSeason(c *fiber.Ctx) error {

	var data models.ContentSeason

	conn := db.GetDB()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	tx, err := conn.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al iniciar la transacción",
		})
	}

	_, err = tx.Exec(`
		DELETE FROM episode
		WHERE season_id = $1 AND content_id = $2
	`, data.Season_Id, data.Content_Id)

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar episodios existentes",
		})
	}

	for _, episode := range data.Episodes {
		_, err := tx.Exec(`
			INSERT INTO episode (episode_number, episode_name, episode_url, season_id, content_id)
			VALUES ($1, $2, $3, $5, $6)
		`,
			episode.Episode_Number,
			strings.ToUpper(episode.Episode_Name),
			episode.Episode_Video,
			data.Season_Id,
			data.Content_Id,
		)

		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al insertar nuevo episodio",
			})
		}

	}

	if err := tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al confirmar la transacción",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Episodios actualizados correctamente 🚀",
	})

}

func DeleteContent(c *fiber.Ctx) error {

	id := c.Params("id")

	conn := db.GetDB()

	tx, err := conn.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al iniciar la transacción",
		})
	}

	_, err = tx.Exec(`
		DELETE FROM episode
		WHERE content_id = $1
	`, id)

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar episodios asociados",
		})
	}

	result, err := tx.Exec(`
		DELETE FROM content_type
		WHERE content_id = $1
	`, id)

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar el contenido",
		})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al confirmar eliminación",
		})
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontró el contenido para eliminar",
		})
	}

	if err := tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al confirmar la transacción",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Contenido y episodios eliminados correctamente 🚀",
	})

}
