package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/dal"
	"github.com/cold-runner/simpleTikTok/service/relation/mq"
	"github.com/cold-runner/simpleTikTok/service/relation/rpc"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService create new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *RelationService.RelationActionRequest) error {
	var err error
	uid, toUid, actionType := req.Uid, req.ToUserId, req.ActionType

	// 修改数据之前先检查关系是否已经存在
	following, err := dal.QueryFollowing(s.ctx, uid, toUid)
	if err != nil {
		return err
	}
	if following && actionType == 1 {
		return errno.ErrAlreadyFollow
	}
	if !following && actionType == 2 {
		return errno.ErrNotFollow
	}

	// 先尝试修改user模块的粉丝数据库
	_, err = rpc.UpdateUserFollowCount(s.ctx,
		&UserService.ChangeUserFollowCountRequest{
			Id:         uid,
			ToUserId:   toUid,
			ActionType: actionType,
		})
	if err != nil {
		log.Errorw("RelationServer failed to call rpc to update user follow count", "err", err)
		return err
	}
	// 成功后再添加本地数据库中的关系, 使用消息队列异步处理
	request := mq.RelationActionRequest{
		UserId:     uid,
		ToUserId:   toUid,
		ActionType: actionType,
	}
	err = mq.RelationProducer.Publish(request)
	if err != nil {
		log.Errorw("RelationServer failed to publish message", "err", err)
	}
	log.Debugw("RelationServer success to publish message", "request",
		request)

	return nil
}
