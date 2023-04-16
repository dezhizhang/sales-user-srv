package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"sales-user-srv/global"
	"sales-user-srv/model"
)

func InitNacosConfig(config *model.NacosConfig) {

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.Host, config.Port),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(config.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("tmp/nacos/log"),
		constant.WithCacheDir("tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	//获取nacos配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
	})

	serverConfig := &model.ServerConfig{}
	err = json.Unmarshal([]byte(content), &serverConfig)
	if err != nil {
		panic(err)
	}
	global.ServerConfig = serverConfig
}
