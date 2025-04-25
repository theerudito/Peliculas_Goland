package models

type Season struct {
	Season_Id   uint   `json:"season_id"`
	Season_Name string `json:"season_name"`
}

type SeasonDTO struct {
	Season_Id   uint          `json:"season_id"`
	Season_Name string        `json:"season_name"`
	Episodes    []EpisodieDTO `json:"episodes"`
}
