package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	dtm "github.com/cold-runner/simpleTikTok/service/dtm/request"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction 点赞/取消点赞
// 给一个视频点赞需要完成如下的业务逻辑：
// 1. 在视频点赞关系表中添加一条记录 （mysql 数据库操作)
// 2. 在视频表中的点赞数 favorite_count +1 (rpc 调用)
// 3. 在用户表中的用户喜欢数favorite_count +1 (rpc 调用)
// 4. 在用户表中的作者获赞数量 total_favorited +1 (rpc 调用)
func (s *FavoriteActionService) FavoriteAction(req *SocialService.
	FavoriteActionRequest) (err error) {
	log.Debugw("start doing favorite action request", "req", req)
	uid, toVid, actionType := req.GetUserId(), req.GetToVideoId(), req.ActionType
	// 获取视频作者的 id
	authorIdResp, err := rpc.VideoClient.GetAuthorIdByVideoId(s.ctx,
		&VideoService.GetAuthorIdByVideoIdRequest{
			VideoId: toVid,
		})
	if err != nil {
		log.Errorw("get author id by video id failed", "err", err)
		return err
	}

	authorID := authorIdResp.GetAuthorId()
	err = dtm.FavoriteActionRequest(&dtm.FavoriteActionRequestToDTM{
		UserId:     uid,
		VideoId:    toVid,
		AuthorId:   authorID,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorw("favorite action request failed", "err", err)
		return err
	}
	log.Debugw("favorite action request success")
	return nil
}
