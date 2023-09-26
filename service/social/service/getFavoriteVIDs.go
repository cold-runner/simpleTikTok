package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
)

type FavoriteVIDService struct {
	ctx context.Context
}

func NewFavoriteVIDService(ctx context.Context) *FavoriteVIDService {
	return &FavoriteVIDService{ctx: ctx}
}

func (s *FavoriteVIDService) GetFavoriteIDsByUserID(ctx context.Context,
	req *SocialService.GetFavoriteVideoByUidRequest) ([]int64, error) {

	vidList, err := dal.GetFavoriteVideoIDListByUid(ctx, req.GetUserId())

	if err != nil {
		log.Errorw("get user favorite video id list failed", "err", err)
	}

	return vidList, nil
}
