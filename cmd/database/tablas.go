package database

// import (
// 	"fmt"
// 	"log"
// )

// // type service struct {
// // 	db *sql.DB
// // }

// func (s *service) RunMigrations() error {
// 	log.Println("Iniciando migraciones de la base de datos...")

// 	if err := s.CreateUsersTable(); err != nil {
// 		return fmt.Errorf("error creando tabla users: %v", err)
// 	}

// 	if err := s.CreateEventsTable(); err != nil {
// 		return fmt.Errorf("error creando tabla events: %v", err)
// 	}

// 	if err := s.CreateAttendeesTable(); err != nil {
// 		return fmt.Errorf("error creando tabla attendees: %v", err)
// 	}

// 	log.Println("Todas las migraciones completadas exitosamente")
// 	return nil
// }

// func (s *service) CreateUsersTable() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INT AUTO_INCREMENT PRIMARY KEY,
// 		name VARCHAR(100) NOT NULL,
// 		email VARCHAR(255) NOT NULL UNIQUE,
// 		password VARCHAR(255) NOT NULL,
// 	) `

// 	_, err := s.db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("error creando tabla users: %v", err)
// 	}

// 	log.Println("Tabla 'users' creada exitosamente")
// 	return nil
// }
// func (s *service) CreateEventsTable() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS events (
// 		id INT AUTO_INCREMENT PRIMARY KEY,
// 		owner_id INT NOT NULL,  
// 		name VARCHAR(100) NOT NULL,
// 		description VARCHAR(255) NOT NULL,
// 		date DATETIME NOT NULL,
// 		location VARCHAR(255) NOT NULL,
// 		FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
// 	) `

// 	_, err := s.db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("error creando tabla eventos: %v", err)
// 	}

// 	log.Println("Tabla 'events' creada exitosamente")
// 	return nil
// }
// func (s *service) CreateAttendeesTable() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS attendees (
// 		id INT AUTO_INCREMENT PRIMARY KEY,
// 		event_id INT NOT NULL,
// 		user_id INT NOT NULL,
// 		FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
// 		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
// 	) `

// 	_, err := s.db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("error creando tabla attendees: %v", err)
// 	}

// 	log.Println("Tabla 'attendees' creada exitosamente")
// 	return nil
// }
