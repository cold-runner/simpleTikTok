package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// 封装基本报文格式
func newBaseResp(err *errno.Errno) *BaseResponse.BaseResp {
	if err.HTTP == 200 || err.HTTP == 0 {
		return &BaseResponse.BaseResp{
			StatusCode: 0,
			StatusMsg:  err.Message,
		}
	} else {
		return &BaseResponse.BaseResp{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
		}
	}
}

func newUserRegisterResponse(baseResp *BaseResponse.BaseResp, userId int64) *UserService.UserRegisterResponse {
	return &UserService.UserRegisterResponse{
		BaseResp: baseResp,
		Id:       userId,
	}
}

func newUserLoginResponse(baseResp *BaseResponse.BaseResp,
	userId int64) *UserService.UserLoginResponse {
	return &UserService.UserLoginResponse{
		BaseResp: baseResp,
		Id:       userId,
	}
}

func newUserInfoResponse(baseResp *BaseResponse.BaseResp,
	user *UserService.User) *UserService.UserInfoResponse {
	return &UserService.UserInfoResponse{
		BaseResp: baseResp,
		User:     user,
	}
}

func newMUserInfoResponse(baseResp *BaseResponse.BaseResp,
	userInfos []*UserService.User) *UserService.MGetUserInfoResponse {
	return &UserService.MGetUserInfoResponse{
		BaseResp: baseResp,
		Users:    userInfos,
	}
}

func BuildUserRegisterResp(err error, userId int64) *UserService.UserRegisterResponse {
	var baseResp *BaseResponse.BaseResp
	// 如果没有错误，返回成功
	baseResp = BuildBaseResp(err)

	return newUserRegisterResponse(baseResp, userId)
}

func BuildUserLoginResp(err error, userId int64) *UserService.UserLoginResponse {
	var baseResp *BaseResponse.BaseResp
	baseResp = BuildBaseResp(err)

	return newUserLoginResponse(baseResp, userId)
}

func BuildUserInfoResp(err error, user *UserService.User) *UserService.UserInfoResponse {
	var baseResp *BaseResponse.BaseResp
	baseResp = BuildBaseResp(err)
	return newUserInfoResponse(baseResp, user)
}

func BuildBaseResp(err error) *BaseResponse.BaseResp {
	var baseResp *BaseResponse.BaseResp
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
