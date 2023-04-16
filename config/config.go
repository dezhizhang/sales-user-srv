package config

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"sales-user-srv/model"
)

func InitConfig() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("9cbeaaea-c313-4ff4-b0ce-3d9f00925555"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("tmp/nacos/log"),
		constant.WithCacheDir("tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "sales-user-srv",
		Group:  "dev",
	})

	serverConfig := &model.ServerConfig{}
	err = json.Unmarshal([]byte(content), &serverConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println(serverConfig.Name)

	//err = client.ListenConfig(vo.ConfigParam{
	//	DataId: "sales-usr-srv",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件发生改变")
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//time.Sleep(10 * time.Second)
	//fmt.Println("-------- :" + content)
}