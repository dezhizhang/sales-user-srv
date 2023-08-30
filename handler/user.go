package handler

import (
	"context"
	"sales-user-srv/global"
	"sales-user-srv/model"
	"sales-user-srv/proto"
	"sales-user-srv/utils"
)

type UserServer struct {
}

//GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
//GetUserByMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
//GetUserById(context.Context, *IdRequest) (*UserInfoResponse, error)
//CreateUser(context.Context, *CreateUserInfo) (*UserInfoResponse, error)
//UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
//DeleteUser(context.Context, *IdRequest) (*emptypb.Empty, error)
//CheckPassword(context.Context, *CheckInfo) (*CheckResponse, error)

func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResponse := proto.UserInfoResponse{
		Id:       user.Id,
		Mobile:   user.Mobile,
		Password: user.Password,
		Name:     user.Name,
		Gender:   user.Gender,
		Nickname: user.NickName,
		Role:     int32(user.Role),
	}

	if user.Birthday != nil {
		userInfoResponse.Birthday = uint64(user.Birthday.Unix())
	}
	//return userInfoResponse
	return proto.UserInfoResponse{}
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
