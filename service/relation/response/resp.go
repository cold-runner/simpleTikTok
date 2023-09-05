package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
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

func BuildBaseResp(err error) *BaseResponse.BaseResp {
	var baseResp *BaseResponse.BaseResp
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

func BuildRelationActionResp(err error) *RelationService.RelationActionResponse {
	var baseResp *BaseResponse.BaseResp
	baseResp = BuildBaseResp(err)
	return &RelationService.RelationActionResponse{
		BaseResp: baseResp,
	}
}

func BuildQueryActionResp(isFollow bool, err error) *RelationService.
	RelationQueryResponse {
	baseResp := BuildBaseResp(err)
	return &RelationService.RelationQueryResponse{
		BaseResp: baseResp,
		IsFollow: isFollow,
	}
}

func BuildGetUserFollowListResp(userInfos []*UserService.User,
	err error) *RelationService.RelationFollowListResponse {
	baseResp := BuildBaseResp(err)
	return &RelationService.RelationFollowListResponse{
		BaseResp:   baseResp,
		FollowList: userInfos,
	}
}

func BuildGetUserFollowerListResp(userInfos []*UserService.User,
	err error) *RelationService.RelationFollowerListResponse {
	baseResp := BuildBaseResp(err)
	return &RelationService.RelationFollowerListResponse{
		BaseResp:     baseResp,
		FollowerList: userInfos,
	}
}

func BuildGetUserInfosWithRelationResp(userInfos []*UserService.User, err error) *RelationService.QueryUserInfosWithRelationResponse {
	return &RelationService.QueryUserInfosWithRelationResponse{
		UserInfoList: userInfos,
	}
}
