package models

import "time"

type Task struct {
	TaskId      int       `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type APIKey struct {
	KeyId  int
	ApiKey string
	UserId int
}

type User struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}
