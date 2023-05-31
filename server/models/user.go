package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// 其他字段
}

func (User) TableName() string {
	return "my_users"
}
