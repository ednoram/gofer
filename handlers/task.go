package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gofer/db"
	"gofer/models"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User id required"})
		return
	}
	_, err := db.GetConn().Exec(
		"INSERT INTO task (title, description, completed, created_by) VALUES ($1, $2, $3, $4)",
		task.Title, task.Description, task.Completed, userId,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task

	rows, err := db.GetConn().Query("SELECT task_id, title, description, completed, created_by, created_at, updated_at FROM task")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskId, &task.Title, &task.Description, &task.Completed, &task.CreatedBy, &task.CreatedAt, &task.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	row := db.GetConn().QueryRow("SELECT task_id, title, description, completed, created_by, created_at, updated_at FROM task WHERE task_id = ?", id)
	var task models.Task
	err := row.Scan(&task.TaskId, &task.Title, &task.Description, &task.Completed, &task.CreatedBy, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTime := time.Now().UTC().Format("2006-01-02T15:04:05")
	_, err := db.GetConn().Exec(
		"UPDATE task SET title = ?, description = ?, completed = ?, updated_at = ? WHERE task_id = ?",
		task.Title, task.Description, task.Completed, updateTime, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.GetConn().Exec("DELETE FROM task WHERE task_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
