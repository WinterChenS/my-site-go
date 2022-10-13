package dao

import (
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
)

func UpdateUser(user *models.User) error {
	return global.DB.Model(&user).Updates(
		map[string]interface{}{
			"screenName": user.ScreenName,
			"email":      user.Email,
		},
	).Error
}

func GetById(id int) (*models.User, error) {
	var user models.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetByUserNameAndPassword(username, password string) (*models.User, error) {
	var user models.User
	err := global.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	// 判断&user是否为空
	if user.Uid == 0 {
		return nil, err
	}

	return &user, err
}
