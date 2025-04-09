package db

import (
	"log"

	"github.com/theerudito/peliculas/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("media.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar la base de datos:", err)
	}

	database.AutoMigrate(&models.Movies{}, &models.Series{}, &models.Seasons{}, &models.Episodes{}, &models.Animes{}, &models.Genders{})

	DB = database
}
