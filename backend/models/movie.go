package models

type Movie struct {
	Movie_Id        uint   `json:"movie_id"`
	Movie_Title     string `json:"movie_title"`
	Movie_Year      int    `json:"movie_year"`
	Movie_Cover     []byte `json:"movie_cover"`
	Movie_Video     []byte `json:"movie_video"`
	Movie_Gender_Id uint   `json:"movie_gender_id"`
}

type MovieDTO struct {
	Movie_Id        uint   `json:"movie_id"`
	Movie_Title     string `json:"movie_title"`
	Movie_Year      int    `json:"movie_year"`
	Movie_Cover     string `json:"movie_cover"`
	Movie_Video     string `json:"movie_video"`
	Movie_Gender_Id uint   `json:"movie_gender_id"`
	Movie_Gender    string `json:"movie_gender"`
}
