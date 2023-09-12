package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"sales-user-srv/global"
	"sales-user-srv/handler"
	"sales-user-srv/initialize"
	"sales-user-srv/proto"
)

func main() {
	// 初始化日志
	initialize.Logger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化nacos
	initialize.InitNacos()

	// 初始化数据库
	initialize.InitDB()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		global.ServerConfig.UserSrvConfig.Host, global.ServerConfig.UserSrvConfig.Port,
	))
	if err != nil {
		zap.S().Errorw("监听失败%s", err.Error())
	}
	//fmt.Println(global.ServerConfig)
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port,
	)
	client, err1 := api.NewClient(cfg)
	if err1 != nil {
		panic(err1)
	}

	check := &api.AgentServiceCheck{
		GRPC:                           "127.0.0.1:8082",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = global.ServerConfig.UserSrvConfig.Port
	registration.Tags = []string{"sales-user-srv", "xiaozhi"}
	registration.Address = "127.0.0.1:8082"
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		zap.S().Errorw("启动服务失败%s", err.Error())
	}
	zap.S().Infof("服务运行在端口:%d", global.ServerConfig.UserSrvConfig.Port)
}
