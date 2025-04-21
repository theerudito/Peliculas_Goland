package models

type Season struct {
	Season_Id   uint       `json:"season_id"`
	Season_Name string     `json:"season_name"`
	Episodes    []Episodie `json:"episodes"`
}

type SeasonDTO struct {
	Season_Id   uint   `json:"season_id"`
	Season_Name string `json:"season_name"`
}
