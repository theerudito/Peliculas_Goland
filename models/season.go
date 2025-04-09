package models

type Season struct {
	Title    string
	AnimeID  *uint // Nullable: puede ser de anime o series
	SeriesID *uint
	Episodes []Episode
}
