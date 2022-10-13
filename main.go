package main

import (
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"winterchen.com/my-site-go/src/global"
	"winterchen.com/my-site-go/src/initialize"
	"winterchen.com/my-site-go/src/middlewares"
)

func main() {
	// init config
	initialize.InitConfig()
	// init logger
	initialize.InitLogger()
	// init router
	Router := initialize.InitRouters()
	// init db
	initialize.InitDB()
	// init minio
	initialize.InitMinIO()
	// init translation
	initialize.InitTrans("zh")

	// init cache
	initialize.InitCache()

	color.Cyan("patient-go is running")

	// set logger middleware
	Router.Use(middlewares.LoggerForGin())

	// use cors
	Router.Use(middlewares.Cors())

	// start server
	err := Router.Run(fmt.Sprintf(":%d", global.Configs.Port))
	if err != nil {
		global.Log.Info("the patient-go ", zap.String("error", "run failed"))
	}

}
