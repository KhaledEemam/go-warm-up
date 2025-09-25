package main

import (
	"database/sql"
	"github.com/KhaledEemam/go-warm-up/internal/database"
	"github.com/KhaledEemam/go-warm-up/internal/env"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type application struct {
	port      int
	models    database.Models
	jwtSecret string
}

func main() {

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		port:      env.GetIntEnv("PORT_NUMBER", 8080),
		models:    database.CreateDataBaseModels(db),
		jwtSecret: env.GetStringEnv("JWT_SECRET", "new-secret-jwt"),
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}

}
