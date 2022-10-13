package dao

import (
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
)

func AddContent(content *models.Content) error {
	return global.DB.Create(content).Error
}
