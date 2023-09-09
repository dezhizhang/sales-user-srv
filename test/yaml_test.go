package test

import (
	"sales-user-srv/model"
	"testing"
)

type ServerConfig struct {
	ServiceName string             `mapstructure:"name"`
	MySqlInfo   model.ServerConfig `mapstructure:"mysql"`
}

func InitConfig() {
	//v := viper.New()
	//v.SetConfigFile("../config.yaml")
	//err := v.ReadInConfig()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("--user", v.Get("user"))
	//server := &ServerConfig{}
	//fmt.Println(server)
	//err = v.Unmarshal(&server)
	//if err != nil {
	//	panic(err)
	//}
	//
	//v.WatchConfig()
	//v.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println("config.yaml发生变化")
	//	v.ReadInConfig()
	//	v.Unmarshal(&server)
	//	fmt.Println(server.MySqlInfo.Host)
	//})
	//
	//fmt.Println("--", server.MySqlInfo.Host)

}

func TestYaml(t *testing.T) {
	InitConfig()
}
