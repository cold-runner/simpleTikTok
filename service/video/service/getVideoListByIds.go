package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/dal/model"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
	"github.com/spf13/viper"
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
	publicHost := viper.GetString("minio.public.host")
	authorInfos := authorInfosResp.UserInfoList

	// 获取视频点赞信息
	resp, err := rpc.SocialClient.GetFavoriteVidList(s.ctx,
		&SocialService.GetFavoriteVideoByUidRequest{
			UserId: req.UserId,
		})
	if err != nil {
		log.Errorw("VideoFeedService failed to get favorite videos")
	}

	// 将用户 id 转化为集合方便查询
	favoriteSet := make(map[int64]struct{})
	for _, vid := range resp.GetFavoriteIdList() {
		favoriteSet[vid] = struct{}{}
	}
	for i := 0; i < len(videoList); i++ {
		_, isFavorite := favoriteSet[videoList[i].VideoID]

		videos = append(videos, &VideoService.Video{
			Id:            videoList[i].VideoID,
			Author:        authorInfos[i],
			PlayUrl:       GetPublicURL(videoList[i].PlayUrl, publicHost),
			CoverUrl:      GetPublicURL(videoList[i].CoverUrl, publicHost),
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
			IsFavorite:    isFavorite,
			Title:         videoList[i].Title,
		})
	}
	log.Debugw("VideoListService get video list successfully.")
	return videos, nil

}
