package models

type Episodie struct {
	Episode_ID     int    `json:"episode_id"`
	Episode_Number string `json:"episode_number"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}
