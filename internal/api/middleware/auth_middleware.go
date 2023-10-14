package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session")
		if err != nil {
			value := uuid.New().String()
			c.SetCookie("session", value, 60*60*24*400, "/", domain, false, false)
			c.Set("session", value)
		}

		c.Next()
	}
}
