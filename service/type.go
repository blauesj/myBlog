package service

import "myBlog/models"

type RegisterRequest struct {
	UserName string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string
	User  *models.User
}

type CommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostId  uint   `json:"postId" binding:"required"`
}
