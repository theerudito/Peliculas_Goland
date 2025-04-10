package models

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	GenderID uint   `json:"gender_id"`
}

type MovieDTO struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Gender string `json:"gender"`
}
