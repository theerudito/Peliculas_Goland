package models

type Example_EpisodeDTO struct {
	Episode_Id     uint   `json:"episode_id"`
	Episode_Number int    `json:"episode_number"`
	Episode_Name   string `json:"episode_name"`
	Episode_Url    string `json:"episode_url"`
}

type Example_SeasonDTO struct {
	Season_Id   uint                 `json:"season_id"`
	Season_Name string               `json:"season_name"`
	Episodes    []Example_EpisodeDTO `json:"episodes"`
}

type Example_FullContentDTO struct {
	Content_Id     uint                `json:"content_id"`
	Content_Title  string              `json:"content_title"`
	Content_Type   string              `json:"content_type"`
	Content_Cover  string              `json:"content_cover"`
	Content_Year   int                 `json:"content_year"`
	Content_Gender string              `json:"content_gender"`
	Seasons        []Example_SeasonDTO `json:"seasons"`
}
