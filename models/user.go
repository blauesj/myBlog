package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null;type:varchar(50)"`
	Password string `gorm:"not null;type:varchar(100)"`
	Email    string `gorm:"unique;not null;type:varchar(500)"`
}

func GetUserByUser(userName string) (*User, error) {
	var user User
	err := DB.Where("user_name = ? ", userName).First(&user).Error
	return &user, err

}

func AddNewUser(user *User) error {
	err := DB.Create(user).Error
	if err != nil {
		return errors.New("用户注册失败")
	}
	return nil
}

func GetUser(id uint) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return &user, err
}
