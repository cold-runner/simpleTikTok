package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/dao/model"
)

type Dal interface {
	CreateComment(context.Context, *model.Comment) error
	DeleteComment(context.Context, int64) error
	CommentList(context.Context, int64) ([]*model.Comment, error)
	LikeAction(context.Context, *model.VideoLike) error
	LikeList(context.Context, int64) ([]*model.Video, error)
	CancelLikeAction(context.Context, int64, int64) error
}
