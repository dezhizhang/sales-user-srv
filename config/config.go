package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sales-user-srv/model"
)

func ReadInConfig() (*model.NacosConfig, error) {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	nacosInfo := new(model.NacosConfig)

	err = v.Unmarshal(&nacosInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println(nacosInfo)
	return nacosInfo, err
}
