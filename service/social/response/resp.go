package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// 封装基本报文格式
func newBaseResp(err *errno.Errno) *SocialService.BaseResp {
	if err.HTTP == 200 {
		return &SocialService.BaseResp{
			StatusCode: 0,
			StatusMsg:  err.Message,
		}
	} else {
		return &SocialService.BaseResp{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
		}
	}
}

func BuildBaseResp(err error) *SocialService.BaseResp {
	var baseResp *SocialService.BaseResp
	// 如果没有错误，返回成功
	if err == nil {
		baseResp = newBaseResp(errno.OK)
	} else {
		// 如果有错误，判断是否是自定义的错误
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

func BuildFavoriteActionResp(err error) *SocialService.FavoriteActionResponse {
	baseResp := BuildBaseResp(err)
	return &SocialService.FavoriteActionResponse{
		BaseResp: baseResp,
	}
}
