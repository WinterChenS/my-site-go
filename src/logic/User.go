package logic

import (
	"crypto/md5"
	"encoding/hex"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"winterchen.com/my-site-go/src/dao"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/models"
	"winterchen.com/my-site-go/src/requests"
	"winterchen.com/my-site-go/src/responses"
)

func UpdateUser(user *models.User) error {
	if user.Uid == 0 {
		global.Log.Error("UpdateUser: user.Uid is nil")
		panic("UpdateUser: user.Uid is nil")
	}
	return dao.UpdateUser(user)
}

func GetUserById(id int) (*models.User, error) {
	return dao.GetById(id)
}

func Login(c *gin.Context) {
	var login requests.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		global.Log.Error("Login: username or password is nil")
		responses.Error(c, 400, 400, "username or password is nil", nil)
		return
	}
	pwd := md5V(login.Username + login.Password)
	user, err := dao.GetByUserNameAndPassword(login.Username, pwd)
	if err != nil {
		responses.Error(c, 400, 400, "username or password is wrong", nil)
		return
	}
	token := uuid.New().String()
	userJson, err := json.Marshal(user)
	if err != nil {
		global.Log.Error("Login: json.Marshal(user) error")
		responses.Error(c, 500, 500, "system error", nil)
		return
	}
	global.Cache.Set([]byte(token), userJson, 86400)
	c.Header("Authorization", token)
	responses.Success(c, 200, "success", user)
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
