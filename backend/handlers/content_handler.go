package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
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
	season.season_id,
	season.season_name,
	gender.gender_name,
	CASE
			WHEN content.content_type = 1 THEN 'ANIME'
			ELSE 'SERIE'
	END AS type
	FROM content_types content
	INNER JOIN seasons AS season ON season.content_id = content.content_id 
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
			&content.Season_Id,
			&content.Season_Name,
			&content.Content_Gender,
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
		season.season_id,
		season.season_name,
		gender.gender_name,
		CASE
			WHEN content.content_type = 1 THEN 'ANIME'
			ELSE 'SERIE'
		END AS type
	FROM content_types content
	INNER JOIN seasons AS season ON season.content_id = content.content_id
	INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
	WHERE content.content_id = ?
`, pID).Scan(
		&dto.Content_Id,
		&dto.Content_Title,
		&dto.Content_Cover,
		&dto.Content_Year,
		&dto.Season_Id,
		&dto.Season_Name,
		&dto.Content_Gender,
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

	value := (c.Params("id"))

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
	SELECT
	MIN(content.content_id) AS content_id,
	content.content_title,
	MIN(content.content_cover) AS content_cover,
	MIN(content.content_year) AS content_year,
	season.season_id,
	season.season_name,
	MIN(gender.gender_name) AS gender_name,
	CASE
		WHEN content.content_type = 1 THEN 'ANIME'
		ELSE 'SERIE'
	END AS type
	FROM content_types content
	INNER JOIN seasons AS season ON season.content_id = content.content_id
	INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
	WHERE content_type = ?
	GROUP BY content.content_title, content.content_type
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
			&content.Season_Id,
			&content.Season_Name,
			&content.Content_Gender,
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

func GetContentSeason(c *fiber.Ctx) error {

	value := helpers.QuitarGuiones(c.Params("value"))

	search := "%" + strings.ToUpper(value) + "%"

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
	SELECT
	content.content_id,
	content.content_title,
	content.content_cover,
	content.content_year,
	season.season_id,
	season.season_name,
	CASE
	WHEN content.content_type = 1 THEN 'ANIME'
	ELSE 'SERIE'
	END AS type,
	gender.gender_name
	FROM content_types AS content
	INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
	INNER JOIN seasons AS season ON season.content_id = content.content_id
	WHERE UPPER(content.content_title) LIKE ?
	`,
		search)

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
			&content.Season_Id,
			&content.Season_Name,
			&content.Content_Type,
			&content.Content_Gender,
		)
		if err != nil {
			return err
		}
		dto = append(dto, content)
	}

	return c.JSON(dto)

}

func GetContentEpisode(c *fiber.Ctx) error {
	value := helpers.QuitarGuiones(c.Params("value"))
	idSeason := c.Params("id")
	search := "%" + strings.ToUpper(value) + "%"

	var dto []models.ContentDataDTO

	rows, err := db.DB.Query(`
		SELECT 
			ct.content_id,
			ct.content_title,
			CASE ct.content_type
				WHEN 1 THEN 'SERIE'
				WHEN 2 THEN 'ANIME'
			END as content_type,
			ct.content_cover,
			ct.content_year,
			g.gender_name,
			
			s.season_id,
			s.season_name,

			e.episode_id,
			e.episode_number,
			e.episode_name,
			e.episode_url
		FROM 
			content_types ct
		JOIN 
			genders g ON g.gender_id = ct.gender_id
		JOIN 
			seasons s ON s.content_id = ct.content_id
		JOIN 
			episodes e ON e.season_id = s.season_id
		WHERE 
			UPPER(ct.content_title) LIKE ?
			AND s.season_id = ?
		ORDER BY 
			s.season_id, e.episode_number
	`, search, idSeason)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error al consultar la base de datos",
			"details": err.Error(),
		})
	}
	defer rows.Close()

	contentMap := make(map[uint]*models.ContentDataDTO)

	for rows.Next() {
		var (
			contentID     uint
			contentTitle  string
			contentType   string
			contentCover  string
			contentYear   int
			genderName    string
			seasonID      uint
			seasonName    string
			episodeID     uint
			episodeNumber int
			episodeName   string
			episodeURL    string
		)

		err := rows.Scan(
			&contentID, &contentTitle, &contentType, &contentCover, &contentYear, &genderName,
			&seasonID, &seasonName,
			&episodeID, &episodeNumber, &episodeName, &episodeURL,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Error al escanear los resultados",
				"details": err.Error(),
			})
		}

		// Buscar o crear entrada del contenido
		contentData, exists := contentMap[contentID]
		if !exists {
			contentData = &models.ContentDataDTO{
				Content: models.ContentDTO{
					Content_Id:     contentID,
					Content_Title:  contentTitle,
					Content_Type:   contentType,
					Content_Cover:  contentCover,
					Content_Year:   contentYear,
					Content_Gender: genderName,
					Season_Id:      seasonID,
					Season_Name:    seasonName,
				},
				Seasons: []models.SeasonDTO{},
			}
			contentMap[contentID] = contentData
		}

		// Buscar la temporada dentro del contenido
		var season *models.SeasonDTO
		for i := range contentData.Seasons {
			if contentData.Seasons[i].Season_Id == seasonID {
				season = &contentData.Seasons[i]
				break
			}
		}

		// Si no existe, la creamos
		if season == nil {
			newSeason := models.SeasonDTO{
				Season_Id:   seasonID,
				Season_Name: seasonName,
				Episodes:    []models.EpisodieDTO{},
			}
			contentData.Seasons = append(contentData.Seasons, newSeason)
			season = &contentData.Seasons[len(contentData.Seasons)-1]
		}

		// Agregar episodio a la temporada
		season.Episodes = append(season.Episodes, models.EpisodieDTO{
			Episode_Id:     episodeID,
			Episode_Number: episodeNumber,
			Episode_Name:   episodeName,
			Episode_Url:    episodeURL,
		})
	}

	// Convertir el mapa a slice
	for _, v := range contentMap {
		dto = append(dto, *v)
	}

	return c.JSON(dto)
}

func FindContent(c *fiber.Ctx) error {

	value := helpers.QuitarGuiones(c.Params("value"))

	search := "%" + strings.ToUpper(value) + "%	"

	var dto []models.ContentDTO

	rows, err := db.DB.Query(`
		SELECT
			content.content_id,
			content.content_title,
			content.content_cover,
			content.content_year,
			season.season_id,
			season.season_name,
			gender.gender_name,
			CASE
				WHEN content.content_type = 1 THEN 'ANIME'
				ELSE 'SERIE'
			END AS type
		FROM content_types content
		INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
		INNER JOIN genders AS gender ON gender.gender_id = content.gender_id
		WHERE UPPER(content.content_title) LIKE ?
		AND content.content_type = ?
		GROUP BY content.content_title, gender.gender_name, content.content_type

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
			&content.Season_Id,
			&content.Season_Name,
			&content.Content_Gender,
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

	for _, ep := range content.Episodie {
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
