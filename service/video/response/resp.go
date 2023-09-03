package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

// 封装基本报文格式
func newBaseResp(err *errno.Errno) *VideoService.BaseResp {
	if err.HTTP == 200 || err.HTTP == 0 {
		return &VideoService.BaseResp{
			StatusCode: 0,
			StatusMsg:  err.Message,
		}
	} else {
		return &VideoService.BaseResp{
			StatusCode: err.HTTP,
			StatusMsg:  err.Message,
		}
	}
}

func BuildBaseResp(err error) *VideoService.BaseResp {
	var baseResp *VideoService.BaseResp
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

func BuildPublishActionResp(err error) *VideoService.VideoPublishActionResponse {
	var baseResp *VideoService.BaseResp
	baseResp = BuildBaseResp(err)
	return &VideoService.VideoPublishActionResponse{
		BaseResp: baseResp,
	}
}

func BuildVideoFeedResp(videoList []*VideoService.Video, nextTime int64, err error) *VideoService.VideoFeedResponse {
	var baseResp *VideoService.BaseResp
	baseResp = BuildBaseResp(err)
	return &VideoService.VideoFeedResponse{
		BaseResp:  baseResp,
		VideoList: videoList,
		NextTime:  nextTime,
	}
}

func BuildVideoListResp(videoList []*VideoService.Video,
	err error) *VideoService.VideoPublishListResponse {
	var baseResp *VideoService.BaseResp
	baseResp = BuildBaseResp(err)
	return &VideoService.VideoPublishListResponse{
		BaseResp:  baseResp,
		VideoList: videoList,
	}
}

func BuildVideoInfoUpdateResp(err error) *VideoService.UpdateVideoInfoResponse {
	var baseResp *VideoService.BaseResp
	baseResp = BuildBaseResp(err)
	return &VideoService.UpdateVideoInfoResponse{
		BaseResp: baseResp,
	}
}

func BuildGetAuthorIdByVideoIdResponse(authorId int64, err error) *VideoService.GetAuthorIdByVideoIdResponse {
	var baseResp *VideoService.BaseResp
	baseResp = BuildBaseResp(err)
	return &VideoService.GetAuthorIdByVideoIdResponse{
		BaseResp: baseResp,
		AuthorId: authorId,
	}
}
