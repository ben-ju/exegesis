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

type Book struct {
	Title              string `json:"title"`
	Abbreviation       string `json:"abbreviation"`
	IsDeuterocanonical int    `json:"is_deuterocanonical"`
	IsOldTestament     int    `json:"is_old_testament"`
	IsNewTestament     int    `json:"is_new_testament"`
}

// SeedBooks est la liste de livres à insérer initialement.
var SeedBooks = []Book{
	// Ancien Testament
	{Title: "Genesis", Abbreviation: "Gen", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Exodus", Abbreviation: "Exod", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Leviticus", Abbreviation: "Lev", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Numbers", Abbreviation: "Num", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Deuteronomy", Abbreviation: "Deut", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Joshua", Abbreviation: "Josh", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Judges", Abbreviation: "Judg", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ruth", Abbreviation: "Ruth", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Samuel", Abbreviation: "1 Sam", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Samuel", Abbreviation: "2 Sam", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Kings", Abbreviation: "1 Kgs", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Kings", Abbreviation: "2 Kgs", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Chronicles", Abbreviation: "1 Chr", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Chronicles", Abbreviation: "2 Chr", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ezra", Abbreviation: "Ezra", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Nehemiah", Abbreviation: "Neh", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Tobit", Abbreviation: "Tobit", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Judith", Abbreviation: "Judith", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Esther", Abbreviation: "Est", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "1 Maccabees", Abbreviation: "1 Macc", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "2 Maccabees", Abbreviation: "2 Macc", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Job", Abbreviation: "Job", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Psalms", Abbreviation: "Ps", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Proverbs", Abbreviation: "Prov", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ecclesiastes", Abbreviation: "Eccl", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Songs of Songs", Abbreviation: "Song", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Wisdom of Solomon", Abbreviation: "Wisd", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Sirach", Abbreviation: "Sir", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Isaiah", Abbreviation: "Isa", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Jeremiah", Abbreviation: "Jer", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Lamentations", Abbreviation: "Lam", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Baruch", Abbreviation: "Baru", IsDeuterocanonical: 1, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Ezekiel", Abbreviation: "Ezek", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Daniel", Abbreviation: "Dan", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Hosea", Abbreviation: "Hos", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Joel", Abbreviation: "Joel", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Amos", Abbreviation: "Amos", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Obadiah", Abbreviation: "Obad", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Jonah", Abbreviation: "Jonah", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Micah", Abbreviation: "Mic", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Nahum", Abbreviation: "Nah", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Habakkuk", Abbreviation: "Hab", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Zephaniah", Abbreviation: "Zeph", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Haggai", Abbreviation: "Hag", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Zechariah", Abbreviation: "Zech", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},
	{Title: "Malachi", Abbreviation: "Mal", IsDeuterocanonical: 0, IsOldTestament: 1, IsNewTestament: 0},

	// Nouveau Testament
	{Title: "Matthew", Abbreviation: "Matt", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Mark", Abbreviation: "Mark", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Luke", Abbreviation: "Luke", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "John", Abbreviation: "John", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Acts", Abbreviation: "Acts", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Romans", Abbreviation: "Rom", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Corinthians", Abbreviation: "1 Cor", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Corinthians", Abbreviation: "2 Cor", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Galatians", Abbreviation: "Gal", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Ephesians", Abbreviation: "Eph", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Philippians", Abbreviation: "Phil", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Colossians", Abbreviation: "Col", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Thessalonians", Abbreviation: "1 Thess", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Thessalonians", Abbreviation: "2 Thess", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Timothy", Abbreviation: "1 Tim", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Timothy", Abbreviation: "2 Tim", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Titus", Abbreviation: "Titus", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Philemon", Abbreviation: "Philem", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Hebrews", Abbreviation: "Heb", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "James", Abbreviation: "Jas", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 Peter", Abbreviation: "1 Pet", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 Peter", Abbreviation: "2 Pet", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "1 John", Abbreviation: "1 Jn", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "2 John", Abbreviation: "2 Jn", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "3 John", Abbreviation: "3 Jn", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Jude", Abbreviation: "Jude", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
	{Title: "Revelation", Abbreviation: "Rev", IsDeuterocanonical: 0, IsOldTestament: 0, IsNewTestament: 1},
}

// InitDB initialise la connexion à la base de données en utilisant les variables d'environnement.
func initDb(config config.Config) (*sql.DB, error) {
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
		return nil, err
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

	if err := setupDatabase(db); err != nil {
		log.Printf("ERROR WHILE SETUP DATABASE %v\n", err)
	}
	return db, nil
}

// SetupDatabase crée les tables et index si nécessaire.
func setupDatabase(db *sql.DB) error {
	log.Println("IN SETUP DATABASE")
	schema := `
	CREATE TABLE IF NOT EXISTS category (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		abbreviation TEXT
	);

	CREATE TABLE IF NOT EXISTS bible_books (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		abbreviation TEXT NOT NULL,
		is_deuterocanonical BOOLEAN DEFAULT FALSE,
		is_old_testament BOOLEAN DEFAULT FALSE,
		is_new_testament BOOLEAN DEFAULT FALSE
	);

	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		abbreviation TEXT,
		language TEXT,
		authors TEXT,
		cover TEXT,
		category_id INTEGER REFERENCES category(id)
	);

	CREATE TABLE IF NOT EXISTS chapters (
		id SERIAL PRIMARY KEY,
		bible_book_id INTEGER NOT NULL REFERENCES bible_books(id),
		number INTEGER NOT NULL,
		is_ambiguous BOOLEAN DEFAULT FALSE
	);

	CREATE TABLE IF NOT EXISTS verses (
		id SERIAL PRIMARY KEY,
		chapter_id INTEGER NOT NULL REFERENCES chapters(id),
		number INTEGER NOT NULL,
		is_ambiguous BOOLEAN DEFAULT FALSE
	);

	CREATE TABLE IF NOT EXISTS contents (
		id SERIAL PRIMARY KEY,
		book_id INTEGER NOT NULL REFERENCES books(id),
		start_verse_id INTEGER REFERENCES verses(id),
		end_verse_id INTEGER REFERENCES verses(id),
		text TEXT
	);
`
	if _, err := db.Exec(schema); err != nil {
		return err
	}

	log.Println("Tables created successfully.")
	// if err := seedBibleBooks(db); err != nil {
	// 	return err
	// }
	//
	// log.Println("Seed done!")
	return nil
}

func seedBibleBooks(db *sql.DB) error {
	log.Println("IN SEEDER")
	for _, book := range SeedBooks {
		query := `
			INSERT INTO bible_books (title, abbreviation, is_deuterocanonical, is_old_testament, is_new_testament)
			SELECT $1, $2, $3, $4, $5
			WHERE NOT EXISTS (
				SELECT 1 FROM bible_books WHERE title = $1 AND abbreviation = $2
			);
`

		_, err := db.Exec(query, book.Title, book.Abbreviation, book.IsDeuterocanonical, book.IsOldTestament, book.IsNewTestament)
		if err != nil {
			return err
		}
	}
	log.Println("Seed data inserted successfully into bible_books.")
	return nil
}
