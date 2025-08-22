package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: up or down")
	}
	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")

	if err != nil {
		log.Fatalf("Error opening migration files: %v", err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)

	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error running migration: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error running migration: %v", err)
		}
	default:
		log.Fatal("Invalid migration direction. Please use 'up' or 'down'.")
	}
}
