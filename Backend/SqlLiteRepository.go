package Forum

import (
	"database/sql"
	"errors"
	// "fmt"
	// "log"
	// "github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("duplicate exists")
	ErrNotExists    = errors.New("not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

// regarde si il y a des erreurs dans la BDD

// function nouveau
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `

    CREATE DATABASE FORUM; -- Crée une BDD --

    USE FORUM; -- Entre dans la base de données --
    
    CREATE TABLE USER (
    id VARCHAR(255) PRIMARY KEY, -- uid tomorrow
    password VARCHAR(255),
    email VARCHAR(255),
    pseudo VARCHAR(255)
    );
    
    CREATE TABLE SUUSER (
    id VARCHAR(255) PRIMARY KEY, --uid tomorrow
    password VARCHAR(255),
    email VARCHAR(255),
    pseudo VARCHAR(255)
    );
    
    CREATE TABLE LOCATIONUSER (
    id VARCHAR(255) PRIMARY KEY, -- uid tomorrow
    country VARCHAR(255)
    );
    
    
`

	_, err := r.db.Exec(query)
	return err
}
