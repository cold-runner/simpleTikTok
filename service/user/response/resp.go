package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// 封装基本报文格式
func newBaseResp(err *errno.Errno) *UserService.BaseResp {
	if err.HTTP == 200 {
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

func BuildUserRegisterResp(err error, userId int64) *UserService.UserRegisterResponse {
	var baseResp *UserService.BaseResp
	// 如果没有错误，返回成功
	if err == nil {
		baseResp = newBaseResp(errno.OK)
	} else {
		// 如果有错误，判断是否是自定义的错误
		var e *errno.Errno
		if errors.As(err, &e) {
			baseResp = newBaseResp(e)
		} else {
			// 返回未知错误
			unknownErr := errno.Errno{HTTP: 400, Message: err.Error()}
			baseResp = newBaseResp(&unknownErr)
		}
	}

	return newUserRegisterResponse(baseResp, userId)
}

func BuildUserLoginResp(err error, userId int64) *UserService.UserLoginResponse {
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

	return newUserLoginResponse(baseResp, userId)
}
