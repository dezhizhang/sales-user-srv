package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sales-user-srv/proto"
	"sales-user-srv/utils"
	"testing"
)

func TestUserList(t *testing.T) {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败%s", err.Error())
	}
	defer conn.Close()

	c := proto.NewUserClient(conn)

	user, err := c.CreateUser(context.Background(), &proto.CreateUserInfo{
		Name:     " 刘德华",
		Password: utils.Md5("admin123"),
		Mobile:   "15992478441",
		Id:       "456",
	})
	if err != nil {
		log.Printf("创建用户失败%s", err.Error())
	}
	fmt.Println(user)

	//user, err := c.GetUserById(context.Background(), &proto.IdRequest{
	//	Id: "123",
	//})
	//if err != nil {
	//	log.Printf("获取用户失败%s", err)
	//}
	fmt.Println(user)
}
