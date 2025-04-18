package models

type Movie struct {
	Movie_Movie_Id int    `json:"movie_movie_id"`
	Movie_Title    string `json:"movie_title"`
	Movie_Year     int    `json:"movie_year"`
	Movie_Cover    string `json:"movie_cover"`
	Movie_Url      string `json:"movie_url"`
	Gender_Id      uint   `json:"gender_id"`
}

type MovieDTO struct {
	Movie_Movie_Id int    `json:"movie_movie_id"`
	Movie_Title    string `json:"movie_title"`
	Movie_Year     int    `json:"movie_year"`
	Movie_Cover    string `json:"movie_cover"`
	Movie_Url      string `json:"movie_url"`
	Gender         string `json:"gender"`
}
