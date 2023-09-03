package main

import (
	"context"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/response"
	"github.com/cold-runner/simpleTikTok/service/social/service"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// FavoriteAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FavoriteAction(ctx context.Context, req *SocialService.FavoriteActionRequest) (resp *SocialService.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	uid, toVid, actionType := req.GetUserId(), req.GetToVideoId(), req.GetActionType()
	if uid <= 0 || toVid <= 0 {
		log.Errorw("SocialServiceImpl.FavoriteAction", "err", errno.ErrInvalidParameter)
	}
	if actionType <= 0 || actionType > 2 {
		log.Errorw("SocialServiceImpl.FavoriteAction", "err", errno.ErrInvalidParameter)
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = response.BuildFavoriteActionResp(err)
		return resp, err
	}
	resp = response.BuildFavoriteActionResp(errno.OK)
	return resp, nil
}

// CommentAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentAction(ctx context.Context, req *SocialService.CommentActionRequest) (resp *SocialService.CommentActionResposne, err error) {
	// TODO: Your code here...
	return
}
