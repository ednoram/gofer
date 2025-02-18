package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gofer/db"
	"gofer/utils"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
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
	// Check if a argument is passed
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <userId>")
	}

	userIdArg := os.Args[1]

	userId, err := strconv.Atoi(userIdArg)
	if err != nil {
		log.Fatalf("Invalid argument userId: %v", err)
	}

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
