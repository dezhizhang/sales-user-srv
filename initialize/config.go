package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sales-user-srv/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//serverConfig := model.SalesServerConfig{}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：%v", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件发生变化：%s", e.Name)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(&global.ServerConfig); err != nil {
			panic(err)
		}
		zap.S().Infof("配置信息：%v", global.ServerConfig)
	})
}
