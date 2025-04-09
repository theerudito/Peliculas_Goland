package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite3", "./peliculas.db")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	// Testear conexión
	if err = DB.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("Conectado a la base de datos SQLite")

	// Ejecutar el archivo init.sql
	initSQL, err := ioutil.ReadFile("init.sql")
	if err != nil {
		log.Fatalf("No se pudo leer el archivo init.sql: %v", err)
	}

	_, err = DB.Exec(string(initSQL))
	if err != nil {
		log.Fatalf("Error al ejecutar init.sql: %v", err)
	}

	log.Println("Archivo init.sql ejecutado correctamente")

}
