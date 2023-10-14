package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware(dockerMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session")
		if err != nil {
			var domain string
			if dockerMode {
				domain = "larek.itatmisis.ru"
			} else {
				domain = "localhost"
			}
			c.SetCookie("session", uuid.New().String(), 60*60*24*400, "/", domain, false, true)
		}

		c.Next()
	}
}
