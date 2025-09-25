package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a valid migration direction. ('up'-'down')")
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrationFiles, err := (&file.File{}).Open("cmd/migrate/migrations")

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", migrationFiles, "sqlite3", instance)

	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	default:
		log.Fatal("Invalid direction, Please provide 'up' or 'down'.")

	}

}
