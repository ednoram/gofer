package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"gofer/db"
	"gofer/db/sqlc"
	"gofer/schemas"
)

func CreateTask(c *gin.Context) {
	var req schemas.CreateUpdateTask
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctxUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID required"})
		return
	}
	userId, ok := ctxUserId.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := db.GetQueries().CreateTask(c, sqlc.CreateTaskParams{
		Title:       req.Title,
		Description: sql.NullString{String: req.Description},
		Completed:   sql.NullBool{Bool: req.Completed},
		CreatedBy:   userId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetTasks(c *gin.Context) {
	tasks, err := db.GetQueries().ListTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var taskResponses []schemas.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, schemas.TaskResponse{
			TaskId:      task.TaskID,
			Title:       task.Title,
			Description: task.Description.String,
			Completed:   task.Completed.Bool,
			CreatedBy:   task.CreatedBy,
			CreatedAt:   task.CreatedAt.Time,
			UpdatedAt:   task.UpdatedAt.Time,
		})
	}

	c.JSON(http.StatusOK, taskResponses)
}

func GetTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := db.GetQueries().GetTask(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	taskResponse := schemas.TaskResponse{
		TaskId:      task.TaskID,
		Title:       task.Title,
		Description: task.Description.String,
		Completed:   task.Completed.Bool,
		CreatedBy:   task.CreatedBy,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}

	c.JSON(http.StatusOK, taskResponse)
}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req schemas.CreateUpdateTask
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTime := time.Now().UTC()

	err = db.GetQueries().UpdateTask(c, sqlc.UpdateTaskParams{
		TaskID:      id,
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: true},
		Completed:   sql.NullBool{Bool: req.Completed, Valid: true},
		UpdatedAt:   sql.NullTime{Time: updateTime, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.GetQueries().DeleteTask(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
