package controllers

import (
	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/logic"
)

func Login(c *gin.Context) {
	logic.Login(c)
}
