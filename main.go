package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sales-user-srv/config"
	"sales-user-srv/global"
	"sales-user-srv/handler"
	"sales-user-srv/initialize"
	"sales-user-srv/proto"
)

func main() {
	config, err := config.ReadInConfig()
	if err != nil {
		zap.L().Error("读取配置文件失败%s")
	}

	// 初始化nacos配置文件
	initialize.InitNacosConfig(config)

	// 获取全局配置
	serverConfig := global.ServerConfig

	//初始化grpc服务
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port))
	if err != nil {
		panic("监听服务失败" + err.Error())
	}
	err = server.Serve(listen)

	// 注册服务发现
	initialize.InitConsul(server)

	if err != nil {
		panic("grpc启动失败" + err.Error())
	}
	zap.S().Info("服务器运行在:8000端口上")
}
