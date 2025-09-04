package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/cmd/database"
	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/env"
	// "github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/services"
)

func NewServer() *http.Server {

	NewApplication := &Application{
		port: env.GetEnvInt("PORT", 8080),
		db:   database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewApplication.port),
		Handler:      NewApplication.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

// func (app *application) serve() error {
// 	server := http.Server{
// 		Addr:         fmt.Sprintf(":%d", app.port),
// 		Handler:      app.routes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Minute,
// 		WriteTimeout: 30 * time.Minute,
// 	}

// 	log.Println("Starting server on", server.Addr)
// 	return server.ListenAndServe()
// }
