package main

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"github.com/cold-runner/simpleTikTok/service/user/response"
	"github.com/cold-runner/simpleTikTok/service/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *UserService.UserRegisterRequest) (resp *UserService.UserRegisterResponse, err error) {
	// TODO: Your code here...
	log.Infow("Validating registering parameters...", "req", req)
	if len(req.Password) < 1 || len(req.Username) < 1 || len(req.
		Password) > 32 || len(req.Username) > 32 {
		resp = response.BuildUserRegisterResp(errno.ErrInvalidParameter, -1)
		return resp, err
	}
	log.Infow("Validating parameters success.")

	log.Infow("Creating user...")
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		log.Infow("Create user failed.", "err", err)
		resp = response.BuildUserRegisterResp(err, -1)
		return resp, nil
	}
	log.Infow("User registered success.")
	// 获取创建后的 user 信息
	u, err := service.NewGetUserInfoService(ctx).GetUserInfoByName(req.Username)
	resp = response.BuildUserRegisterResp(errno.OK, u.ID)

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *UserService.UserLoginRequest) (resp *UserService.UserLoginResponse, err error) {
	// TODO: Your code here...
	log.Infow("Validating logging in parameters...", "req", req)
	if len(req.Password) < 1 || len(req.Username) < 1 || len(req.
		Password) > 32 || len(req.Username) > 32 {
		resp = response.BuildUserLoginResp(errno.ErrInvalidParameter, -1)
		return resp, err
	}
	// 验证用户的登录信息，如果成功会返回用户的 id
	uid, err := service.NewValidateLoginService(ctx).ValidateLogin(req)
	if err != nil {
		resp = response.BuildUserLoginResp(err, -1)
		return resp, nil
	}
	resp = response.BuildUserLoginResp(errno.OK, uid)

	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *UserService.UserInfoRequest) (resp *UserService.UserInfoResponse, err error) {
	// TODO: Your code here...
	if req.GetFromId() <= 0 || req.GetToId() <= 0 {
		resp = response.BuildUserInfoResp(errno.ErrInvalidParameter, nil)
		return resp, nil
	}
	uInfo, err := service.NewGetUserInfoService(ctx).GetUserInfoById(req.
		GetFromId(),
		req.GetToId())
	if err != nil {
		resp = response.BuildUserInfoResp(err, nil)
		log.Errorw("get user info failed", "err", err)
		return resp, nil
	}
	// 将数据访问层的 user 结构体转换为 rpc 的 user 结构体
	resp = response.BuildUserInfoResp(errno.OK, convertToRPCUser(uInfo))
	return
}

func convertToRPCUser(uInfo *model.UserInfo) *UserService.User {
	log.Debugw("Convert user info to rpc user info", "uInfo", uInfo)
	return &UserService.User{
		Id:              uInfo.ID,
		Name:            uInfo.Name,
		FollowCount:     uInfo.FollowCount,
		FollowerCount:   uInfo.FollowerCount,
		Avatar:          uInfo.Avatar,
		BackgroundImage: uInfo.BackgroundImage,
		Signature:       uInfo.Signature,
		TotalFavorited:  uInfo.TotalFavorited,
		WorkCount:       uInfo.WorkCount,
		FavoriteCount:   uInfo.FavoriteCount,
		IsFollow:        uInfo.IsFollow,
	}
}

func mConvertToRPCUser(uInfos []*model.UserInfo) []*UserService.User {
	rpcUserInfos := make([]*UserService.User, 0, len(uInfos))
	for _, u := range uInfos {
		k := convertToRPCUser(u)
		log.Debugw("after convert", "uinfo", k)
		rpcUserInfos = append(rpcUserInfos, k)
	}
	log.Debugw("User info with follow info in service level", "userInfos",
		rpcUserInfos)
	return rpcUserInfos
}

// ChangeUserFollowCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserFollowCount(ctx context.Context, req *UserService.ChangeUserFollowCountRequest) (resp *BaseResponse.BaseResp, err error) {
	// TODO: Your code here...

	if req.Id <= 0 || req.ToUserId < 0 || req.ActionType < 1 || req.ActionType > 2 {
		resp = response.BuildBaseResp(errno.ErrInvalidParameter)
		return resp, nil
	}

	// 更新用户的关注数量
	err = service.NewChangeUserFollowService(ctx).UpdateUserFollowInfo(req)
	if err != nil {
		resp = response.BuildBaseResp(err)
		return resp, nil
	}

	resp = response.BuildBaseResp(errno.OK)
	return
}

// MGetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUserInfo(ctx context.Context, req *UserService.MGetUserInfoRequest) (resp *UserService.MGetUserInfoResponse, err error) {
	// TODO: Your code here...
	if len(req.Uids) <= 0 {
		resp = response.BuildMGetUserInfoResp(errno.ErrInvalidParameter, nil)
		return resp, nil
	}
	userInfos, err := service.NewGetUserInfoService(ctx).MGetUserInfoById(
		req.GetUids())
	if err != nil {
		resp = response.BuildMGetUserInfoResp(err, nil)
		return resp, nil
	}
	resp = response.BuildMGetUserInfoResp(errno.OK, mConvertToRPCUser(userInfos))
	return resp, nil
}

// ChangeUserTotalFavoritedCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserTotalFavoritedCount(ctx context.Context, req *UserService.ChangeUserTotalFavoritedCountRequest) (resp *UserService.UpdateUserInfoResponse, err error) {
	// TODO: Your code here...
	if req.GetId() <= 0 {
		resp = response.BuildUpdateUserInfoResp(errno.ErrInvalidParameter)
		return resp, nil
	}
	err = service.NewUpdateUserInfoService(ctx).UpdateUserInfo(req)
	if err != nil {
		return nil, err
	}
	response.BuildUpdateUserInfoResp(errno.OK)
	return
}

// ChangeUserWorkCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserWorkCount(ctx context.Context, req *UserService.ChangeUserWorkCountRequest) (resp *UserService.UpdateUserInfoResponse, err error) {
	// TODO: Your code here...
	if req.GetId() <= 0 {
		resp = response.BuildUpdateUserInfoResp(errno.ErrInvalidParameter)
		return resp, nil
	}
	err = service.NewUpdateUserInfoService(ctx).UpdateUserInfo(req)
	if err != nil {
		return nil, err
	}
	response.BuildUpdateUserInfoResp(errno.OK)
	return
}

// ChangeUserFavoriteCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserFavoriteCount(ctx context.Context, req *UserService.ChangeUserFavoriteCountRequest) (resp *UserService.UpdateUserInfoResponse, err error) {
	// TODO: Your code here...
	if req.GetId() <= 0 {
		resp = response.BuildUpdateUserInfoResp(errno.ErrInvalidParameter)
		return resp, nil
	}
	err = service.NewUpdateUserInfoService(ctx).UpdateUserInfo(req)
	if err != nil {
		return nil, err
	}
	response.BuildUpdateUserInfoResp(errno.OK)
	return
}
