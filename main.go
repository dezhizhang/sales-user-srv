package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sales-user-srv/handler"
	"sales-user-srv/initialize"
	"sales-user-srv/proto"
)

func main() {

	// 初始化日志
	initialize.Logger()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
	zap.S().Info("服务器运行在:8000端口上")
}
