package initialize

//func InitConsul(server *grpc.Server) {
//	serverConfig := global.ServerConfig
//	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
//
//	grpcConfig := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)
//	cfg := api.DefaultConfig()
//	client, err := api.NewClient(cfg)
//
//	check := &api.AgentServiceCheck{
//		GRPC:                           grpcConfig,
//		Timeout:                        "5s",
//		Interval:                       "10s",
//		DeregisterCriticalServiceAfter: "10s",
//	}
//
//	registration := new(api.AgentServiceRegistration)
//	registration.Name = serverConfig.Name
//	registration.ID = serverConfig.Name
//	registration.Check = check
//	registration.Port = serverConfig.Port
//	registration.Tags = []string{"sales-srv-usr"}
//
//	if err != nil {
//		panic(err)
//	}
//
//	err = client.Agent().ServiceRegister(registration)
//	if err != nil {
//		panic(err)
//	}
//
//}
