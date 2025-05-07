package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null;type:text" json:"content"`
	UserID  uint   `json:"userID"`
	User    User
	PostID  uint `json:"postID"`
	Post    Post
}

func AddComment(comment *Comment) error {
	return DB.Create(comment).Error
}

func ListComment(postId uint) ([]Comment, error) {
	var comments []Comment
	err := DB.Where("post_id = ?", postId).Find(&comments).Error
	return comments, err
}
