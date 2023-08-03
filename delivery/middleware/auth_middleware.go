package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizaationHeader string `header:"authorization"`
}

func AuthMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "api/v1/login" {
			c.Next()
			fmt.Println("login")
		} else {
			var h authHeader
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}

			if h.AuthorizaationHeader != "token123" {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}
	}
}
