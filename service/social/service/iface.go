package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
)

type Srv interface {
	CommentAction(context.Context, *SocialService.CommentActionRequest) (*SocialService.Comment, error)
	CommentList(context.Context, *SocialService.CommentListRequest) ([]*SocialService.Comment, error)
	LikeAction(context.Context, *SocialService.LikeActionRequest) error
	LikeList(context.Context, *SocialService.LikeListRequest) ([]*SocialService.Video, error)
}
