package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	dtm "github.com/cold-runner/simpleTikTok/service/dtm/request"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/dal/model"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
	"time"
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
	if actionType == 2 {
		authorID, err := dal.GetAuthorByCommentID(ctx, req.GetCommentId())
		if err != nil {
			log.Errorw("get author by comment id failed", "err", err)
			return nil, err
		}
		if authorID != req.UserId {
			log.Errorw("can not delete other's comment", "err", err)
			return nil, errno.ErrInvalidDeleteComment
		}
	}

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

func (s *CommentService) GetUserLatestComment(ctx context.Context,
	userId int64, commentModel *model.VideoComment,
	commentText string) (*SocialService.Comment,
	error) {
	info, err := rpc.UserClient.GetUserInfo(ctx, &UserService.UserInfoRequest{
		FromId: userId,
		ToId:   userId,
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
		Id:         commentUid,
		User:       user,
		Content:    commentText,
		CreateDate: time.Now().Format("01-02"),
	}
	return comment, nil
}
