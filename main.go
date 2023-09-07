package main

import (
	"fmt"
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

	host := global.ServerConfig.SalesUserSrvConfig.Host
	port := global.ServerConfig.SalesUserSrvConfig.Port

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Printf("监听失败%s", err.Error())
	}

	err = server.Serve(listen)
	if err != nil {
		log.Printf("启动服务失败%s", err.Error())
	}
}
