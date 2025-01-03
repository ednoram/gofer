package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"gofer/db"
	"gofer/utils"
)

func generateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	apiKey := hex.EncodeToString(bytes)
	return apiKey, nil
}

func storeAPIKey(apiKey string, userId int) error {
	hashedKey := utils.HashAPIKey(apiKey)

	_, err := db.GetDbConn().Exec("INSERT INTO api_key (api_key, user_id) VALUES (?, ?)", hashedKey, userId)
	if err != nil {
		return fmt.Errorf("Failed to insert API key: %v", err)
	}
	return nil
}

func main() {
	var userId int = 1 // Change this to the user ID you want to generate API key for

	db.InitializeDatabase()
	defer db.GetDbConn().Close()

	apiKey, err := generateAPIKey()
	if err != nil {
		log.Fatalf("Error generating API key: %v", err)
	}

	err = storeAPIKey(apiKey, userId)
	if err != nil {
		log.Fatalf("Error storing API key: %v", err)
	}

	fmt.Printf("Generated API Key: %s\n", apiKey)
}
