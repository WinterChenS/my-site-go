package dao

import (
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
)

func AddMeta(meta *models.Meta) error {
	return global.DB.Create(meta).Error
}

func DeleteMetaById(id int) error {
	return global.DB.Where("id = ?", id).Delete(&models.Meta{}).Error
}

func UpdateMeta(meta *models.Meta) error {
	return global.DB.Model(&meta).Updates(
		map[string]interface{}{
			"name":        meta.Name,
			"slug":        meta.Slug,
			"type":        meta.Type,
			"description": meta.Description,
			"sort":        meta.Sort,
			"parent":      meta.Parent,
		},
	).Error
}

func GetMetaById(id int) (*models.Meta, error) {
	var meta models.Meta
	err := global.DB.Where("id = ?", id).First(&meta).Error
	return &meta, err
}

func GetMetasByIfNameOrType(name, metaType string) ([]*models.Meta, error) {
	var metas []*models.Meta
	err := global.DB.Where(&models.Meta{Name: name, Type: metaType}).Find(&metas).Error
	return metas, err
}

func GetCountByType(metaType string) (int, error) {
	var count int
	err := global.DB.Model(&models.Meta{}).Where("type = ?", metaType).Count(&count).Error
	return count, err
}

func GetMetasJoinCount(metaType, order string, limit int) ([]*models.Meta, error) {
	var metas []*models.Meta
	err := global.DB.Model(&models.Meta{}).Select("meta.*, count(relationships.cid) as count").
		Joins("left join relationships on relationships.mid = meta.mid").
		Where("meta.type = ?", metaType).
		Group("meta.id").
		Order("meta." + order).
		Limit(limit).
		Find(&metas).Error
	return metas, err
}
