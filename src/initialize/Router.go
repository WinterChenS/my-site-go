package initialize

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/router"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("/api/v1/")
	router.HomeRouter(ApiGroup)
	return r
}
