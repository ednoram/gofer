package middleware

import (
	// "gofer/models"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// apiKey := c.GetHeader("X-API-Key")
		// if apiKey == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "API key is required"})
		// 	c.Abort()
		// 	return
		// }

		// Validate api key here

		// // Add user to context
		// c.Set("userID", key.UserID)

		c.Next()
	}
}
