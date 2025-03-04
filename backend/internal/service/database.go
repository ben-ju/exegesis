package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	MAX_CONNECTIONS      = 25
	MAX_IDLE_CONNECTIONS = 25
	MAX_CONNECTION_TIME  = 5 * time.Minute
)

func InitDB() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error while initializing database : %v", err)
	}

	db.SetMaxOpenConns(MAX_CONNECTIONS)
	db.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	db.SetConnMaxLifetime(MAX_CONNECTION_TIME)

	if err := db.Ping(); err != nil {
		log.Fatalf("Error while trying to ping the database : %v", err)
	}
	log.Println("[INFO] CONNECTED TO DATABASE")
	return db
}
