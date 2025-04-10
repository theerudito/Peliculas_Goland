package models

type Movie struct {
	ID       int    `json:"movie_id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Cover    string `json:"cover"`
	URL      string `json:"url"`
	GenderID uint   `json:"gender_id"`
}

type MovieDTO struct {
	ID     int    `json:"movie_id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Cover  string `json:"cover"`
	URL    string `json:"url"`
	Gender string `json:"gender"`
}
