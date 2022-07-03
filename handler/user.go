package handler

import (
	"context"
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
		Id:       user.Id,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Name:     user.Name,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}

	if user.Birthday != nil {
		userInfoResponse.Birthday = uint64(user.Birthday.Unix())
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
	result := driver.DB.Where("mobile?=", req.Mobile).Find(&user)
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
