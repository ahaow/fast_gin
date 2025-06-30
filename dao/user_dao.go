package dao

import (
	"fast_gin/global"
	"fast_gin/models"
)

// 根据用户名查用户
func GetUserByUsername(username string) (*models.UserModel, error) {
	var user models.UserModel
	err := global.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// CreateUser 创建用户
func CreateUser(user *models.UserModel) error {
	return global.DB.Create(user).Error
}
