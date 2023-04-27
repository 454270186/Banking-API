package initial

import (
	"RESTful/config"
	"RESTful/global"

	"github.com/spf13/viper"
)

func InitialConfig() {
	v := viper.New()
	v.SetConfigFile("./setting-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := config.ServerConfig{}

	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	global.Settings = serverConfig
}