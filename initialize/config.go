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
	debug := GetEnvInfo(SALES_USER_DRV)

	configFilePrefix := "config"

	configFileName := fmt.Sprintf("sales-user-srv/%s-pro.yaml", configFilePrefix)

	if debug {
		configFileName = fmt.Sprintf("sales-user-srv/%s-debug.yaml", configFileName)
	}
	v := viper.New()
	viper.SetConfigName(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

}
