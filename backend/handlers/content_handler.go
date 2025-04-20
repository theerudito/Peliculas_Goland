package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func GetContent(c *fiber.Ctx) error {

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
	SELECT
	content.content_id,
	content.content_title,
	content.content_cover,
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
		err := rows.Scan(
			&content.Content_Id,
			&content.Content_Title,
			&content.Content_Cover,
			&content.Content_Year,
			&content.Gender,
			&content.Content_Type)
		if err != nil {
			return err
		}
		dto = append(dto, content)
	}
	return c.JSON(dto)

}

func GetContentID(c *fiber.Ctx) error {

	var dto models.ContentDTO

	pID := c.Params("id")

	err := db.DB.QueryRow(`
	SELECT
		content.content_id,
		content.content_title,
		content.content_cover,
		content.content_year,
		gender.gender_name,
		CASE
			WHEN content.content_type = 1 THEN 'ANIME'
			ELSE 'SERIE'
		END AS type
	FROM content_types content
	INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
	WHERE content.content_id = ?
`, pID).Scan(
		&dto.Content_Id,
		&dto.Content_Title,
		&dto.Content_Cover,
		&dto.Content_Year,
		&dto.Gender,
		&dto.Content_Type,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Contenido no encontrado ❌",
			})
		}
		return err
	}

	return c.JSON(dto)

}

func GetContenType(c *fiber.Ctx) error {
	value := c.Params("type")

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
		SELECT
			content.content_id,
			content.content_title,
			content.content_cover,
			content.content_year,
			gender.gender_name,
			CASE
				WHEN content.content_type = 1 THEN 'ANIME'
				ELSE 'SERIE'
			END AS type
		FROM content_types content
		INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
		WHERE content.content_type = ?
	`, value)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error al consultar la base de datos",
			"details": err.Error(),
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
			&content.Gender,
			&content.Content_Type,
		)
		if err != nil {
			return err
		}
		dto = append(dto, content)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No se encontró contenido para el tipo especificado",
		})
	}

	return c.JSON(dto)
}

func FindContent(c *fiber.Ctx) error {
	value := c.Params("value")

	search := "%" + strings.ToUpper(value) + "%"

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
		SELECT
			content.content_id,
			content.content_title,
			content.content_cover,
			content.content_year,
			gender.gender_name,
			CASE
				WHEN content.content_type = 1 THEN 'ANIME'
				ELSE 'SERIE'
			END AS type
		FROM content_types content
		INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
		WHERE UPPER(content.content_title) LIKE ?
	`, search)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error al buscar contenido",
			"details": err.Error(),
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
			&content.Gender,
			&content.Content_Type,
		)
		if err != nil {
			return err
		}
		dto = append(dto, content)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No se encontró contenido con ese título ❌",
		})
	}

	return c.JSON(dto)
}

func GetContentData(c *fiber.Ctx) error {

	value := c.Params("value")

	var dto []models.ContentData

	rows, err := db.DB.Query(`
		SELECT
			content.content_id,
			content.content_title,
			content.content_type,
			content.content_cover,
			content.content_year,
			content.gender_id,

			season.season_id,
			season.season_name,

			episode.episode_id,
			episode.episode_number,
			episode.episode_name,
			episode.episode_url

		FROM content_types AS content
		INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
		LEFT JOIN seasons AS season ON season.content_id = content.content_id
		LEFT JOIN episodes AS episode ON episode.season_id = season.season_id
		WHERE season.season_id = ?
		ORDER BY season.season_id, episode.episode_number;
	`, value)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error al consultar la base de datos",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	// Map para agrupar por Season ID
	contentMap := make(map[uint]*models.ContentData)

	for rows.Next() {
		var (
			content     models.Content
			season      models.Season
			episode     models.Episodie
			seasonID    sql.NullInt64
			seasonName  sql.NullString
			episodeID   sql.NullInt64
			episodeNum  sql.NullInt64
			episodeName sql.NullString
			episodeURL  sql.NullString
		)

		err := rows.Scan(
			&content.Content_Id,
			&content.Content_Title,
			&content.Content_Type,
			&content.Content_Cover,
			&content.Content_Year,
			&content.Gender_Id,

			&seasonID,
			&seasonName,

			&episodeID,
			&episodeNum,
			&episodeName,
			&episodeURL,
		)

		if err != nil {
			return err
		}

		// Validamos la Season
		if seasonID.Valid {
			season.Season_Id = uint(seasonID.Int64)
			season.Season_Name = seasonName.String
		}

		// Validamos Episodio
		if episodeID.Valid {
			episode = models.Episodie{
				Episode_Id:     uint(episodeID.Int64),
				Episode_Number: int(episodeNum.Int64),
				Episode_Name:   episodeName.String,
				Episode_Url:    episodeURL.String,
			}
		}

		// Agrupamos por Season ID
		if existing, ok := contentMap[season.Season_Id]; ok {
			if episodeID.Valid {
				existing.Seasons.Episodes = append(existing.Seasons.Episodes, episode)
			}
		} else {
			newData := models.ContentData{
				Content: content,
				Seasons: models.Season{
					Season_Id:   season.Season_Id,
					Season_Name: season.Season_Name,
				},
			}

			if episodeID.Valid {
				newData.Seasons.Episodes = append(newData.Seasons.Episodes, episode)
			}

			contentMap[season.Season_Id] = &newData
		}
	}

	// Convertimos el mapa a slice
	for _, v := range contentMap {
		dto = append(dto, *v)
	}

	return c.JSON(dto)
}

func PostContent(c *fiber.Ctx) error {

	var content models.ContentData

	if err := c.BodyParser(&content); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error al parsear JSON"})
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error iniciando transacción"})
	}

	res, err := tx.Exec(`
	INSERT INTO content_types (content_title, content_type, content_cover, content_year, gender_id)
	VALUES (?, ?, ?, ?, ?)`,
		strings.ToUpper(content.Content.Content_Title),
		content.Content.Content_Type,
		content.Content.Content_Cover,
		content.Content.Content_Year,
		content.Content.Gender_Id,
	)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Error insertando contenido", "details": err.Error()})
	}

	contentID, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Error obteniendo content_id", "details": err.Error()})
	}

	res, err = tx.Exec(`
		INSERT INTO seasons (season_name, content_id)
		VALUES (?, ?)`,
		strings.ToUpper(content.Seasons.Season_Name),
		contentID,
	)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Error insertando temporada", "details": err.Error()})
	}

	seasonID, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Error obteniendo season_id", "details": err.Error()})
	}

	for _, ep := range content.Seasons.Episodes {
		_, err := tx.Exec(`
			INSERT INTO episodes (episode_number, episode_name, episode_url, season_id)
			VALUES (?, ?, ?, ?)`,
			ep.Episode_Number,
			strings.ToUpper(ep.Episode_Name),
			ep.Episode_Url,
			seasonID,
		)
		if err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{"error": "Error insertando episodio", "details": err.Error()})
		}
	}

	if err := tx.Commit(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al confirmar transacción", "details": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Contenido creado correctamente 🚀"})
}

func PutContent(c *fiber.Ctx) error {

	// var content models.ContenData

	// if err := c.BodyParser(&content); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": "Error al parsear JSON"})
	// }

	// tx, err := db.DB.Begin()
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"error": "Error iniciando transacción"})
	// }

	// _, err = tx.Exec(`
	// 	UPDATE content_types SET content_title = ?, content_type_id = ?, content_cover = ?, content_year = ?, gender_id = ?
	// 	WHERE content_id=?`,
	// 	strings.ToUpper(content.Content_Title),
	// 	content.Content_Type_Id,
	// 	content.Content_Cover,
	// 	content.Content_Year,
	// 	content.Gender_Id,
	// )
	// if err != nil {
	// 	tx.Rollback()
	// 	return c.Status(500).JSON(fiber.Map{"error": "Error actualizando contenido"})
	// }

	// for _, ep := range content.Seasons[0].Episodes {
	// 	_, err := tx.Exec(`
	// 		UPDATE episodes
	// 		SET episode_number=?, episode_name=?, episode_url=?
	// 		WHERE episode_id=? AND season_id=?`,
	// 		ep.Episode_Number,
	// 		ep.Episode_Name,
	// 		ep.Episode_Url,
	// 		ep.Episode_Id,
	// 		ep.Season_Id,
	// 	)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return c.Status(500).JSON(fiber.Map{"error": "Error actualizando episodios"})
	// 	}
	// }

	// tx.Commit()

	return c.Status(200).JSON(fiber.Map{"message": "Contenido actualizado correctamente ✅"})

}

func DeleteContent(c *fiber.Ctx) error {

	contentID := c.Params("id")

	if contentID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Falta content_id"})
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error iniciando transacción"})
	}

	_, err = tx.Exec(`
		DELETE FROM episodes
		WHERE EXISTS (
			SELECT 1 FROM seasons 
			WHERE seasons.season_id = episodes.season_id 
			AND seasons.content_id = ?
		)
	`, contentID)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error borrando episodios",
			"details": err.Error(),
		})
	}

	_, err = tx.Exec(`
		DELETE FROM seasons
		WHERE content_id = ?
	`, contentID)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error borrando temporadas",
			"details": err.Error(),
		})
	}

	res, err := tx.Exec(`
		DELETE FROM content_types
		WHERE content_id = ?
	`, contentID)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error borrando contenido",
			"details": err.Error(),
		})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error al obtener filas afectadas",
			"details": err.Error(),
		})
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return c.Status(404).JSON(fiber.Map{
			"error":      "Contenido no encontrado ❌",
			"content_id": contentID,
		})
	}

	if err := tx.Commit(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error al confirmar transacción",
			"details": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Contenido eliminado correctamente 🗑️",
	})

}
