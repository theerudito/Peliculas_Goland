package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	db "github.com/theerudito/peliculas/database"
	"github.com/theerudito/peliculas/models"
)

func Get_Movies(c *fiber.Ctx) error {

	var dto []models.MovieDTO

	rows, err := db.DB.Query(`SELECT movie.movie_movie_id, movie.movie_title, movie.movie_year,movie.movie_cover, movie.movie_url, gender.name
	FROM movies AS movie
	INNER JOIN genders AS gender ON movie.movie_id = gender.gender_id`)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.MovieDTO
		err := rows.Scan(&movie.Movie_Movie_Id, &movie.Movie_Title, &movie.Movie_Year, &movie.Movie_Cover, &movie.Movie_Url, &movie.Gender)
		if err != nil {
			return err
		}
		dto = append(dto, movie)
	}
	return c.JSON(dto)
}

func Get_MovieByID(c *fiber.Ctx) error {

	id := c.Params("id")

	var movie models.MovieDTO

	err := db.DB.QueryRow(`SELECT movie.movie_id, movie.title, movie.year,movie.cover, movie.url, gender.descripcion
	FROM movies AS movie
	INNER JOIN genders AS gender ON movie.movie_id = gender.gender_id
	WHERE movie.movie_id = ?`, id).Scan(&movie.Movie_Movie_Id, &movie.Movie_Title, &movie.Movie_Year, &movie.Movie_Cover, &movie.Movie_Url, &movie.Gender)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Movie not found"})
		}
		return err
	}
	return c.JSON(movie)
}

func Post_Movie(c *fiber.Ctx) error {

	movie := new(models.Movie)

	if err := c.BodyParser(movie); err != nil {
		return err
	}

	_, err := db.DB.Exec(`INSERT INTO movies (title, year, gender_id) 
	VALUES (?, ?, ?)`, movie.Movie_Title, movie.Movie_Cover, movie.Movie_Year, movie.Movie_Url, movie.Gender_Id)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"message": "Movie added"})
}

func Put_Movies(c *fiber.Ctx) error {

	id := c.Params("id")

	movie := new(models.Movie)

	if err := c.BodyParser(movie); err != nil {
		return err
	}

	_, err := db.DB.Exec(`UPDATE movies SET 
		title = ?, year = ?, gender_id = ? WHERE id = ?`,
		movie.Movie_Title, movie.Movie_Cover, movie.Movie_Year, movie.Movie_Url, movie.Gender_Id, id)

	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Movie Updated"})
}

func Delete_Movies(c *fiber.Ctx) error {

	id := c.Params("id")

	_, err := db.DB.Exec("DELETE FROM movies WHERE id = ?", id)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Movie Deleted"})
}
