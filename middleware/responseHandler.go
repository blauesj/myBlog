package middleware

import "github.com/gin-gonic/gin"

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(400, gin.H{
				"error": c.Errors.ByType(gin.ErrorTypePublic).String(),
			})
		}
	}
}
