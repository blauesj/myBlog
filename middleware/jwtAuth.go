package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myBlog/define"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var user define.UserClaims
		_, err := jwt.ParseWithClaims(token, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte(define.JwtKey), nil
		})
		if err != nil {
			c.Abort()
			c.JSON(401, gin.H{
				"code":    401,
				"message": "token验证失败",
			})
		}
		c.Set("user", &user)
		c.Next()
	}
}
