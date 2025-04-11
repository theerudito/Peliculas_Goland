package models

type Content_Type struct {
	Content_Type_ID int    `json:"content_type_id"`
	Title           string `json:"title"`
	Descripcion     string `json:"descripcion"`
	Cover           string `json:"cover"`
	Year            string `json:"year"`
}
