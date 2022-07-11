package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user_srv/driver"
	"user_srv/model"
	"user_srv/proto"
	"user_srv/utils"
)

type UserServer struct {
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResponse := proto.UserInfoResponse{
		Id:       uint64(user.Id),
		Mobile:   user.Mobile,
		Password: user.Password,
		Name:     user.Name,
		Birthday: uint64(user.Birthday),
		Gender:   int32(user.Gender),
		Role:     int32(user.Role),
	}

	return userInfoResponse
}

//反回用户列表信息

func (u *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := driver.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	//分页
	driver.DB.Scopes(utils.Paginate(int(req.PageIndex), int(req.PageSize))).Find(&users)

	for _, user := range users {
		userInfoResp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoResp)
	}
	return rsp, nil

}

//通过电话查询用户信息

func (u *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := driver.DB.Where("mobile=?", req.Mobile).Find(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// 通过id查询用记信息

func (u *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := driver.DB.Where("id=?", req.Id).Find(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// 创建用户

func (u *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := driver.DB.Where("mobile?=", req.Mobile).Find(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Id = int64(req.Id)
	user.Name = req.Name
	user.Mobile = req.Mobile
	user.Birthday = int(req.Birthday)
	user.Gender = int(req.Gender)
	user.Password = utils.Md5(req.Password)

	result = driver.DB.Create(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// 更新用户

func (u *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	result := driver.DB.Where("id=?", req.Id).Find(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user.Name = req.Name
	user.Role = int(req.Role)
	user.Birthday = int(req.Birthday)
	user.Gender = int(req.Gender)

	result = driver.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

// 删除用户

func (u *UserServer) DeleteUser(ctx context.Context, req *proto.IdRequest) (*empty.Empty, error) {
	var user model.User
	result := driver.DB.Where("id=?", req.Id).Delete(&user)
	if result.RowsAffected == 1 {
		return &empty.Empty{}, nil
	}
	return nil, result.Error

}
