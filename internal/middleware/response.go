package middleware

import "github.com/gin-gonic/gin"

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
