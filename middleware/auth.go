package middleware

import (
	"gofer/db"
	"gofer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyHeader := c.GetHeader("x-api-key")
		if apiKeyHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		hashedKey := utils.HashAPIKey(apiKeyHeader)

		apiKey, err := db.GetQueries().GetApiKey(c, hashedKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userId", apiKey.UserID)

		c.Next()
	}
}
