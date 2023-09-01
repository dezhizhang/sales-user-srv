package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"sales-user-srv/handler"
	"sales-user-srv/proto"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	listen, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Printf("监听失败%s", err.Error())
	}

	err = server.Serve(listen)
	if err != nil {
		log.Printf("启动服务失败%s", err.Error())
	}
	//
	//tx := driver.DB.Create(&model.User{ID: "123", Name: "刘德华12", Mobile: "15992478448", Password: "123456"})
	//if tx.RowsAffected == 0 {
	//	log.Printf("创建用户失败%s", tx.Error.Error())
	//	return
	//}
	//fmt.Println("创建用户成功")
}
