package repository

import (
	"database/sql"
)

type IBookRepository interface {
	ListVersions() ([]string, error)
}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) ListVersions() string {
	return "version"
	// rows, err := r.db.Query("SELECT version_name FROM bible_versions")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	//
	// var versions []string
	// for rows.Next() {
	// 	var version string
	// 	if err := rows.Scan(&version); err != nil {
	// 		return nil, err
	// 	}
	// 	versions = append(versions, version)
	// }
	// return versions, nil
}
