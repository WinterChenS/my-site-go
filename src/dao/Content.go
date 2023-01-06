package dao

import (
	"github.com/jinzhu/gorm"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
	"winterchen.com/my-site-go/src/requests"
)

func AddContent(content *models.Content) error {
	return global.DB.Create(content).Error
}

func GetContentsByType(contentType string) ([]models.Content, error) {
	var contents []models.Content
	err := global.DB.Where("type = ?", contentType).Find(&contents).Error
	return contents, err
}

func UpdateContentCategoryById(id int, category string) error {
	return global.DB.Model(&models.Content{}).Where("id = ?", id).Update("categories", category).Error
}

func UpdateContentById(content *models.Content, tx *gorm.DB) error {
	err := global.DB.Model(&content).Updates(
		map[string]interface{}{
			"title":        content.Title,
			"titlePic":     content.TitlePic,
			"slug":         content.Slug,
			"content":      content.Content,
			"modified":     content.Modified,
			"type":         content.Type,
			"status":       content.Status,
			"categories":   content.Categories,
			"tags":         content.Tags,
			"hits":         content.Hits,
			"commentsNum":  content.CommentsNum,
			"allowComment": content.AllowComment,
			"allowPing":    content.AllowPing,
			"allowFeed":    content.AllowFeed,
		},
	)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	return nil
}

func GetContentById(id int) (*models.Content, error) {
	var content models.Content
	err := global.DB.Where("id = ?", id).First(&content).Error
	return &content, err
}

func GetCount(request *requests.ContentSearch) (int, error) {
	var count int
	err := global.DB.Model(&models.Content{}).Where("type = ?", request.Type).Count(&count).Error
	return count, err
}
