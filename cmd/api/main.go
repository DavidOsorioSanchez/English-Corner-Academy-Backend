package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/services"

	_ "github.com/DavidOsorioSanchez/englishcorneracademy-gim/docs"

	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
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
	models    services.Models
}

var (
	dbname   = os.Getenv("DB_DATABASENAME")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	defer db.Close()

	models := services.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "mysecret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
