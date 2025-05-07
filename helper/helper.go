package helper

import (
	"github.com/dgrijalva/jwt-go"
	"myBlog/define"
)

func GenerateToken(id uint, name string, expireAt int64) (tokenString string, err error) {
	uc := define.UserClaims{
		Id:       id,
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err = token.SignedString([]byte(define.JwtKey))
	return
}
