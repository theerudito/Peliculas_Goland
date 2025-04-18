package handlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/helpers"
	"github.com/theerudito/peliculas/models"
)

func GetMovies(c *fiber.Ctx) error {

	var dto []models.MovieDTO

	rows, err := db.DB.Query(`SELECT 
	movie.movie_movie_id, 
	movie.movie_title, 
	movie.movie_year,
	movie.movie_cover, 
	movie.movie_url, 
	gender.gender_name
	FROM movies AS movie
	INNER JOIN genders AS gender ON movie.movie_movie_id = gender.gender_id`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.MovieDTO
		err := rows.Scan(
			&movie.Movie_Movie_Id,
			&movie.Movie_Title,
			&movie.Movie_Year,
			&movie.Movie_Cover,
			&movie.Movie_Url,
			&movie.Gender)

		if err != nil {
			return err
		}
		dto = append(dto, movie)
	}
	return c.JSON(dto)
}

func GetMoviebyid(c *fiber.Ctx) error {

	id := c.Params("id")

	var movie models.MovieDTO

	err := db.DB.QueryRow(`FROM movies AS movie
	INNER JOIN genders AS gender ON movie.movie_movie_id = gender.gender_id
	WHERE  movie.movie_movie_id = ?`, id).Scan(
		&movie.Movie_Movie_Id,
		&movie.Movie_Title,
		&movie.Movie_Year,
		&movie.Movie_Cover,
		&movie.Movie_Url,
		&movie.Gender)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Movie not found"})
		}
		return err
	}
	return c.JSON(movie)
}

func PostMovie(c *fiber.Ctx) error {

	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	movieYear, err := helpers.ConvertToInt(body["movie_year"])
	if err != nil {
		return fmt.Errorf("invalid movie_year: %v", err)
	}

	genderId, err := helpers.ConvertToUInt(body["gender_id"])
	if err != nil {
		return fmt.Errorf("invalid gender_id: %v", err)
	}

	movie := &models.Movie{
		Movie_Title: body["movie_title"].(string),
		Movie_Year:  movieYear,
		Movie_Cover: body["movie_cover"].(string),
		Movie_Url:   body["movie_url"].(string),
		Gender_Id:   genderId,
	}

	_, err = db.DB.Exec(`INSERT INTO movies (
	movie_title, 
	movie_year, 
	movie_cover, 
	movie_url, 
	gender_id)
		VALUES (?, ?, ?, ?, ?)`,
		strings.ToUpper(movie.Movie_Title),
		movie.Movie_Year,
		movie.Movie_Cover,
		movie.Movie_Url,
		movie.Gender_Id,
	)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"message": "Movie added"})
}

func PutMovies(c *fiber.Ctx) error {

	id := c.Params("id")

	movie := new(models.Movie)

	if err := c.BodyParser(movie); err != nil {
		return err
	}

	_, err := db.DB.Exec(`UPDATE movies SET
	movie_title = ?, 
	movie_year = ?, 
	movie_cover = ?, 
	movie_url = ?, 
	gender_id = ?`,
		strings.ToUpper(movie.Movie_Title),
		movie.Movie_Year,
		movie.Movie_Cover,
		movie.Movie_Url,
		movie.Gender_Id, id)

	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Movie Updated"})
}

func DeleteMovies(c *fiber.Ctx) error {

	id := c.Params("id")

	_, err := db.DB.Exec(`DELETE FROM movies 
	WHERE movie_movie_id = ?`, id)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Movie Deleted"})
}
