package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sales-user-srv/global"
	"sales-user-srv/model"
	"sales-user-srv/proto"
	"sales-user-srv/utils"
)

type UserServer struct {
	//proto.UnimplementedGreeterServer
	proto.UnimplementedUserServer
}

//

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResponse := proto.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		Name:     user.Name,
		Gender:   user.Gender,
		Nickname: user.NickName,
		Role:     int32(user.Role),
	}

	return userInfoResponse
}

// GetUserList 反回用户列表信息
func (u *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	//分页
	global.DB.Scopes(utils.Paginate(int(req.PageIndex), int(req.PageSize))).Find(&users)

	for _, user := range users {
		userInfoResp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoResp)
	}
	return rsp, nil
}

// GetUserByMobile 通过mobile查询用户
func (u *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	tx := global.DB.Where(&model.User{Mobile: req.Mobile}).Find(&user)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// GetUserById 通过id查询用户
func (u *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	tx := global.DB.Where("id =? ", req.Id).Find(&user)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// CreateUser 新建用户
func (u *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	tx := global.DB.Where(&model.User{Mobile: req.Mobile}).Find(&user)
	if tx.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Name = req.Name
	user.Mobile = req.Mobile
	user.Password = utils.Md5(req.Password)

	tx = global.DB.Create(&user)
	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, tx.Error.Error())
	}
	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

// UpdateUser 更新用户
func (u *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	tx := global.DB.Where("id = ?", req.Id).Find(&user)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user.Name = req.Name
	user.NickName = req.Name
	user.Gender = req.Gender
	tx = global.DB.Save(&user)
	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, tx.Error.Error())
	}
	return &empty.Empty{}, nil
}

// DeleteUser 删除用户
func (u *UserServer) DeleteUser(ctx context.Context, req *proto.IdRequest) (*empty.Empty, error) {
	var user model.User
	tx := global.DB.Where("id = ?", req.Id).Find(&user)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	tx = global.DB.Delete(&user)
	if tx.Error != nil {
		return nil, status.Errorf(codes.Internal, tx.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (u *UserServer) CheckPassword(ctx context.Context, req *proto.CheckInfo) (*proto.CheckResponse, error) {
	return nil, nil
}
