package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
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

	authorInfo := &VideoService.UserInfo{
		Id:              author.Id,
		Name:            author.Name,
		FollowCount:     author.FollowCount,
		FollowerCount:   author.FollowerCount,
		IsFollow:        author.IsFollow,
		Avatar:          author.Avatar,
		BackgroundImage: author.BackgroundImage,
		Signature:       author.Signature,
		TotalFavorited:  author.TotalFavorited,
		WorkCount:       author.WorkCount,
		FavoriteCount:   author.FavoriteCount,
	}

	// 获取视频列表
	videos, err := dal.GetVideoListByUserId(s.ctx, req.ToUserId)
	if err != nil {
		log.Errorw("VideoPublishListService failed to get video list", "err", err)
	}

	for _, video := range videos {
		videoList = append(videoList, &VideoService.Video{
			Id:            video.VideoID,
			Author:        authorInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			//IsFavorite:    false,
			Title: video.Title,
		})
	}
	log.Debugw("VideoPublishListService get video list successfully.")
	return videoList, nil
}
