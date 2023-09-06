package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
	"github.com/spf13/viper"
)

type VideoPublishListService struct {
	ctx context.Context
}

func NewVideoPublishListService(ctx context.Context) *VideoPublishListService {
	return &VideoPublishListService{ctx: ctx}
}

func (s *VideoPublishListService) GetPublishVideoList(req *VideoService.
	VideoPublishListRequest) ([]*VideoService.Video, error) {

	var videoList []*VideoService.Video
	// rpc 调用获取作者信息(含有是否关注信息)
	infoResp, err := rpc.UserClient.GetUserInfo(s.ctx, &UserService.UserInfoRequest{
		FromId: req.UserId,
		ToId:   req.ToUserId,
	})
	if err != nil {
		log.Errorw("VideoPublishListService failed to get user info", "err", err)
		return nil, err
	}

	if infoResp.User == nil {
		log.Errorw("VideoPublishListService get null user info", "err", err)
		return nil, errno.ErrUserNotExist
	}

	author := infoResp.User

	// 获取视频列表
	videos, err := dal.GetVideoListByUserId(s.ctx, req.ToUserId)
	if err != nil {
		log.Errorw("VideoPublishListService failed to get video list", "err", err)
	}
	publicHost := viper.GetString("minio.public.host")

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

	for _, video := range videos {
		_, isFavorite := favoriteSet[video.VideoID]
		videoList = append(videoList, &VideoService.Video{
			Id:            video.VideoID,
			Author:        author,
			PlayUrl:       GetPublicURL(video.PlayUrl, publicHost),
			CoverUrl:      GetPublicURL(video.CoverUrl, publicHost),
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}
	log.Debugw("VideoPublishListService get video list successfully.")
	return videoList, nil
}
