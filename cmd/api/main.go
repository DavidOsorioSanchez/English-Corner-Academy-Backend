package main

import (
	"log"

	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/cmd/database"
	"github.com/DavidOsorioSanchez/englishcorneracademy-gim/internal/services"

	_ "github.com/DavidOsorioSanchez/englishcorneracademy-gim/docs"

	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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

type Application struct {
	port      int
	jwtSecret string
	db        database.Service
	models    services.Models
}

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	done := make(chan bool, 1)

	server := NewServer()

	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server: ", err)
		return
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
