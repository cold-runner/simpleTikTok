package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

func (s *CommentListService) GetCommentList(ctx context.Context,
	req *SocialService.CommentListRequest) ([]*SocialService.Comment, error) {

	// 1. 从数据库中获取评论列表
	comments, authorIds, err := dal.GetCommentListByVideoId(ctx, req.GetVideoId())
	log.Debugw("get comment list", "comments", comments)
	log.Debugw("get author ids", "authorIds", authorIds)
	if err != nil {
		log.Debugw("get comment list failed", "err", err)
		return nil, err
	}
	// 2. rpc 调用获取作者信息
	userInfos, err := rpc.RelationClient.QueryUserInfosWithRelation(ctx,
		&RelationService.QueryUserInfosWithRelationRequest{
			Uid:    req.GetUserId(),
			IdList: authorIds,
		})
	if err != nil {
		log.Errorw("query user infos with relation failed", "err", err)
		return nil, err
	}
	log.Debugw("get user infos", "userInfos", userInfos)
	// 将作者信息存入map中
	userInfoMap := make(map[int64]*UserService.User)
	for _, userInfo := range userInfos.GetUserInfoList() {
		userInfoMap[userInfo.GetId()] = userInfo
	}
	log.Debugw("get user info map", "userInfoMap", userInfoMap)
	// 3. 将评论列表和作者信息进行组装
	commentList := make([]*SocialService.Comment, 0)
	for _, comment := range comments {
		commentList = append(commentList, &SocialService.Comment{
			Id:         comment.CommentId,
			User:       userInfoMap[comment.AuthorId],
			Content:    comment.Content,
			CreateDate: comment.CreatedAt,
		})
	}
	log.Debugw("get comment list", "commentList", commentList)
	return commentList, nil
}
