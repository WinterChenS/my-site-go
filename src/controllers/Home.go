package controllers

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/logic"
)

func GetAbout(c *gin.Context) {
	logic.GetAbout(c)
}
