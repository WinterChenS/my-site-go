package dao

import (
	"github.com/jinzhu/gorm"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
)

func AddRelationShip(relationShip *models.RelationShip) error {
	return global.DB.Create(relationShip).Error
}

func DeleteRelationShipById(cid, mid int) error {
	return global.DB.Where("cid = ? AND mid = ?", cid, mid).Delete(&models.RelationShip{}).Error
}

func DeleteRelationShipByCid(cid int, tx *gorm.DB) error {
	err := global.DB.Where("cid = ?", cid).Delete(&models.RelationShip{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func DeleteRelationShipByMid(mid int) error {
	return global.DB.Where("mid = ?", mid).Delete(&models.RelationShip{}).Error
}

func UpdateRelationShipByCid(cid, mid int) error {
	return global.DB.Model(&models.RelationShip{}).Where("cid = ?", cid).Update("mid", mid).Error
}

func GetRelationShipByCid(cid int) ([]*models.RelationShip, error) {
	var relationShips []*models.RelationShip
	err := global.DB.Where("cid = ?", cid).Find(&relationShips).Error
	return relationShips, err
}

func GetRelationShipByMid(mid int) ([]*models.RelationShip, error) {
	var relationShips []*models.RelationShip
	err := global.DB.Where("mid = ?", mid).Find(&relationShips).Error
	return relationShips, err
}

func GetRelationShipCountById(cid, mid int) (int64, error) {
	var count int64
	err := global.DB.Model(&models.RelationShip{}).Where("cid = ? AND mid = ?", cid, mid).Count(&count).Error
	return count, err
}
