package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

var dbPath string = "gofer.db"

func InitializeDatabase() {
	log.Println("Initializing database")

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	fmt.Println("Successfully connected to the database!")
}

func GetConn() *sql.DB {
	if db == nil {
		InitializeDatabase()
	}

	return db
}
