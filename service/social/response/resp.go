package response

import (
	"errors"
	"github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
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

func BuildFavoriteActionResp(err error) *SocialService.FavoriteActionResponse {
	baseResp := BuildBaseResp(err)
	return &SocialService.FavoriteActionResponse{
		BaseResp: baseResp,
	}
}

func BuildFavoriteListResp(videoList []*VideoService.Video, err error) *SocialService.
	FavoriteListResponse {
	baseResp := BuildBaseResp(err)
	return &SocialService.FavoriteListResponse{
		BaseResp:  baseResp,
		VideoList: videoList,
	}
}

func BuildCommentActionResp(comment *SocialService.Comment, err error) *SocialService.
	CommentActionResposne {
	//log.Debugw("BuildCommentActionResp", "comment", comment, "err", err)
	baseResp := BuildBaseResp(err)
	if comment == nil {
		return &SocialService.CommentActionResposne{
			BaseResp: baseResp,
			Comment:  nil,
		}
	} else {
		return &SocialService.CommentActionResposne{
			BaseResp: baseResp,
			Comment:  comment,
		}
	}
}

func BuildCommentListResp(commentList []*SocialService.Comment,
	err error) *SocialService.CommentListResponse {
	baseResp := BuildBaseResp(err)
	return &SocialService.CommentListResponse{
		BaseResp:    baseResp,
		CommentList: commentList,
	}
}

func BuildFavoriteIdList(vidList []int64, err error) *SocialService.
	GetFavoriteVideoByUidResponse {
	if err == nil {
		return &SocialService.GetFavoriteVideoByUidResponse{
			FavoriteIdList: vidList,
		}
	} else {
		return &SocialService.GetFavoriteVideoByUidResponse{
			FavoriteIdList: nil,
		}
	}
}
