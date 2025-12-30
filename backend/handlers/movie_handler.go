package handlers

import (
	"database/sql"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
			COALESCE(c.url, '') AS cover,
			COALESCE(v.url, '') AS video,
			g.gender_name
		FROM movie AS m
			LEFT JOIN gender AS g ON m.gender_id = g.gender_id
			LEFT JOIN storage AS c ON m.cover_id = c.storage_id
    		LEFT JOIN storage AS v ON m.video_id = v.storage_id
	`)

	if err != nil {

		_ = helpers.InsertLogsError(conn, "movie", err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al ejecutar la consulta"})

	}

	defer rows.Close()

	for rows.Next() {
		var movie models.MovieDTO

		err := rows.Scan(
			&movie.Movie_Id,
			&movie.Movie_Title,
			&movie.Movie_Year,
			&movie.Movie_Cover,
			&movie.Movie_Video,
			&movie.Gender,
		)

		if err != nil {

			_ = helpers.InsertLogsError(conn, "movie", err.Error())

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al leer los registros"})
		}

		dto = append(dto, movie)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No se encontraron registros"})
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
			COALESCE(c.url, '') AS cover,
			COALESCE(v.url, '') AS video,
			g.gender_name
		FROM movie AS m
			INNER JOIN gender AS g ON m.gender_id = g.gender_id
			LEFT JOIN storage AS c ON m.cover_id = c.storage_id
    		LEFT JOIN storage AS v ON m.video_id = v.storage_id
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
			&movie.Movie_Video,
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
			COALESCE(c.url, '') AS cover,
			COALESCE(v.url, '') AS video,
			g.gender_name
		FROM movie AS m
			INNER JOIN gender AS g ON m.gender_id = g.gender_id
			LEFT JOIN storage AS c ON m.cover_id = c.storage_id
    		LEFT JOIN storage AS v ON m.video_id = v.storage_id
		WHERE m.movie_title LIKE $1
`, search)

	if err != nil {

		_ = helpers.InsertLogsError(conn, "movie", err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al ejecutar la consulta"})
	}

	defer rows.Close()

	for rows.Next() {
		var movie models.MovieDTO

		err := rows.Scan(
			&movie.Movie_Id,
			&movie.Movie_Title,
			&movie.Movie_Year,
			&movie.Movie_Cover,
			&movie.Movie_Video,
			&movie.Gender,
		)
		if err != nil {

			_ = helpers.InsertLogsError(conn, "movie", err.Error())

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al leer los registros"})
		}

		dto = append(dto, movie)
	}

	if len(dto) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No se encontraron registros"})
	}

	return c.JSON(dto)

}

func PostMovie(c *fiber.Ctx) error {
	var (
		coverID int
		videoID int
		movieID int
	)

	conn := db.GetDB()

	// =============================
	// 1️⃣ LEER CAMPOS DE TEXTO
	// =============================
	movieTitle := strings.ToUpper(c.FormValue("movie_title"))
	movieYear, err := strconv.Atoi(c.FormValue("movie_year"))

	if err != nil {
		return c.Status(400).JSON("movie_year inválido")
	}

	genderID, err := strconv.Atoi(c.FormValue("gender_id"))

	if err != nil {
		return c.Status(400).JSON("gender_id inválido")
	}

	// =============================
	// 2️⃣ VALIDAR DUPLICADO
	// =============================
	var existingID int
	err = conn.QueryRow(
		"SELECT movie_id FROM movie WHERE movie_title = $1",
		movieTitle,
	).Scan(&existingID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.Status(500).JSON("error al validar duplicado")
	}

	if existingID != 0 {
		return c.Status(409).JSON("la película ya existe")
	}

	// =============================
	// 3️⃣ LEER ARCHIVOS
	// =============================
	coverFile, err := c.FormFile("cover")

	if err != nil {
		return c.Status(400).JSON("cover requerido")
	}

	videoFile, err := c.FormFile("video")

	if err != nil {
		return c.Status(400).JSON("video requerido")
	}

	// =============================
	// 4️⃣ FUNCIÓN PARA LEER ARCHIVO
	// =============================
	readFile := func(fh *multipart.FileHeader) ([]byte, []byte, error) {

		src, err := fh.Open()
		if err != nil {
			return nil, nil, err
		}

		defer src.Close()

		header := make([]byte, 512)
		_, _ = src.Read(header)

		_, _ = src.Seek(0, 0)

		data, err := io.ReadAll(src)
		return data, header, err

	}

	coverData, coverHeader, err := readFile(coverFile)

	if err != nil {
		return c.Status(500).JSON("error leyendo cover")
	}

	videoData, videoHeader, err := readFile(videoFile)

	if err != nil {
		return c.Status(500).JSON("error leyendo video")
	}

	// =============================
	// 5️⃣ DETECTAR EXTENSIONES
	// =============================
	coverExt := helpers.InfoExtention(coverHeader)
	videoExt := helpers.InfoExtention(videoHeader)

	if coverExt == "" || videoExt == "" {
		return c.Status(400).JSON("tipo de archivo no permitido")
	}

	// =============================
	// 6️⃣ TRANSACCIÓN
	// =============================
	tx, err := conn.Begin()

	if err != nil {
		return c.Status(500).JSON("error iniciando transacción")
	}

	defer tx.Rollback()

	// =============================
	// 7️⃣ GUARDAR ARCHIVOS + STORAGE
	// =============================
	files := []struct {
		Data []byte
		Ext  string
		Type string
	}{
		{coverData, coverExt, "cover"},
		{videoData, videoExt, "video"},
	}

	for _, f := range files {

		dir := map[string]string{
			"cover": os.Getenv("Images"),
			"video": os.Getenv("Videos"),
		}[f.Type]

		fileName := uuid.New().String()

		url, err := helpers.SaveImageToDirectory(f.Data, fileName, f.Ext, dir)

		if err != nil {
			return c.Status(500).JSON("error guardando archivo")
		}

		var storageID int
		err = tx.QueryRow(`
			INSERT INTO storage (file_name, url, extension)
			VALUES ($1, $2, $3)
			RETURNING storage_id`,
			fileName,
			url,
			f.Ext,
		).Scan(&storageID)

		if err != nil {
			return c.Status(500).JSON("error insertando storage")
		}

		if f.Type == "cover" {
			coverID = storageID
		} else {
			videoID = storageID
		}
	}

	// =============================
	// 8️⃣ INSERTAR MOVIE
	// =============================
	err = tx.QueryRow(`
		INSERT INTO movie (
			movie_title,
			movie_year,
			cover_id,
			video_id,
			gender_id
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING movie_id`,
		movieTitle,
		movieYear,
		coverID,
		videoID,
		genderID,
	).Scan(&movieID)

	if err != nil {
		return c.Status(500).JSON("error insertando movie")
	}

	// =============================
	// 9️⃣ COMMIT
	// =============================
	if err := tx.Commit(); err != nil {
		return c.Status(500).JSON("error confirmando transacción")
	}

	return c.Status(201).JSON(fiber.Map{"message": "Película creada correctamente 🚀"})

}

func PutMovie(c *fiber.Ctx) error {

	panic("")
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
