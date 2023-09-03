package rpc

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService/userservice"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
)

var userClient userservice.Client

// Register 定义Register函数，它接收上下文对象和用户注册请求，返回用户ID和可能的错误。
func Register(ctx context.Context, req *UserService.UserRegisterRequest) (
	int64, error) {

	log.Debugw("Calling rpc register request:", "req", req)
	resp, err := userClient.Register(ctx, req)
	// 处理调用rpc 服务时可能发生的错误
	if err != nil {
		return -1, err
	}
	// 处理调用结果不成功返回的响应
	if resp.BaseResp.StatusCode != 0 {
		return -1, errno.NewErr(resp.BaseResp.StatusCode,
			"", resp.BaseResp.StatusMsg)
	}
	// 如果注册成功，则返回用户ID和nil错误。
	return resp.Id, nil
}

// Login 定义Login函数，它接收上下文对象和用户登录请求，返回用户ID和可能的错误。
func Login(ctx context.Context, req *UserService.UserLoginRequest) (int64,
	error) {
	log.Debugw("Calling rpc login request:", "req", req)
	resp, err := userClient.Login(ctx, req)
	// 处理调用rpc 服务时可能发生的错误
	if err != nil {
		return -1, err
	}
	// 处理调用结果不成功返回的响应
	if resp.BaseResp.StatusCode != 0 {
		return -1, errno.NewErr(resp.BaseResp.StatusCode,
			"", resp.BaseResp.StatusMsg)
	}
	// 如果登录成功，则返回用户ID和nil错误。
	return resp.Id, nil
}

func GetUserInfo(ctx context.Context, req *UserService.UserInfoRequest) (
	*UserService.UserInfoResponse, error) {
	log.Debugw("Calling rpc get user info request:", "req", req)
	resp, err := userClient.GetUserInfo(ctx, req)
	// 处理调用 rpc 服务时可能发生的错误
	if err != nil {
		return nil, err
	}
	// 处理调用结果不成功返回的响应
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErr(resp.BaseResp.StatusCode,
			"", resp.BaseResp.StatusMsg)
	}
	// 如果登录成功，则返回用户信息和nil错误。
	return resp, nil
}
