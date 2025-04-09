package models

type Series struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`

	SeasonID uint    `json:"season_id"`
	Season   Seasons `json:"season" gorm:"foreignKey:SeasonID"`

	GenderID uint    `json:"gender_id"`
	Gender   Genders `json:"gender" gorm:"foreignKey:GendersID"`
}
