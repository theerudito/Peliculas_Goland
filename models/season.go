package models

type Seasons struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	Number int  `json:"number"`

	EpisodeID uint     `json:"episode_id"`
	Episode   Episodes `json:"episode" gorm:"foreignKey:EpisodeID"`

	Anime []Animes `json:"animes" gorm:"foreignKey:SeasonID"`
	Serie []Series `json:"series" gorm:"foreignKey:SeasonID"`
}
