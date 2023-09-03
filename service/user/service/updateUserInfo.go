package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/mq"
)

type UpdateUserInfoService struct {
	ctx context.Context
}

func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

func (s *UpdateUserInfoService) UpdateUserInfo(req interface{}) error {
	var mqRequest mq.UpdateUserInfoRequest
	log.Debugw("start update user info", "req", req)
	switch r := req.(type) {
	case *UserService.ChangeUserFavoriteCountRequest:
		mqRequest = mq.UpdateUserInfoRequest{
			ToUserId:            r.Id,
			ActionType:          r.ActionType,
			UpdateFavoriteCount: true,
		}

	case *UserService.ChangeUserWorkCountRequest:
		mqRequest = mq.UpdateUserInfoRequest{
			ToUserId:        r.Id,
			ActionType:      r.ActionType,
			UpdateWorkCount: true,
		}

	case *UserService.ChangeUserTotalFavoritedCountRequest:
		mqRequest = mq.UpdateUserInfoRequest{
			ToUserId:             r.Id,
			ActionType:           r.ActionType,
			UpdateTotalFavorited: true,
		}
	}

	err := mq.SocialInfoUpdateProducer.Publish(mqRequest)
	if err != nil {
		log.Errorw("user service failed to publish message", "err", err)
		return err
	}
	log.Debugw("user service publish message success!", "request:", mqRequest)
	return nil
}
