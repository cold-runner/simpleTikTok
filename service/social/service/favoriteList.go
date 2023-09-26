package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) GetFavoriteList(ctx context.Context,
	req *SocialService.FavoriteListRequest) ([]*VideoService.Video, error) {
	uid, toUid := req.GetUserId(), req.GetToUserId()
	// 获取点赞的视频列表
	vidList, err := dal.GetFavoriteVideoListByUid(ctx, toUid)
	if err != nil {
		log.Errorw("get user favorite video list failed", "err", err)
		return nil, err
	}

	// 获取用户个人点赞的视频列表集合
	userFavorVidSet, err := dal.GetFavoriteVideoSetByUid(ctx, uid)
	if err != nil {
		log.Errorw("get user favorite video set failed", "err", err)
		return nil, err
	}

	// rpc 获取视频的信息
	videoResp, err := rpc.VideoClient.VideoPublishListByIds(ctx,
		&VideoService.VideoPublishListByIdsRequest{
			UserId:     uid,
			ToVideoIds: vidList,
		})
	if err != nil {
		log.Errorw("get video publish list by vids failed", "err", err)
		return nil, err
	}
	videoList := videoResp.GetVideoList()
	// 处理视频中的是否点赞信息
	for _, video := range videoList {
		if _, ok := userFavorVidSet[video.Id]; ok {
			video.IsFavorite = true
		}
	}
	return videoList, nil
}
