package main

import (
	"English-Corner-Academy-Gim/internal/database"
	"English-Corner-Academy-Gim/internal/env"
	"database/sql"
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "mysecret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
