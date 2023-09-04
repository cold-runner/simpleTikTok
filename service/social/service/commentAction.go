package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	dtm "github.com/cold-runner/simpleTikTok/service/dtm/request"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/dal/model"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) CommentAction(ctx context.Context,
	req *SocialService.CommentActionRequest) (*model.VideoComment, error) {
	actionType := req.GetActionType()

	err := dtm.CommentActionRequest(&dtm.CommentRequestToDTM{
		CommentText: req.GetCommentText(),
		VideoId:     req.GetVideoId(),
		UserId:      req.GetUserId(),
		CommentId:   req.GetCommentId(),
		ActionType:  actionType,
	})
	if err != nil {
		log.Errorw("comment action request failed", "err", err)
		return nil, err
	}
	log.Debugw("comment action request success")
	if actionType == 1 {
		comment, err := dal.GetUserLatestComment(ctx, req.GetUserId())
		return comment, err
	}
	if err != nil {
		log.Errorw("get user latest comment failed", "err", err)
		return nil, nil
	}
	return nil, nil
}
