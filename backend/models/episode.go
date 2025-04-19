package models

type Episodie struct {
	Episode_Id     uint   `json:"episode_id"`
	Episode_Number int    `json:"episode_number"`
	Episode_Name   string `json:"episode_name"`
	Episode_Url    string `json:"episode_url"`
}

type EpisodieDTO struct {
	Episode_Id     uint   `json:"episode_id"`
	Episode_Number int    `json:"episode_number"`
	Episode_Name   string `json:"episode_name"`
	Episode_Url    string `json:"episode_url"`
}
