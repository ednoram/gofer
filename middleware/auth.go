package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gofer/db"
	"gofer/models"
	"gofer/utils"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyHeader := c.GetHeader("x-api-key")
		if apiKeyHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		hashed_key := utils.HashAPIKey(apiKeyHeader)

		var apiKey models.APIKey
		row := db.GetConn().QueryRow("SELECT user_id, api_key FROM api_key WHERE api_key = ?", hashed_key)
		err := row.Scan(&apiKey.UserId, &apiKey.ApiKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userId", apiKey.UserId)

		c.Next()
	}
}
