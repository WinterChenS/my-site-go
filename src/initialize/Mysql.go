package initialize

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"winterchen.com/my-site-go/src/global"
)

func InitDB() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", global.Configs.Mysql.Username, global.Configs.Mysql.Password, global.Configs.Mysql.Host, global.Configs.Mysql.Port, global.Configs.Mysql.Dbname))
	if err != nil {
		global.Log.Error("mysql connect error", zap.Any("err", err))
		panic(err)
	}
	global.DB = db
}
