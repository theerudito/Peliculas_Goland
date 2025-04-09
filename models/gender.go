package models

type Genders struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`

	Anime []Animes `json:"animes" gorm:"foreignKey:GendersID"`
	Movie []Movies `json:"movies" gorm:"foreignKey:GendersID"`
	Serie []Series `json:"series" gorm:"foreignKey:GendersID"`
}
