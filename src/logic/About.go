package logic

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/responses"
)

func GetAbout(c *gin.Context) {
	global.Log.Info("GetAbout")
	responses.Success(c, 200, "success", "test")
}
