package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"sales-user-srv/global"
	"sales-user-srv/handler"
	"sales-user-srv/initialize"
	"sales-user-srv/proto"
)

func main() {
	// 初始化配置文件
	initialize.InitConfig()
	srvConfig := global.ServerConfig.UserSrvConfig

	// 初始化数据库
	initialize.InitDB()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", srvConfig.Host, srvConfig.Port))
	if err != nil {
		log.Printf("监听失败%s", err.Error())
	}

	err = server.Serve(listen)
	if err != nil {
		log.Printf("启动服务失败%s", err.Error())
	}
	zap.S().Infof("服务运行在端口:%d", srvConfig.Port)
}
