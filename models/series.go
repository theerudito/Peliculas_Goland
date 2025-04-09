package models

type Series struct {
	Title    string
	Year     int
	GenderID uint
	Gender   Gender
	Seasons  []Season
}
