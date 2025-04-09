package models

type Anime struct {
	Title    string
	Year     int
	GenderID uint
	Gender   Gender
	Seasons  []Season
}
