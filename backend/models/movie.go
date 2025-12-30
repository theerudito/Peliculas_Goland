package models

type Movie struct {
	Movie_Id    uint   `json:"movie_id"`
	Movie_Title string `json:"movie_title"`
	Movie_Year  int    `json:"movie_year"`
	Cover       []byte `json:"cover"`
	Video       []byte `json:"video"`
	Gender_Id   uint   `json:"gender_id"`
}

type MovieDTO struct {
	Movie_Id    uint   `json:"movie_id"`
	Movie_Title string `json:"movie_title"`
	Movie_Year  int    `json:"movie_year"`
	Movie_Cover string `json:"movie_cover"`
	Movie_Video string `json:"movie_video"`
	Gender      string `json:"gender"`
}
