package social

import (
	"context"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/pack"
	"github.com/cold-runner/simpleTikTok/service/social/service"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct {
	ctx context.Context
	srv service.Srv
}

func NewServiceImpl(ctx context.Context, srv service.Srv) *ServiceImpl {
	return &ServiceImpl{
		ctx: ctx,
		srv: srv,
	}
}

// CommentAction implements the ServiceImpl interface.
func (s *ServiceImpl) CommentAction(ctx context.Context, req *SocialService.CommentActionRequest) (*SocialService.CommentActionResponse, error) {
	comment, err := s.srv.CommentAction(ctx, req)
	if err != nil {
		// TODO VLOG?
		log.Errorw("评论操作失败！err: %v", err)
		return pack.BuildCommentActionResp(err, nil), nil
	}
	return pack.BuildCommentActionResp(nil, comment), nil
}

// CommentList implements the ServiceImpl interface.
func (s *ServiceImpl) CommentList(ctx context.Context, req *SocialService.CommentListRequest) (*SocialService.CommentListResponse, error) {
	commentList, err := s.srv.CommentList(ctx, req)
	if err != nil {
		log.Errorw("获取评论列表失败！err: %v", err)
		return pack.BuildCommentListResp(err, commentList), nil
	}
	return pack.BuildCommentListResp(nil, commentList), nil
}

// LikeAction implements the ServiceImpl interface.
func (s *ServiceImpl) LikeAction(ctx context.Context, req *SocialService.LikeActionRequest) (*SocialService.LikeActionResponse, error) {
	err := s.srv.LikeAction(ctx, req)
	if err != nil {
		log.Errorw("点赞/取消点赞失败！err: %v", err)
		return pack.BuildLikeResp(err), nil
	}
	return pack.BuildLikeResp(err), nil
}

// LikeList implements the ServiceImpl interface.
func (s *ServiceImpl) LikeList(ctx context.Context, req *SocialService.LikeListRequest) (*SocialService.LikeListResponse, error) {
	likeList, err := s.srv.LikeList(ctx, req)
	if err != nil {
		log.Errorw("获取点赞列表失败！err: %v", err)
		return pack.BuildLikeListResp(err, nil), nil
	}
	return pack.BuildLikeListResp(err, likeList), nil
}
