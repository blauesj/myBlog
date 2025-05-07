package models

import (
	"errors"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null;type:varchar(50)" json:"title" binding:"required"`
	Content string `gorm:"not null;type:text" json:"content" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
	User    User
}

func AddPost(post *Post) error {
	//验证用户存在
	_, err := GetUser(post.UserID)
	if err != nil {
		return errors.New("user is not variable")
	}

	err = DB.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPosts() ([]Post, error) {
	var posts []Post
	err := DB.Preload("User").Find(&posts).Error
	return posts, err
}

func GetPost(id uint) (*Post, error) {
	var post Post
	err := DB.Preload("User").First(&post, id).Error
	return &post, err
}

func UpdatePost(id uint, post *Post) error {
	return DB.Model(&post).Where("id = ?", id).Updates(post).Error
}

func DeletePost(id uint) error {
	return DB.Delete(&Post{}, id).Error
}
