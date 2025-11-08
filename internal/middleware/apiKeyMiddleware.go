package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuthMiddlware(expectedKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" || apiKey != expectedKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized: invalid api key",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
