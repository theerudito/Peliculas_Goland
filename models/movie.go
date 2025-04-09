package models

type Movies struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`

	GenderID uint    `json:"gender_id"`
	Gender   Genders `json:"gender" gorm:"foreignKey:GendersID"`
}
