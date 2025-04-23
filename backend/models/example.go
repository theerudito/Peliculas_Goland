package models

type ExampleContent struct {
	Content_Id    uint   `json:"content_id"`
	Content_Title string `json:"content_title"`
	Content_Type  string `json:"content_type"`
	Content_Cover string `json:"content_cover"`
	Content_Year  int    `json:"content_year"`
	Gender        string `json:"gender"`
}

type ExampleSeason struct {
	Season_Id   uint              `json:"season_id"`
	Season_Name string            `json:"season_name"`
	Episodes    []ExampleEpisodie `json:"episodes"`
}

type ExampleEpisodie struct {
	Episode_Id     uint   `json:"episode_id"`
	Episode_Number int    `json:"episode_number"`
	Episode_Name   string `json:"episode_name"`
	Episode_Url    string `json:"episode_url"`
}

type ExampleData struct {
	Content ExampleContent  `json:"content"`
	Seasons []ExampleSeason `json:"seasons"`
}
