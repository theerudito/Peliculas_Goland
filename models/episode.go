package models

type Episodes struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Number int    `json:"number"`

	Serie []Seasons `json:"season" gorm:"foreignKey:EpisodeID"`
}
