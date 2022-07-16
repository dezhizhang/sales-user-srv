package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user_srv/global"
	"user_srv/model"
)

func Init() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	global.UserSrv = &model.ServerConfig{}
	err = v.Unmarshal(&global.UserSrv)
	if err != nil {
		panic(err)
	}
	zap.S().Infof("初始化配置文件")
}
