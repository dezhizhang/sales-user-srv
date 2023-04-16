package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"sales-user-srv/global"
)

const SALES_USER_DRV = "SALES_USER_DEBUG"

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

// 初始化配置文件

func InitConfig() {

	v := viper.New()
	viper.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println("---------", global.ServerConfig)

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

}
