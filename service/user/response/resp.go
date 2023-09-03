package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// 封装基本报文格式
func newBaseResp(err *errno.Errno) *UserService.BaseResp {
	if err.HTTP == 200 || err.HTTP == 0 {
		return &UserService.BaseResp{
			StatusCode: 0,
			StatusMsg:  err.Message,
		}
	} else {
		return &UserService.BaseResp{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
		}
	}
}

func newUserRegisterResponse(baseResp *UserService.BaseResp, userId int64) *UserService.UserRegisterResponse {
	return &UserService.UserRegisterResponse{
		BaseResp: baseResp,
		Id:       userId,
	}
}

func newUserLoginResponse(baseResp *UserService.BaseResp,
	userId int64) *UserService.UserLoginResponse {
	return &UserService.UserLoginResponse{
		BaseResp: baseResp,
		Id:       userId,
	}
}

func newUserInfoResponse(baseResp *UserService.BaseResp,
	user *UserService.User) *UserService.UserInfoResponse {
	return &UserService.UserInfoResponse{
		BaseResp: baseResp,
		User:     user,
	}
}

func newMUserInfoResponse(baseResp *UserService.BaseResp,
	userInfos []*UserService.User) *UserService.MGetUserInfoResponse {
	return &UserService.MGetUserInfoResponse{
		BaseResp: baseResp,
		Users:    userInfos,
	}
}

func BuildUserRegisterResp(err error, userId int64) *UserService.UserRegisterResponse {
	var baseResp *UserService.BaseResp
	// 如果没有错误，返回成功
	baseResp = BuildBaseResp(err)

	return newUserRegisterResponse(baseResp, userId)
}

func BuildUserLoginResp(err error, userId int64) *UserService.UserLoginResponse {
	var baseResp *UserService.BaseResp
	baseResp = BuildBaseResp(err)

	return newUserLoginResponse(baseResp, userId)
}

func BuildUserInfoResp(err error, user *UserService.User) *UserService.UserInfoResponse {
	var baseResp *UserService.BaseResp
	baseResp = BuildBaseResp(err)
	return newUserInfoResponse(baseResp, user)
}

func BuildBaseResp(err error) *UserService.BaseResp {
	var baseResp *UserService.BaseResp
	if err == nil {
		baseResp = newBaseResp(errno.OK)
	} else {
		var e *errno.Errno
		if errors.As(err, &e) {
			baseResp = newBaseResp(e)
		} else {
			unknownErr := errno.Errno{HTTP: 400, Message: err.Error()}
			baseResp = newBaseResp(&unknownErr)
		}
	}

	return baseResp
}

func BuildMGetUserInfoResp(err error, userInfos []*UserService.User) *UserService.
	MGetUserInfoResponse {
	baseResp := BuildBaseResp(err)
	return newMUserInfoResponse(baseResp, userInfos)
}

func BuildUpdateUserInfoResp(err error) *UserService.UpdateUserInfoResponse {
	baseResp := BuildBaseResp(err)
	return &UserService.UpdateUserInfoResponse{
		BaseResp: baseResp,
	}
}
