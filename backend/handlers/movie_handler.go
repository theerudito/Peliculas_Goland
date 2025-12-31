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

	movieTitle := strings.ToUpper(c.FormValue("movie_title"))
	movieYear, err := strconv.Atoi(c.FormValue("movie_year"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"messaje": "movie_year inválido"})
	}

	genderID, err := strconv.Atoi(c.FormValue("gender_id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"messaje": "gender_id inválido"})
	}

	var existingID int
	err = conn.QueryRow("SELECT movie_id FROM movie WHERE movie_title = $1", movieTitle).Scan(&existingID)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.Status(500).JSON(fiber.Map{"messaje": "error al validar duplicado"})
	}

	if existingID != 0 {
		return c.Status(409).JSON(fiber.Map{"messaje": "la película ya existe"})
	}

	coverFile, err := c.FormFile("cover")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"messaje": "cover requerido"})
	}

	videoFile, err := c.FormFile("video")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"messaje": "video requerido"})
	}

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
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error leyendo cover"})
	}

	videoData, videoHeader, err := readFile(videoFile)

	if err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error leyendo video"})
	}

	coverExt := helpers.InfoExtention(coverHeader)
	videoExt := helpers.InfoExtention(videoHeader)

	if coverExt == "" || videoExt == "" {
		return c.Status(400).JSON(fiber.Map{"messaje": "tipo de archivo no permitido"})
	}

	tx, err := conn.Begin()

	if err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error iniciando transacción"})
	}

	defer tx.Rollback()

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
			_ = helpers.InsertLogsError(conn, "movie", err.Error())
			return c.Status(500).JSON(fiber.Map{"messaje": "error guardando archivo"})
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
			_ = helpers.InsertLogsError(conn, "movie", err.Error())
			return c.Status(500).JSON(fiber.Map{"messaje": "error insertando storage"})
		}

		if f.Type == "cover" {
			coverID = storageID
		} else {
			videoID = storageID
		}
	}

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
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error insertando movie"})
	}

	if err := tx.Commit(); err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error confirmando transacción"})
	}

	_ = helpers.InsertLogs(conn, "INSERT", "movie", movieID, "Película creada correctamente 🚀")

	return c.Status(201).JSON(fiber.Map{"message": "Película creada correctamente 🚀"})

}

func PutMovie(c *fiber.Ctx) error {

	var (
		oldCoverID, oldVideoID int
		oldCoverFileToDelete   string
		oldVideoFileToDelete   string
	)

	conn := db.GetDB()
	var movie models.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "body inválido"})
	}

	if movie.Movie_Id == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "movie_id requerido"})
	}

	movieTitle := strings.ToUpper(movie.Movie_Title)
	movieYear := movie.Movie_Year
	genderID := movie.Gender_Id

	err := conn.QueryRow(`
		SELECT cover_id, video_id
		FROM movie
		WHERE movie_id = $1`,
		movie.Movie_Id,
	).Scan(&oldCoverID, &oldVideoID)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return c.Status(404).JSON(fiber.Map{"message": "la película no existe"})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error validando película"})
	}

	coverFile, _ := c.FormFile("cover")
	videoFile, _ := c.FormFile("video")

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

	tx, err := conn.Begin()

	if err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"message": "error iniciando transacción"})
	}

	defer tx.Rollback()

	/* =========================
	   COVER
	========================= */
	if coverFile != nil && coverFile.Size > 0 {

		err = tx.QueryRow(`
			SELECT CONCAT(file_name, extension)
			FROM storage
			WHERE storage_id = $1`,
			oldCoverID,
		).Scan(&oldCoverFileToDelete)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error obteniendo cover antiguo"})
		}

		coverData, coverHeader, err := readFile(coverFile)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error leyendo cover"})
		}

		coverExt := helpers.InfoExtention(coverHeader)
		fileName := uuid.New().String()

		url, err := helpers.SaveImageToDirectory(
			coverData,
			fileName,
			coverExt,
			os.Getenv("Images"),
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error guardando cover"})
		}

		_, err = tx.Exec(`
			UPDATE storage
			SET file_name = $1,
			    url = $2,
			    extension = $3
			WHERE storage_id = $4`,
			fileName,
			url,
			coverExt,
			oldCoverID,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error actualizando cover"})
		}
	}

	/* =========================
	   VIDEO
	========================= */
	if videoFile != nil && videoFile.Size > 0 {

		err = tx.QueryRow(`
			SELECT CONCAT(file_name, extension)
			FROM storage
			WHERE storage_id = $1`,
			oldVideoID,
		).Scan(&oldVideoFileToDelete)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error obteniendo video antiguo"})
		}

		videoData, videoHeader, err := readFile(videoFile)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error leyendo video"})
		}

		videoExt := helpers.InfoExtention(videoHeader)
		fileName := uuid.New().String()

		url, err := helpers.SaveImageToDirectory(
			videoData,
			fileName,
			videoExt,
			os.Getenv("Videos"),
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error guardando video"})
		}

		_, err = tx.Exec(`
			UPDATE storage
			SET file_name = $1,
			    url = $2,
			    extension = $3
			WHERE storage_id = $4`,
			fileName,
			url,
			videoExt,
			oldVideoID,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error actualizando video"})
		}
	}

	/* =========================
	   MOVIE
	========================= */
	_, err = tx.Exec(`
		UPDATE movie SET
			movie_title = $1,
			movie_year = $2,
			gender_id = $3
		WHERE movie_id = $4`,
		movieTitle,
		movieYear,
		genderID,
		movie.Movie_Id,
	)

	if err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"message": "error actualizando movie"})
	}

	if err := tx.Commit(); err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"message": "error confirmando transacción"})
	}

	/* =========================
	   DELETE FILES AFTER COMMIT
	========================= */
	if oldCoverFileToDelete != "" {
		_ = helpers.DeleteImageFromDirectory(
			oldCoverFileToDelete,
			os.Getenv("Images"),
		)
	}

	if oldVideoFileToDelete != "" {
		_ = helpers.DeleteImageFromDirectory(
			oldVideoFileToDelete,
			os.Getenv("Videos"),
		)
	}

	_ = helpers.InsertLogs(
		conn,
		"UPDATE",
		"movie",
		int(movie.Movie_Id),
		"Película actualizada correctamente ✨",
	)

	return c.Status(200).JSON(fiber.Map{
		"message": "Película actualizada correctamente ✨",
	})
}

func DeleteMovie(c *fiber.Ctx) error {

	var (
		coverID, videoID int
		storageIds       []int
	)

	type StorageFile struct {
		FileName  string
		Extension string
	}

	id, _ := strconv.Atoi(c.Params("id"))

	conn := db.GetDB()

	err := conn.QueryRow(`SELECT cover_id, video_id FROM movie WHERE movie_id = $1`, id).Scan(&coverID, &videoID)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return c.Status(404).JSON(fiber.Map{"messaje": "la película no existe"})
	}

	storageIds = append(storageIds, coverID, videoID)

	var files []StorageFile

	for _, storageID := range storageIds {

		var file StorageFile

		err = conn.QueryRow(`SELECT file_name, extension FROM storage WHERE storage_id = $1`, storageID).Scan(&file.FileName, &file.Extension)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"messaje": "error obteniendo datos"})
		}

		files = append(files, file)
	}

	tx, err := conn.Begin()

	if err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error iniciando transacción"})
	}

	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM movie WHERE movie_id = $1`, id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"messaje": "error eliminando movie"})
	}

	_, err = tx.Exec(`DELETE FROM storage WHERE storage_id IN ($1, $2)`, coverID, videoID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"messaje": "error eliminando movie"})
	}

	if err := tx.Commit(); err != nil {
		_ = helpers.InsertLogsError(conn, "movie", err.Error())
		return c.Status(500).JSON(fiber.Map{"messaje": "error confirmando transacción"})
	}

	for _, file := range files {
		switch file.Extension {
		case ".jpg", ".png", "webp":
			_ = helpers.DeleteImageFromDirectory(file.FileName+file.Extension, os.Getenv("Images"))
		case ".mp4":
			_ = helpers.DeleteImageFromDirectory(file.FileName+file.Extension, os.Getenv("Videos"))
		default:
			return c.Status(200).JSON(fiber.Map{"message": "formato incorrecto 🚀"})
		}
	}

	_ = helpers.InsertLogs(conn, "DELETE", "movie", id, "Película eliminada correctamente 🗑️")

	return c.Status(200).JSON(fiber.Map{"message": "Película eliminada correctamente 🗑️"})

}
