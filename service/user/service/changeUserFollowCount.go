package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/mq"
)

type ChangeUserFollowService struct {
	ctx context.Context
}

func NewChangeUserFollowService(ctx context.Context) *ChangeUserFollowService {
	return &ChangeUserFollowService{ctx: ctx}
}

func (s *ChangeUserFollowService) UpdateUserFollowInfo(req *UserService.
	ChangeUserFollowCountRequest) error {
	uid, toUid, actionType := req.Id, req.ToUserId, req.ActionType
	request := mq.ChangeUserFollowCountRequest{
		UserId:     uid,
		ToUserId:   toUid,
		ActionType: actionType,
	}

	err := mq.UserInfoUpdateProducer.Publish(request)
	if err != nil {
		log.Errorw("user service failed to publish message", "err", err)
	} else {
		log.Infow("user service publish message successfully")
	}
	return nil
}
