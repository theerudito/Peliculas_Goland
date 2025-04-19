package models

type Content struct {
	Content_Id      int    `json:"content_id"`
	Content_Type_Id int    `json:"content_type_id"`
	Content_Cover   string `json:"content_cover"`
	Content_Url     string `json:"content_url"`
	Content_Year    string `json:"content_year"`
	Gender_Id       uint   `json:"gender_id"`
}

type ContentDTO struct {
	Content_Id    int    `json:"content_id"`
	Content_Type  string `json:"content_type"`
	Content_Cover string `json:"content_cover"`
	Content_Url   string `json:"content_url"`
	Content_Year  string `json:"content_year"`
	Gender        string `json:"gender"`
}
