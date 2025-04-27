package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/models"
)

func GET_Content(c *fiber.Ctx) error {

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
	SELECT
	c.content_id,
	c.content_title,
	c.content_cover,
	c.content_year,
	g.gender_name,
	CASE
	WHEN c.content_type = 1 THEN 'ANIME'
	ELSE 'SERIE'
	END AS type
	FROM content_type c
	INNER JOIN gender AS g ON g.gender_id = c.gender_id`)

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
			&content.Content_Cover,
			&content.Content_Year,
			&content.Content_Gender,
			&content.Content_Type)

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

func GET_Content_ID(c *fiber.Ctx) error {

	var dto models.ContentDTO

	id := c.Params("id")

	rows, err := db.DB.Query(`
	SELECT
	c.content_id,
	c.content_title,
	c.content_cover,
	c.content_year,
	g.gender_name,
	CASE
	WHEN c.content_type = 1 THEN 'ANIME'
	ELSE 'SERIE'
	END AS type
	FROM content_type c
	INNER JOIN gender AS g ON g.gender_id = c.gender_id
	WHERE c.content_id = ?
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
			&dto.Content_Id,
			&dto.Content_Title,
			&dto.Content_Cover,
			&dto.Content_Year,
			&dto.Content_Gender,
			&dto.Content_Type,
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

	return c.JSON(dto)

}

func GET_Find_Content(c *fiber.Ctx) error {

	value := helpers.QuitarGuiones(c.Params("value"))

	search := "%" + strings.ToUpper(value) + "%"

	id := c.Params("id")

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
	SELECT
	c.content_id,
	c.content_title,
	c.content_cover,
	c.content_year,
	g.gender_name,
	CASE
	WHEN c.content_type = 1 THEN 'ANIME'
	ELSE 'SERIE'
	END AS type
	FROM content_type c
	INNER JOIN gender AS g ON g.gender_id = c.gender_id
	WHERE c.content_title LIKE ? AND c.content_type = ?
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
			&content.Content_Cover,
			&content.Content_Year,
			&content.Content_Gender,
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

func GET_Content_Type(c *fiber.Ctx) error {

	id := (c.Params("id"))

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
	SELECT
	c.content_id,
	c.content_title,
	c.content_cover,
	c.content_year,
	g.gender_name,
	CASE
	WHEN c.content_type = 1 THEN 'ANIME'
	ELSE 'SERIE'
	END AS type
	FROM content_type c
	INNER JOIN gender AS g ON g.gender_id = c.gender_id
	WHERE c.content_type = ?
	GROUP BY c.content_title, c.content_type
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
			&content.Content_Cover,
			&content.Content_Year,
			&content.Content_Gender,
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

func GET_Full_Content(c *fiber.Ctx) error {

	id := c.Params("id")

	var content models.ContentDTO
	seasonsMap := make(map[uint]*models.SeasonDTO)

	rows, err := db.DB.Query(`
		SELECT
			c.content_id,
			c.content_title,
			c.content_cover,
			c.content_year,
			CASE
				WHEN c.content_type = 1 THEN 'ANIME'
				ELSE 'SERIE'
			END AS type,
			e.episode_id,
			e.episode_number,
			e.episode_name,
			e.episode_url,
			s.season_id,
			s.season_name,
			g.gender_name
		FROM episode AS e
		LEFT JOIN season AS s ON s.season_id = e.season_id
		LEFT JOIN content_type AS c ON c.content_id = e.content_id
		LEFT JOIN gender g ON c.gender_id = g.gender_id
		WHERE c.content_id = ?
		ORDER BY s.season_id, e.episode_number
	`, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al ejecutar la consulta",
		})
	}
	defer rows.Close()

	found := false // 👈 controlamos si se encontraron registros

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
			&genderName,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error leyendo registros",
			})
		}

		if content.Content_Id == 0 {
			content = models.ContentDTO{
				Content_Id:     contentID,
				Content_Title:  contentTitle,
				Content_Cover:  contentCover,
				Content_Year:   contentYear,
				Content_Type:   contentType,
				Content_Gender: genderName,
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
			Episode_Url:    episodeURL,
		})

		found = true // 👈 aquí marcamos que encontramos algo
	}

	// 👇 Despues de leer todo, verificamos si no encontró nada
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No se encontraron registros",
		})
	}

	// Armamos seasons
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

func POST_Content(c *fiber.Ctx) error {

	var content models.Content

	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	var existingId int
	err := db.DB.QueryRow(`
		SELECT content_id 
		FROM content_type 
		WHERE UPPER(content_title) = ?
	`,
		strings.ToUpper(content.Content_Title),
	).Scan(&existingId)

	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar existencia del contenido",
		})
	}

	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Ya existe un contenido con ese título",
		})
	}

	_, err = db.DB.Exec(`
		INSERT INTO content_type (content_title, content_type, content_cover, content_year, gender_id)
		VALUES (?, ?, ?, ?, ?)
	`,
		strings.ToUpper(content.Content_Title),
		content.Content_Type,
		content.Content_Cover,
		content.Content_Year,
		content.Gender_Id,
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

func PUT_Content(c *fiber.Ctx) error {

	id := c.Params("id")

	var content models.Content

	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	var existingId int
	err := db.DB.QueryRow(`
		SELECT content_id 
		FROM content_type 
		WHERE UPPER(content_title) = ?
		AND content_id != ?
	`,
		strings.ToUpper(content.Content_Title),
		id,
	).Scan(&existingId)

	if err != nil && err != sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar existencia del contenido",
		})
	}

	if existingId != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Ya existe un contenido con ese título",
		})
	}

	_, err = db.DB.Exec(`
		UPDATE content_type SET
			content_title = ?,
			content_type = ?,
			content_cover = ?,
			content_year = ?,
			gender_id = ?
		WHERE content_id = ?
	`,
		strings.ToUpper(content.Content_Title),
		content.Content_Type,
		content.Content_Cover,
		content.Content_Year,
		content.Gender_Id,
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

func POST_Content_Season(c *fiber.Ctx) error {

	var data models.ContentSeason

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al iniciar la transacción",
		})
	}

	for _, episode := range data.Episodes {
		_, err := tx.Exec(`
			INSERT INTO episode (episode_number, episode_name, episode_url, season_id, content_id)
			VALUES (?, ?, ?, ?, ?)
		`,
			episode.Episode_Number,
			strings.ToUpper(episode.Episode_Name),
			episode.Episode_Url,
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

func PUT_Content_Season(c *fiber.Ctx) error {

	var data models.ContentSeason

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cuerpo de solicitud inválido",
		})
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al iniciar la transacción",
		})
	}

	_, err = tx.Exec(`
		DELETE FROM episode
		WHERE season_id = ? AND content_id = ?
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
			VALUES (?, ?, ?, ?, ?)
		`,
			episode.Episode_Number,
			strings.ToUpper(episode.Episode_Name),
			episode.Episode_Url,
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

func DELETE_Content(c *fiber.Ctx) error {

	id := c.Params("id")

	tx, err := db.DB.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al iniciar la transacción",
		})
	}

	_, err = tx.Exec(`
		DELETE FROM episode
		WHERE content_id = ?
	`, id)

	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar episodios asociados",
		})
	}

	result, err := tx.Exec(`
		DELETE FROM content_type
		WHERE content_id = ?
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
