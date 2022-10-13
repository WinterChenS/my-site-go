package initialize

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"winterchen.com/my-site-go/src/config"
	"winterchen.com/my-site-go/src/global"
)

// init config
func InitConfig() {
	v := viper.New()
	//read env config
	v.SetConfigFile("env.yml")
	if err := v.ReadInConfig(); err != nil {
		color.Red("read env.yml error: %v", err)
		panic(err)
	}
	envConfig := config.EnvConfig{}
	if err := v.Unmarshal(&envConfig); err != nil {
		color.Red("can not unmarshal envConfig, default env is dev")
		envConfig.Env = "dev"
	}

	//read config
	v.SetConfigFile("setting-" + envConfig.Env + ".yml")
	if err := v.ReadInConfig(); err != nil {
		color.Red("read setting-"+envConfig.Env+".yml error: %v", err)
		panic(err)
	}
	configs := config.Config{}
	if err := v.Unmarshal(&configs); err != nil {
		color.Red("can not unmarshal configs")
		panic(err)
	}
	global.Configs = configs
}
