package models

type Episodie struct {
	Episode_ID     int    `json:"episode_id"`
	Episode_Number string `json:"episode_number"`
	Episode_Name   string `json:"episode_name"`
	Episode_Url    string `json:"episode_url"`
	Season_Id      int    `json:"season_id"`
}

type EpisodieDTO []struct {
	Episode_ID     int    `json:"episode_id"`
	Episode_Number string `json:"episode_number"`
	Episode_Name   string `json:"episode_name"`
	Episode_Url    string `json:"episode_url"`
	Season_Id      int    `json:"season_id"`
}
