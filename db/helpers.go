package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"gofer/db/sqlc"
)

var database *sql.DB

var dbPath string = "gofer.db"

func InitializeDatabase() {
	log.Println("Initializing database")

	var err error
	database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Test the connection
	if err := database.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
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
