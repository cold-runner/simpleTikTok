package v1

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/dao/model"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/pack"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
	"github.com/marmotedu/errors"
	"time"
)

type CommentActionService struct {
	ctx context.Context
	dal.Dal
	rpc.UserRpc
}

func NewCommentActionService(dbIns dal.Dal, rpcIns rpc.UserRpc) *CommentActionService {
	return &CommentActionService{
		Dal:     dbIns,
		UserRpc: rpcIns,
	}
}

// CommentAction 创建/删除操作
func (s *CommentActionService) CommentAction(
	ctx context.Context,
	request *SocialService.CommentActionRequest) (*SocialService.Comment, error) {
	// TODO 硬编码
	if request.ActionType == 0 {
		if err := s.Dal.DeleteComment(ctx, request.GetCommentId()); err != nil {
			return nil, errors.WithCode(errno.ErrDatabase, "删除评论失败，err: %v", err)
		}
		return nil, nil
	}

	comment := &model.Comment{
		UserID:    request.GetUserId(),
		VedioID:   request.GetVideoId(),
		Content:   request.GetCommentText(),
		CreatedAt: time.Now(),
	}

	if err := s.Dal.CreateComment(s.ctx, comment); err != nil {
		return nil, errors.WithCode(errno.ErrDatabase, "创建评论失败，err: %v", err)
	}
	// TODO RPC调用用户服务
	user, err := s.UserRpc.GetUserInfo(request.GetUserId())
	if err != nil {
		return nil, errors.WithCode(errno.ErrRpcCall, "rpc请求用户服务失败，err: %v", err)
	}

	// 添加评论成功后拼接comment返回
	respComment := pack.CommentInfoConvert(user, comment)
	return respComment, nil
}

func (s *CommentActionService) CommentList(
	ctx context.Context,
	request *SocialService.CommentListRequest) ([]*SocialService.Comment, error) {

	commentList, err := s.Dal.CommentList(ctx, request.GetVideoId())
	if err != nil {
		return nil, errors.WithCode(errno.ErrDatabase, "获取评论列表失败，err: %v", err)
	}

	respCommentList := make([]*SocialService.Comment, len(commentList))
	for i := range commentList {
		// RPC调用获取用户信息
		// TODO 循环发起远程调用，优化
		user, err := s.UserRpc.GetUserInfo(commentList[i].UserID)
		if err != nil {
			return nil, errors.WithCode(errno.ErrRpcCall, "rpc请求用户服务失败，err: %v", err)
		}
		respCommentList[i] = &SocialService.Comment{
			Id:         commentList[i].ID,
			User:       user,
			Content:    commentList[i].Content,
			CreateDate: commentList[i].CreatedAt.String(),
		}
	}
	return respCommentList, nil
}

func (s *CommentActionService) LikeAction(
	ctx context.Context,
	request *SocialService.LikeActionRequest) error {
	// TODO 硬编码
	if request.GetActionType() == 1 {
		// 点赞
		vl := &model.VideoLike{
			UserID:  request.GetUserId(),
			VideoID: request.GetVideoId(),
		}
		if err := s.Dal.LikeAction(ctx, vl); err != nil {
			return errors.WithCode(errno.ErrDatabase, "点赞失败，err: %v", err)
		}
		return nil
	}

	// 取消点赞
	if err := s.Dal.CancelLikeAction(ctx, request.GetUserId(), request.GetVideoId()); err != nil {
		return errors.WithCode(errno.ErrDatabase, "取消点赞失败，err: %v", err)
	}
	return nil
}

func (s *CommentActionService) LikeList(
	ctx context.Context,
	request *SocialService.LikeListRequest) ([]*SocialService.Video, error) {

	list, err := s.Dal.LikeList(ctx, request.GetUserId())
	if err != nil {
		return nil, errors.WithCode(errno.ErrDatabase, "获取喜欢列表失败，err: %v", err)
	}

	respVideos := make([]*SocialService.Video, len(list))
	for i := range respVideos {
		// TODO 循环RPC调用
		user, err := s.UserRpc.GetUserInfo(request.GetUserId())
		if err != nil {
			return nil, errors.WithCode(errno.ErrRpcCall, "rpc请求用户服务失败，err: %v", err)
		}
		respVideos[i] = &SocialService.Video{
			Id:            list[i].ID,
			Author:        user,
			PlayUrl:       list[i].PlayURL,
			CoverUrl:      list[i].CoverURL,
			FavoriteCount: list[i].FavoriteCount,
			CommentCount:  list[i].CommentCount,
			IsLike:        true,
			Title:         list[i].Title,
		}
	}
	return respVideos, nil
}
