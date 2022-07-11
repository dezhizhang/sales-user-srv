package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"user_srv/handler"
	"user_srv/proto"
)

func main() {

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
	fmt.Println("服务运行在：8000")
}
