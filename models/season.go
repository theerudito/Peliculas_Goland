package models

type Season struct {
	Season_ID   int    `json:"season_id"`
	Title       string `json:"title"`
	Descripcion string `json:"descripcion"`
	Year        string `json:"year"`
}
