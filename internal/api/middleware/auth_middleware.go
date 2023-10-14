package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session")
		if err != nil {
			c.SetCookie("session", uuid.New().String(), 60*60*24*400, "/", "localhost", false, true)
		}
		c.Next()
	}
}