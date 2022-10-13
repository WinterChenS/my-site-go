package router

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/controllers"
	"winterchen.com/my-site-go/src/middlewares"
)

func HomeRouter(Router *gin.RouterGroup) {
	HomeRouter := Router.Group("/home").Use(middlewares.LoggerForGin()).Use(middlewares.Cors())
	{
		HomeRouter.GET("/about", controllers.GetAbout)
		HomeRouter.POST("/login", controllers.Login)
	}
}
