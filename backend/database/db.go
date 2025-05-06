package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	dbExists := fileExists("movies.db")

	var err error
	DB, err = sql.Open("sqlite3", "./movies.db")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("Conectado a la base de datos SQLite")

	// Ejecutar init.sql si la DB no existe o si no tiene tablas de usuario
	if !dbExists || isDBEmpty() {
		initSQL, err := os.ReadFile("init.sql")
		if err != nil {
			log.Fatalf("No se pudo leer el archivo init.sql: %v", err)
		}

		_, err = DB.Exec(string(initSQL))
		if err != nil {
			log.Fatalf("Error al ejecutar init.sql: %v", err)
		}

		log.Println("Archivo init.sql ejecutado correctamente")
	}
}

func isDBEmpty() bool {
	var count int
	row := DB.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'")
	if err := row.Scan(&count); err != nil {
		log.Fatalf("Error al verificar tablas en la base de datos: %v", err)
	}
	return count == 0
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
