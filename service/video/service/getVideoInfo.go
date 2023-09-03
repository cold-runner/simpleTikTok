package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
)

type GetVideoInfoService struct {
	ctx context.Context
}

func NewGetVideoInfoService(ctx context.Context) *GetVideoInfoService {
	return &GetVideoInfoService{ctx: ctx}
}

func (s *GetVideoInfoService) GetAuthorID(req *VideoService.
	GetAuthorIdByVideoIdRequest) (authorId int64, err error) {

	authorId, err = dal.GetAuthorIdByVideoId(s.ctx, req.VideoId)
	if err != nil {
		log.Errorw("GetVideoInfoService failed to get author id", "err", err)
		return -1, err
	}

	return authorId, nil
}
