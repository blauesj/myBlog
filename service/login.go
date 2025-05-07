package service

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"myBlog/define"
	"myBlog/helper"
	"myBlog/models"
	"net/http"
)

func Register(c *gin.Context) {
	in := new(RegisterRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "-1",
			"message": err.Error(),
		})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	}
	user := models.User{
		UserName: in.UserName,
		Password: string(password),
	}
	err = models.AddNewUser(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "-1",
			"message": err.Error(),
		})
		return
	}
}

func Login(c *gin.Context) {
	in := new(LoginRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "-1",
			"message": err.Error(),
		})
		return
	}
	//password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	//	return
	//}
	user, err := models.GetUserByUser(in.UserName)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "-1",
			"message": "用户名或密码错误",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		c.JSON(400, gin.H{
			"code":    "-1",
			"message": "用户名或密码错误",
		})
		return
	}
	tokenString, err := helper.GenerateToken(user.ID, user.UserName, define.TokenExpireDuration)
	c.JSON(200, gin.H{
		"code":    "0",
		"message": "登录成功",
		"data": LoginResponse{
			User:  user,
			Token: tokenString,
		},
	})
}
