package models

type Content struct {
	Content_Id    uint   `json:"content_id"`
	Content_Title string `json:"content_title"`
	Content_Type  int    `json:"content_type"`
	Content_Cover string `json:"content_cover"`
	Content_Year  int    `json:"content_year"`
	Gender_Id     uint   `json:"gender_id"`
}

type ContentDTO struct {
	Content_Id    uint   `json:"content_id"`
	Content_Title string `json:"content_title"`
	Content_Type  int    `json:"content_type"`
	Content_Cover string `json:"content_cover"`
	Content_Year  int    `json:"content_year"`
	Gender        string `json:"gender"`
}

type ContentData struct {
	Content Content `json:"content"`
	Seasons Season  `json:"seasons"`
}
