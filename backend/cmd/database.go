package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ben-ju/exegesis/internal/config"
	_ "github.com/lib/pq"
)

const (
	MAX_CONNECTIONS      = 25
	MAX_IDLE_CONNECTIONS = 25
	MAX_CONNECTION_TIME  = 5 * time.Minute
)

// InitDB initialise la connexion à la base de données en utilisant les variables d'environnement.
func initDb(config config.Config) *sql.DB {
	host := config.PsqlHost
	port := config.PsqlPort
	user := config.PsqlUser
	password := config.PsqlPass
	database := config.PsqlDBName

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// On continue même en cas d'erreur, mais il est recommandé de gérer cela en production.
		log.Print("No database connection but we kept running")
	}

	// Configuration du pool de connexions.
	db.SetMaxOpenConns(MAX_CONNECTIONS)
	db.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	db.SetConnMaxLifetime(MAX_CONNECTION_TIME)

	// Vérification de la connexion.
	if err := db.Ping(); err != nil {
		log.Fatalf("Error while trying to ping the database: %v", err)
	}
	log.Println("[INFO] CONNECTED TO DATABASE")

	setupDatabase(db)
	return db
}

// SetupDatabase crée les tables et index si nécessaire.
func setupDatabase(db *sql.DB) error {
    // Used for dev purposes
	// _, err := db.Exec(`
	// 	DROP TABLE IF EXISTS user_notes;
	// 	DROP TABLE IF EXISTS verse_index;
	// 	DROP TABLE IF EXISTS resources;
	// `)
	// if err != nil {
	// 	return fmt.Errorf("error dropping tables: %v", err)
	// }
	// Création des tables
    _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS resources (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			file_path VARCHAR(512) UNIQUE,
			version VARCHAR(50),
			language VARCHAR(3),
			resource_type VARCHAR(20),
			metadata JSONB
		);

		CREATE TABLE IF NOT EXISTS verse_index (
			resource_id INT REFERENCES resources(id),
			osis_id VARCHAR(20),
			byte_start BIGINT,
			byte_length INT,
			chapter INT,
			verse INT,
			content TEXT,
			PRIMARY KEY (resource_id, osis_id)
		);

		CREATE TABLE IF NOT EXISTS user_notes (
			id SERIAL PRIMARY KEY,
			user_id INT,
			resource_id INT REFERENCES resources(id),
			osis_id VARCHAR(20),
			note TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}

	// Création des index
	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_verse_index_osis ON verse_index(osis_id);
		CREATE INDEX IF NOT EXISTS idx_verse_index_resource ON verse_index(resource_id);
		CREATE INDEX IF NOT EXISTS idx_user_notes_user ON user_notes(user_id);
		CREATE INDEX IF NOT EXISTS idx_user_notes_osis ON user_notes(osis_id);
	`)
	if err != nil {
		return fmt.Errorf("error creating indexes: %v", err)
	}

	log.Println("Database setup completed successfully")
	return nil
}
