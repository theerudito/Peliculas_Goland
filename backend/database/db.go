package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type MyDB struct {
	DB *sql.DB
}

var instance *MyDB

func InitDB() {
	if instance != nil {
		return
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("✅ Conectado a la base de datos:", os.Getenv("DB_DRIVER"))

	instance = &MyDB{DB: db}
}

func GetDB() *sql.DB {
	if instance == nil {
		InitDB()
	}
	return instance.DB
}
