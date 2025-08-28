package main

import (
	"database/sql"
	"log"

	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/database"

	_ "github.com/DavidOsorioSanchez/englishcorneracademy-gim/docs"

	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/env"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// esto es de lo mas loco que e visto jajajajaj,
// para documentar con swagger tienes que comentar los comandos. Por ejemplo

// @title English Corner Academy API
// @version 1.0
// @description API para la gestión de eventos y usuarios en English Corner Academy
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Token de acceso para las operaciones protegidas **Bearer &lt;token&gt;**

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
