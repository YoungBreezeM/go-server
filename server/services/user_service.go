package services

import (
	"server/db"
	"server/models"
)

func GetUserInfo(userID string) (*models.User, error) {
	// 从数据库或其他存储中获取用户信息的逻辑var user models.User
	var user models.User

	if err := db.MYSQL.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserInfo(userID string, user *models.User) error {
	// 更新用户信息的逻辑
	if err := db.MYSQL.Model(&models.User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
