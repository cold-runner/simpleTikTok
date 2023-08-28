package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/mq"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService create new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction
func (s *RelationActionService) RelationAction(req *RelationService.RelationActionRequest) error {
	var err error
	uid, toUid, actionType := req.Uid, req.ToUserId, req.ActionType
	request := mq.RelationActionRequest{
		UserId:     uid,
		ToUserId:   toUid,
		ActionType: actionType,
	}
	err = mq.RelationProducer.Publish(request)
	if err != nil {
		log.Errorw("RelationServer failed to publish message", "err", err)
	} else {
		log.Infow("RelationServer success to publish message", "request",
			request)
	}

	return nil
}

//func FollowAction(uid, toUid int64) error {
//	// 判断是否已经关注
//	ctx := context.Background()
//	isFollowing, err := dal.QueryFollowing(ctx, uid, toUid)
//	if err != nil {
//		log.Errorw("RelationActionService.followAction", "err", err)
//		return err
//	}
//	if isFollowing {
//		log.Errorw("RelationActionService.followAction", "err",
//			errno.ErrAlreadyFollow)
//		return errno.ErrAlreadyFollow
//	}
//	err = dal.CreateFollowRelation(ctx, uid, toUid)
//	if err != nil {
//		log.Errorw("RelationActionService.followAction", "err", err)
//		return err
//	}
//	log.Infow("Success to follow", "uid", uid, "toUid", toUid)
//	return nil
//}
//
//func UnfollowAction(uid, toUid int64) error {
//	// 判断是否已经关注
//	ctx := context.Background()
//	isFollowing, err := dal.QueryFollowing(ctx, uid, toUid)
//	if err != nil {
//		log.Errorw("RelationActionService.unfollowAction", "err", err)
//		return err
//	}
//	if !isFollowing {
//		log.Errorw("RelationActionService.unfollowAction", "err",
//			errno.ErrNotFollow)
//		return errno.ErrNotFollow
//	}
//
//	err = dal.DeleteFollowRelation(ctx, uid, toUid)
//	if err != nil {
//		log.Errorw("RelationActionService.unfollowAction", "err", err)
//	}
//	log.Infow("Success to unfollow", "uid", uid, "toUid", toUid)
//	return nil
//}
