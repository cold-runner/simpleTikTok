package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/mq"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *VideoService.
	VideoPublishActionRequest) error {
	request := mq.VideoPublishRequest{
		AuthorID: req.UserId,
		Data:     req.Data,
		Title:    req.Title,
	}

	_, err := rpc.UserClient.ChangeUserWorkCount(context.Background(),
		&UserService.ChangeUserWorkCountRequest{
			Id:         req.GetUserId(),
			ActionType: 1,
		})

	if err != nil {
		log.Errorw("failed to update user work count", "err", err)
		return err
	}

	err = mq.VideoProducer.Publish(request)
	if err != nil {
		log.Errorw("VideoServer failed to publish message", "err", err)
	}
	log.Debugw("VideoServer success to publish message", "Title",
		request.Title, "AuthorID", request.AuthorID)
	return nil
}
