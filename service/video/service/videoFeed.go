package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type VideoFeedService struct {
	ctx context.Context
}

func NewVideoFeedService(ctx context.Context) *VideoFeedService {
	return &VideoFeedService{ctx: ctx}
}

func (s *VideoFeedService) VideoFeed(req *VideoService.VideoFeedRequest) (
	[]*VideoService.Video, int64, error) {
	var latestTime int64
	// 没有传入latestTime，就默认为当前时间
	if req.LatestTime == 0 {
		latestTime = time.Now().Local().UnixMilli()
	} else {
		latestTime = int64(req.LatestTime)
	}

	videos, authorIDs, err := dal.GetVideosBeforeTimeWithAuthors(s.ctx,
		latestTime)
	if err != nil {
		log.Errorw("VideoFeedService failed to get videos", "err", err)
		return nil, 0, err
	}
	if len(videos) == 0 {
		log.Debugw("VideoFeedService no videos found")
		return nil, 0, nil
	}

	// 更新latestTime
	if len(videos) > 0 {
		latestTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	} else {
		latestTime = 0
	}

	videoList := make([]*VideoService.Video, 0, len(videos))
	// 游客用户，不需要获取关注列表
	if req.UserId == 0 {
		infoResp, err := rpc.UserClient.MGetUserInfo(s.ctx,
			&UserService.MGetUserInfoRequest{
				Uids: authorIDs,
			})
		if err != nil {
			log.Errorw("VideoFeedService failed to get user infos for"+
				" visitor",
				"err", err)
			return nil, 0, err
		}
		authorInfos := infoResp.Users

		for i := 0; i < len(videos) && i < len(authorInfos); i++ {
			videoList = append(videoList, &VideoService.Video{
				Id:            videos[i].VideoID,
				Author:        authorInfos[i],
				PlayUrl:       videos[i].PlayUrl,
				CoverUrl:      videos[i].CoverUrl,
				FavoriteCount: videos[i].FavoriteCount,
				CommentCount:  videos[i].CommentCount,
				IsFavorite:    false,
				Title:         videos[i].Title,
			})
		}
	} else {
		// 非游客用户，需要获取关注信息
		infoResp, err := rpc.RelationClient.QueryUserInfosWithRelation(s.ctx,
			&RelationService.QueryUserInfosWithRelationRequest{
				Uid:    req.UserId,
				IdList: authorIDs,
			})
		authorInfos := infoResp.UserInfoList

		if err != nil {
			log.Errorw("VideoFeedService failed to get user infos")
			return nil, 0, err
		}

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
		publicHost := viper.GetString("minio.public.host")
		for i := 0; i < len(videos) && i < len(authorInfos); i++ {
			_, isFavorite := favoriteSet[videos[i].VideoID]
			videoList = append(videoList, &VideoService.Video{
				Id:            videos[i].VideoID,
				Author:        authorInfos[i],
				PlayUrl:       GetPublicURL(videos[i].PlayUrl, publicHost),
				CoverUrl:      GetPublicURL(videos[i].CoverUrl, publicHost),
				FavoriteCount: videos[i].FavoriteCount,
				CommentCount:  videos[i].CommentCount,
				IsFavorite:    isFavorite,
				Title:         videos[i].Title,
			})
		}
	}

	return videoList, latestTime, nil
}

func GetPublicURL(localURL string, newHost string) string {
	return strings.Replace(localURL, "localhost", newHost, -1)
}
