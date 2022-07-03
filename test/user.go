package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"user_srv/proto"
	"user_srv/utils"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.Password, user.NickName, user.Role)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			Name:     fmt.Sprintf("刘德华%d", i),
			NickName: fmt.Sprintf("阿黄%d", i),
			Role:     1,
			Gender:   1,
			Mobile:   fmt.Sprintf("1599247844%d", i),
			Password: utils.Md5(fmt.Sprintf("123456%d", i)),
		})
	}

}

func TestGetUserById() {
	rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: 22,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name, rsp.NickName, rsp.Mobile, rsp.Gender, rsp.Role)
}

func main() {
	Init()
	defer conn.Close()
	//TestGetUserList()
	//TestCreateUser()
	TestGetUserById()
}
