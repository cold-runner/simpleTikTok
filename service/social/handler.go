package main

import (
	"context"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/response"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
	"github.com/cold-runner/simpleTikTok/service/social/service"
	"time"
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
		uid := req.GetUserId()
		info, err := rpc.UserClient.GetUserInfo(ctx, &UserService.UserInfoRequest{
			FromId: uid,
			ToId:   uid,
		})
		user := info.GetUser()
		if err != nil {
			log.Errorw("get user info failed", "err", err)
			return nil, err
		}
		var commentUid int64
		if commentModel == nil {
			commentUid = 0
		} else {
			commentUid = commentModel.CommentId
		}
		// 处理上传好的 ID 问题
		comment := &SocialService.Comment{
			// id 为虚假 id 如果没有查到
			Id: commentUid,
			User: &SocialService.UserInfo{
				Id:              user.GetId(),
				Name:            user.GetName(),
				FollowCount:     user.GetFollowCount(),
				FollowerCount:   user.GetFollowerCount(),
				IsFollow:        user.GetIsFollow(),
				Avatar:          user.GetAvatar(),
				BackgroundImage: user.GetBackgroundImage(),
				Signature:       user.GetSignature(),
				TotalFavorited:  user.GetTotalFavorited(),
				WorkCount:       user.GetWorkCount(),
				FavoriteCount:   user.GetFavoriteCount(),
			},
			Content:    req.GetCommentText(),
			CreateDate: time.Now().Format("01-02"),
		}
		resp = response.BuildCommentActionResp(comment, errno.OK)
	} else if actionType == 2 {
		resp = response.BuildCommentActionResp(nil, errno.OK)
	}

	return resp, nil
}

// CommentList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) CommentList(ctx context.Context, req *SocialService.CommentListRequest) (resp *SocialService.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
