package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"sales-user-srv/handler"
	"sales-user-srv/proto"
)

func main() {

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	err = server.Serve(listen)

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	cfg := api.DefaultConfig()
	client, err := api.NewClient(cfg)

	check := &api.AgentServiceCheck{
		GRPC:                           "127.0.0.1:8082",
		Timeout:                        "5s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = "sales-user-srv"
	registration.ID = "1"
	registration.Check = check
	registration.Port = 8000
	registration.Tags = []string{"sales-srv-usr"}

	fmt.Println("HELLO")

	if err != nil {
		zap.S().Error("注册服务失败")
		panic(err)
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
	zap.S().Info("服务器运行在:8000端口上")
}
