package main

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
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
