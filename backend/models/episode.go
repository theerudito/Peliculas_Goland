package models

type Episodie struct {
	Episode_Id         uint   `json:"episode_id"`
	Episode_Number     int    `json:"episode_number"`
	Episode_Name       string `json:"episode_name"`
	Episode_Video      []byte `json:"episode_video"`
	Episode_Season_Id  uint   `json:"episode_season_id"`
	Episode_Content_Id uint   `json:"episode_content_id"`
}

type EpisodieDTO struct {
	Episode_Id        uint   `json:"episode_id"`
	Episode_Number    int    `json:"episode_number"`
	Episode_Name      string `json:"episode_name"`
	Episode_Video     string `json:"episode_video"`
	Episode_Season_Id uint   `json:"episode_season_id"`
	Episode_Season    string `json:"episode_season"`
}
