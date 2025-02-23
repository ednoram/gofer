package handlers

import (
	"database/sql"
	"gofer/db"
	"gofer/db/sqlc"
	"gofer/schemas"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var req schemas.CreateTask
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

	task, err := db.GetQueries().CreateTask(c, sqlc.CreateTaskParams{
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: true},
		Completed:   sql.NullBool{Bool: req.Completed, Valid: true},
		CreatedBy:   userId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
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
	c.JSON(http.StatusCreated, taskResponse)
}

func GetTasks(c *gin.Context) {
	completedParam := c.Query("completed")
	validCompletionValues := map[string]bool{"1": true, "0": false}

	var completed sql.NullBool
	if value, exists := validCompletionValues[completedParam]; exists {
		completed = sql.NullBool{Bool: value, Valid: true}
	} else {
		completed = sql.NullBool{Valid: false}
	}

	tasks, err := db.GetQueries().ListTasks(c, completed)
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

	var req schemas.UpdateTask
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTime := time.Now().UTC()

	updateParams := sqlc.UpdateTaskParams{
		TaskID:    id,
		UpdatedAt: sql.NullTime{Valid: true, Time: updateTime},
	}
	if req.Title != nil {
		updateParams.Title = sql.NullString{Valid: true, String: *req.Title}
	}
	if req.Description != nil {
		updateParams.Description = sql.NullString{Valid: true, String: *req.Description}
	}
	if req.Completed != nil {
		updateParams.Completed = sql.NullBool{Valid: true, Bool: *req.Completed}
	}

	err = db.GetQueries().UpdateTask(c, updateParams)
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
