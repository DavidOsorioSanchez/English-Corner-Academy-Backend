package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/golang-migrate/migrate/database/mysql"
)

var (
	dbname     = os.Getenv("DB_DATABASENAME")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)

type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}

	err = dbInstance.RunMigrations()

	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}

// este codigo esta temporal hasta que lo refactorize bien
// ya que asi si me funciona

func (s *service) RunMigrations() error {
	log.Println("Iniciando migraciones de la base de datos...")

	errUser := s.CreateUsersTable()
	if errUser != nil {
		return fmt.Errorf("error creando tabla users: %v", errUser)
	}

	errEvents := s.CreateEventsTable()
	if errEvents != nil {
		return fmt.Errorf("error creando tabla events: %v", errEvents)
	}

	errAttendees := s.CreateAttendeesTable()
	if errAttendees != nil {
		return fmt.Errorf("error creando tabla attendees: %v", errAttendees)
	}

	log.Println("Todas las migraciones completadas exitosamente")
	return nil
}

func (s *service) CreateUsersTable() error {
	log.Println("entrando en el script de crear tabla users")
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	) `

	tabla, err := s.db.Exec(query)

	if tabla != nil {
		fmt.Println("tabla creada")
	}

	if err != nil {
		log.Fatalf("error creando tabla users: %v", err)
		return err
	}

	log.Println("Tabla 'users' creada exitosamente")
	return nil
}

func (s *service) CreateEventsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		owner_id INT NOT NULL,  
		name VARCHAR(100) NOT NULL,
		description VARCHAR(255) NOT NULL,
		date DATETIME NOT NULL,
		location VARCHAR(255) NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
	) `

	tabla, err := s.db.Exec(query)

	if tabla != nil {
		fmt.Println("tabla creada")
	}

	if err != nil {
		log.Fatalf("error creando tabla events: %v", err)
		return err
	}

	log.Println("Tabla 'events' creada exitosamente")
	return nil
}
func (s *service) CreateAttendeesTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS attendees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		event_id INT NOT NULL,
		user_id INT NOT NULL,
		FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	) `

	tabla, err := s.db.Exec(query)

	if tabla != nil {
		fmt.Println("tabla creada")
	}

	if err != nil {
		log.Fatalf("error creando tabla attendees: %v", err)
		return err
	}

	log.Println("Tabla 'attendees' creada exitosamente")
	return nil
}
