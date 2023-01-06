package logic

import (
	"strings"

	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/dao"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
	"winterchen.com/my-site-go/src/requests"
	"winterchen.com/my-site-go/src/responses"
)

func addContent(c *gin.Context) {
	global.Log.Info("addContent")
	var content models.Content
	if err := c.ShouldBindJSON(&content); err != nil {
		global.Log.Error("invalid params")
		responses.Error(c, 400, 400, "invalid params", nil)
	}
	if err := dao.AddContent(&content); err != nil {
		global.Log.Error("add content failed")
		responses.Error(c, 500, 500, "add content failed", nil)
	}
	// Todo 标签和分类的设置

	responses.Success(c, 200, "add content success", nil)
}

func UpdateCategory(ordinal, newCatefory string) {
	contents, _ := dao.GetContentsByType(ordinal)
	if len(contents) > 0 {
		for _, content := range contents {
			content.Categories = strings.Replace(content.Categories, ordinal, newCatefory, -1)
			dao.UpdateContentCategoryById(content.Cid, content.Categories)
		}
	}
}

func UpdateContentById(content *models.Content) error {
	tx := global.DB.Begin()
	if err := dao.UpdateContentById(content, tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := dao.DeleteRelationShipByCid(content.Cid, tx); err != nil {
		tx.Rollback()
		return err
	}
	AddMetas(content.Cid, content.Tags, "tag", tx)
	AddMetas(content.Cid, content.Categories, "category", tx)
	tx.Commit()
	return nil
}

func GetContentById(cid int) (*models.Content, error) {
	return dao.GetContentById(cid)
}

func GetContentsByCond(request *requests.ContentSearch, pageNum int, pageSize int) (*responses.ContentResponse, error) {

}
