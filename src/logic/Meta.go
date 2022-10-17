package logic

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"winterchen.com/my-site-go/src/dao"
	"winterchen.com/my-site-go/src/enums"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
	"winterchen.com/my-site-go/src/responses"
)

func AddMeta(c *gin.Context) {
	var meta models.Meta
	if err := c.ShouldBindJSON(&meta); err != nil {
		responses.Error(c, 400, 400, "invalid params", nil)
	}
	if meta == (models.Meta{}) {
		responses.Error(c, 400, 400, "invalid params", nil)
	}
	dao.AddMeta(&meta)
	responses.Success(c, 200, "success", nil)
}

func SaveMeta(metaType, name string, mid int) error {
	if metaType != "" && name != "" {
		metas, err := dao.GetMetasByIfNameOrType(metaType, name)
		if err != nil {
			return err
		}
		if len(metas) == 0 {
			meta := &models.Meta{
				Type: metaType,
				Name: name,
			}
			if mid != 0 {
				meta, err := dao.GetMetaById(mid)
				if meta != nil {
					meta.Mid = mid
				}
				err = dao.UpdateMeta(meta)
				if err != nil {
					return err
				}
				//更新原有的文章分类 todo
				if meta != nil {
					UpdateCategory(meta.Name, name)
				}
			} else {
				err := dao.AddMeta(meta)
				if err != nil {
					return err
				}
			}
		}

	} else {
		panic("metaType or name is null")
	}
	return nil
}

func AddMetas(cid int, names, contentType string, tx *gorm.DB) error {
	if cid == 0 {
		panic("cid is null")
	}
	if names != "" && contentType != "" {
		names := strings.Split(names, ",")
		for _, name := range names {
			SaveOrUpdate(cid, name, contentType, tx)
		}
	}
	return nil
}

func SaveOrUpdate(cid int, name, contentType string, tx *gorm.DB) error {
	metas, err := GetMetas(name, contentType)
	if err != nil {
		return err
	}
	var mid int
	if len(metas) == 1 {
		meta := metas[0]
		mid = meta.Mid
	} else if len(metas) > 1 {
		tx.Rollback()
		panic("metas is more than one")
	} else {
		meta := &models.Meta{
			Slug: name,
			Name: name,
			Type: contentType,
		}
		err := dao.AddMeta(meta)
		if err != nil {
			tx.Rollback()
			return err
		}
		mid = meta.Mid
	}
	if mid != 0 {
		count, _ := dao.GetRelationShipCountById(cid, mid)
		if count == 0 {
			relationShip := &models.RelationShip{
				Cid: cid,
				Mid: mid,
			}
			err := dao.AddRelationShip(relationShip)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return nil
}

func DeleteMetaById(mid int) error {
	if mid == 0 {
		panic("mid is null")
	}
	tx := global.DB.Begin()
	meta, err := dao.GetMetaById(mid)
	if err != nil {
		return err
	}
	if meta != nil {
		//需要把相关的数据删除
		relationShips, err := dao.GetRelationShipByMid(mid)
		if err != nil {
			return err
		}
		if relationShips != nil {
			for _, relationShip := range relationShips {
				content, err := GetContentById(relationShip.Cid)
				if err != nil {
					return err
				}
				if content != nil {
					contentTmp := &models.Content{
						Cid: content.Cid,
					}
					if meta.Type == enums.CATEGORY.String() {
						metaName, err := ReMeta(meta.Name, content.Categories)
						if err != nil {
							return err
						}
						contentTmp.Categories = metaName
					} else if meta.Type == enums.TAG.String() {
						metaName, err := ReMeta(meta.Name, content.Tags)
						if err != nil {
							return err
						}
						contentTmp.Tags = metaName
					}
					//将删除的资源去除

				}
			}
		}
		dao.DeleteMetaById(mid)
	}
	tx.Commit()
}

func GetMetas(name, contentType string) ([]*models.Meta, error) {
	return dao.GetMetasByIfNameOrType(name, contentType)
}

func ReMeta(name, metas string) (string, error) {
	ms := strings.Split(metas, ",")
	var newMetas string
	for _, m := range ms {
		if m == name {
			continue
		}
		newMetas += m + ","
	}
	if newMetas != "" {
		newMetas = newMetas[:len(newMetas)-1]
	}
	return newMetas, nil
}
