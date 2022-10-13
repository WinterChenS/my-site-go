package logic

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/dao"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
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
