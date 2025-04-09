package models

type Movie struct {
	Title    string
	Year     int
	GenderID uint
	Gender   Gender
}

type MovieDTO struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Gender string `json:"gender"`
}
