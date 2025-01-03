package models

type Task struct {
	TaskId      int    `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type APIKey struct {
	KeyID  int    `json:"key_id"`
	Key    string `json:"key"`
	UserID int    `json:"user_id"`
}

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}
