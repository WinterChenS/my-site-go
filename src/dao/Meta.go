package dao

import (
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
)

func AddMeta(meta *models.Meta) error {
	return global.DB.Create(meta).Error
}
