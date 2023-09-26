package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/mq"
)

type UpdateVideoInfoService struct {
	ctx context.Context
}

func NewUpdateVideoInfoService(ctx context.Context) *UpdateVideoInfoService {
	return &UpdateVideoInfoService{ctx: ctx}
}

func (s *UpdateVideoInfoService) UpdateVideoFavoriteCount(req *VideoService.UpdateVideoFavoriteCountRequest) (
	err error) {
	request := mq.UpdateVideoInfoRequest{
		VideoId:             req.VideoId,
		ActionType:          req.ActionType,
		UpdateFavoriteCount: true,
	}
	err = mq.VideoInfoProducer.Publish(request)
	if err != nil {
		log.Errorw("UpdateVideoFavoriteCount failed to publish message", "err",
			err)
	}
	log.Debugw("UpdateVideoFavoriteCount publish message success", "request",
		request)
	return nil
}

func (s *UpdateVideoInfoService) UpdateVideoCommentCount(req *VideoService.UpdateVideoCommentCountRequest) (
	err error) {
	request := mq.UpdateVideoInfoRequest{
		VideoId:            req.VideoId,
		ActionType:         req.ActionType,
		UpdateCommentCount: true,
	}
	err = mq.VideoInfoProducer.Publish(request)
	if err != nil {
		log.Errorw("UpdateVideoCommentCount failed to publish message", "err",
			err)
	}
	log.Debugw("UpdateVideoCommentCount publish message success", "request",
		request)
	return nil
}
