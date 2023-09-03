package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/dal/model"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
	"sync"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{ctx: ctx}
}

func (s *VideoListService) GetVideoListByIds(req *VideoService.
	VideoPublishListByIdsRequest) ([]*VideoService.Video, error) {
	var wg sync.WaitGroup
	var videoList []*model.Video
	var videos []*VideoService.Video
	var authorIds []int64

	videoList, err := dal.MGetVideoListByIds(s.ctx, req.ToVideoIds)
	if err != nil {
		log.Errorw("VideoListService failed to get video list", "err", err)
		return nil, err
	}

	// 获取作者信息
	wg.Add(1)
	go func() {
		for _, video := range videoList {
			authorIds = append(authorIds, video.AuthorID)
		}
		wg.Done()
	}()
	wg.Wait()

	// rpc 调用获取作者信息(含有是否关注信息)
	authorInfosResp, err := rpc.RelationClient.QueryUserInfosWithRelation(s.
		ctx,
		&RelationService.QueryUserInfosWithRelationRequest{
			Uid:    req.UserId,
			IdList: authorIds,
		})
	if err != nil {
		log.Errorw("VideoListService failed to get user info", "err", err)
	}

	authorInfos := authorInfosResp.UserInfoList
	for i := 0; i < len(videoList); i++ {
		videos = append(videos, &VideoService.Video{
			Id: videoList[i].VideoID,
			Author: &VideoService.UserInfo{
				Id:              authorInfos[i].Id,
				Name:            authorInfos[i].Name,
				FollowCount:     authorInfos[i].FollowCount,
				FollowerCount:   authorInfos[i].FollowerCount,
				IsFollow:        authorInfos[i].IsFollow,
				Avatar:          authorInfos[i].Avatar,
				BackgroundImage: authorInfos[i].BackgroundImage,
				Signature:       authorInfos[i].Signature,
				TotalFavorited:  authorInfos[i].TotalFavorited,
				WorkCount:       authorInfos[i].WorkCount,
				FavoriteCount:   authorInfos[i].FavoriteCount,
			},
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
			//IsFavorite:    false,
			Title: videoList[i].Title,
		})
	}
	log.Debugw("VideoListService get video list successfully.")
	return videos, nil

}
