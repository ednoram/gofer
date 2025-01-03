package handlers

import (
	"net/http"

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

	_, err := db.GetConn().Exec("INSERT INTO task (title, description, completed) VALUES ($1, $2, $3)", task.Title, task.Description, task.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task

	rows, err := db.GetConn().Query("SELECT task_id, title, description, completed FROM task")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskId, &task.Title, &task.Description, &task.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	row := db.GetConn().QueryRow("SELECT task_id, title, description, completed FROM task WHERE task_id = ?", id)
	var task models.Task
	err := row.Scan(&task.TaskId, &task.Title, &task.Description, &task.Completed)
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

	_, err := db.GetConn().Exec("UPDATE task SET title = ?, description = ?, completed = ? WHERE task_id = ?", task.Title, task.Description, task.Completed, id)
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
