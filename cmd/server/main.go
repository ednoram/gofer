package main

import (
	"gofer/db"
	"gofer/handlers"
	"gofer/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server")
	db.InitializeDatabase()

	r := gin.Default()

	// Middleware
	r.Use(middleware.APIKeyAuthMiddleware())

	// Routes
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
