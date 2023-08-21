package pack

import (
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/marmotedu/errors"
)

// BuildCommentActionResp 报文的封装过程
func BuildCommentActionResp(err error, comment *SocialService.Comment) *SocialService.CommentActionResponse {
	// 解析从底层封装来的错误
	parsedErr := errors.ParseCoder(err)
	resp := &SocialService.CommentActionResponse{
		BaseResp: &SocialService.BaseResp{ //指定响应状态，与errno种返回的错误一致
			StatusCode: int32(parsedErr.Code()),
			StatusMsg:  parsedErr.String(),
		},
		Comment: comment,
	}
	return resp
}

func BuildCommentListResp(err error, commentList []*SocialService.Comment) *SocialService.CommentListResponse {
	// 解析从底层封装来的错误
	parsedErr := errors.ParseCoder(err)
	resp := &SocialService.CommentListResponse{
		BaseResp: &SocialService.BaseResp{
			StatusCode: int32(parsedErr.Code()),
			StatusMsg:  parsedErr.String(),
		},
		CommentList: commentList,
	}
	return resp
}

func BuildLikeListResp(err error, likeList []*SocialService.Video) *SocialService.LikeListResponse {
	// 解析从底层封装来的错误
	parsedErr := errors.ParseCoder(err)
	resp := &SocialService.LikeListResponse{
		BaseResp: &SocialService.BaseResp{
			StatusCode: int32(parsedErr.Code()),
			StatusMsg:  parsedErr.String(),
		},
		VideoList: likeList,
	}
	return resp
}

func BuildLikeResp(err error) *SocialService.LikeActionResponse {
	// 解析从底层封装来的错误
	parsedErr := errors.ParseCoder(err)
	resp := &SocialService.LikeActionResponse{
		BaseResp: &SocialService.BaseResp{
			StatusCode: int32(parsedErr.Code()),
			StatusMsg:  parsedErr.String(),
		},
	}
	return resp
}
