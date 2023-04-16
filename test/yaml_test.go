package test

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sales-user-srv/model"
	"testing"
)

func InitConfig() {

	v := viper.New()
	v.SetConfigFile("../config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(v.Get("user"))

	nacos := &model.NacosConfig{}

	err = v.Unmarshal(&nacos)
	if err != nil {
		panic(err)
	}
	fmt.Println("------", nacos.Namespace)
	zap.S().Infof("初始化配置文件")

}

func TestYaml(t *testing.T) {
	InitConfig()
}
