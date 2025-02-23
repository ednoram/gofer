package db

import (
	"database/sql"
	"fmt"
	"gofer/db/sqlc"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func InitializeDatabase() {
	log.Println("Initializing database")

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	dbPath := filepath.Join(cwd, "db", "gofer.db")

	database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test the connection
	if err = database.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Enforce foreign keys
	_, err = database.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatalf("Error enabling foreign keys: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}

func GetDbConn() *sql.DB {
	if database == nil {
		InitializeDatabase()
	}

	return database
}

func GetQueries() *sqlc.Queries {
	dbConn := GetDbConn()
	return sqlc.New(dbConn)
}
