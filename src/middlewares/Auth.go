package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
	"winterchen.com/my-site-go/src/responses"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		user, err := global.Cache.Get([]byte(authHeader))
		if err != nil || user == nil {
			responses.Error(context, http.StatusUnauthorized, 401, "Unauthorized", nil)
			context.Abort()
			return
		}
		var mUser models.User
		err = json.Unmarshal(user, &mUser)
		if err != nil {
			responses.Error(context, http.StatusUnauthorized, 401, "Unauthorized", nil)
			context.Abort()
			return
		}
		global.Log.Info(fmt.Sprintf("Auth success %v", mUser.Username))
		context.Set("user", mUser)
		context.Next()
	}
}
