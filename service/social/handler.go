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

// FavoriteList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FavoriteList(ctx context.Context, req *SocialService.FavoriteListRequest) (resp *SocialService.FavoriteListResponse, err error) {
	// TODO: Your code here...
	uid, toUid := req.GetUserId(), req.GetToUserId()
	if uid <= 0 || toUid <= 0 {
		log.Errorw("SocialServiceImpl.FavoriteList", "err", errno.ErrInvalidParameter)
	}
	videoList, err := service.NewFavoriteListService(ctx).GetFavoriteList(ctx, req)
	if err != nil {
		resp = response.BuildFavoriteListResp(nil, err)
		return resp, nil
	}
	resp = response.BuildFavoriteListResp(videoList, errno.OK)
	return resp, nil
}

// CommentAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentAction(ctx context.Context, req *SocialService.CommentActionRequest) (resp *SocialService.CommentActionResposne, err error) {
	// TODO: Your code here..
	actionType := req.GetActionType()

	if actionType <= 0 || actionType > 2 {
		log.Errorw("SocialServiceImpl.CommentAction", "err", errno.ErrInvalidParameter)
	}

	if actionType == 1 && (req.GetCommentText() == "" || req.GetVideoId() <= 0 || req.GetUserId() <= 0) {
		log.Errorw("SocialServiceImpl.CommentAction", "err", errno.ErrInvalidParameter)
	}

	if actionType == 2 && req.GetCommentId() <= 0 {
		log.Errorw("SocialServiceImpl.CommentAction", "err", errno.ErrInvalidParameter)
	}
	commentModel, err := service.NewCommentService(ctx).CommentAction(ctx, req)
	if err != nil {
		log.Errorw("comment action request failed", "err", err)
		return nil, err
	}
	log.Debugw("comment action request success")
	if actionType == 1 {
		comment, err := service.NewCommentService(ctx).GetUserLatestComment(
			ctx, req.GetUserId(), commentModel, req.GetCommentText())
		if err != nil {
			log.Errorw("get user latest comment failed", "err", err)
			resp = response.BuildCommentActionResp(nil, err)
		} else {
			resp = response.BuildCommentActionResp(comment, errno.OK)
		}
	} else if actionType == 2 {
		resp = response.BuildCommentActionResp(nil, errno.OK)
	}

	return resp, nil
}

// CommentList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentList(ctx context.Context, req *SocialService.CommentListRequest) (resp *SocialService.CommentListResponse, err error) {
	// TODO: Your code here...
	if req.GetVideoId() <= 0 || req.GetUserId() <= 0 {
		log.Errorw("SocialServiceImpl.CommentList", "err", errno.ErrInvalidParameter)
	}
	commentList, err := service.NewCommentListService(ctx).GetCommentList(ctx, req)
	if err != nil {
		resp = response.BuildCommentListResp(nil, err)
		return nil, err
	} else {
		resp = response.BuildCommentListResp(commentList, errno.OK)
		return resp, nil
	}
}

// GetFavoriteVidList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) GetFavoriteVidList(ctx context.Context, req *SocialService.GetFavoriteVideoByUidRequest) (resp *SocialService.GetFavoriteVideoByUidResponse, err error) {
	// TODO: Your code here...
	if req.GetUserId() <= 0 {
		log.Errorw("SocialServiceImpl.GetFavoriteVidList", "err", errno.ErrInvalidParameter)
	}
	vidList, err := service.NewFavoriteVIDService(ctx).GetFavoriteIDsByUserID(ctx, req)
	if err != nil {
		resp = response.BuildFavoriteIdList(nil, nil)
		return resp, err
	} else {
		resp = response.BuildFavoriteIdList(vidList, err)
		return resp, nil
	}
}
