package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	JwtKey                     = "myBlog"
	TokenExpireDuration        = time.Now().Add(time.Second * 60 * 60 * 24 * 7).Unix()
	RefreshTokenExpireDuration = time.Now().Add(time.Second * 60 * 60 * 24 * 14).Unix()
)

type UserClaims struct {
	Id       uint
	Username string
	jwt.StandardClaims
}
