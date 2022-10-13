package router

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/controllers"
	"winterchen.com/my-site-go/src/middlewares"
)

/**
need auth
*/
func AdminRouter(Router *gin.RouterGroup) {
	HomeRouter := Router.Group("/home").Use(middlewares.LoggerForGin()).Use(middlewares.Cors()).Use(middlewares.Auth())
	{
		HomeRouter.GET("/about", controllers.GetAbout)
	}
}
