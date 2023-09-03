package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"github.com/cold-runner/simpleTikTok/service/user/dal/redis"
	"github.com/cold-runner/simpleTikTok/service/user/rpc"
)

type GetUserInfoService struct {
	ctx context.Context
}

// NewGetUserInfoService create new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// GetUserInfoByName get user info by username
func (s *GetUserInfoService) GetUserInfoByName(username string) (*model.User,
	error) {
	u, err := dal.GetUserByUserName(s.ctx, username)
	if err != nil {
		return u, errno.ErrUserNotExist
	}
	return u, nil
}

// GetUserInfoById get user info by user id
func (s *GetUserInfoService) GetUserInfoById(id int64, toUid int64) (*model.UserInfo, error) {

	userInfo := &model.UserInfo{ID: toUid}

	// RPC 调用
	resp, err := rpc.RelationClient.QueryRelation(s.ctx, &RelationService.RelationQueryRequest{
		Uid:      id,
		ToUserId: toUid,
	})
	if err != nil {
		log.Errorw("get relation failed", "err", err)
	} else {
		userInfo.IsFollow = resp.IsFollow
	}

	// 尝试从 Redis 获取稳定字段

	sFields, err := redis.GetStableUserFieldsByUserId(toUid)
	if err != nil {
		log.Errorw("get stable user fields cache failed", "err", err)
	}

	// 如果在 Redis 中找到了稳定字段
	if sFields != nil {
		log.Debugw("get stable user fields from redis success!", "sFields",
			sFields)
		// 设置稳定字段
		userInfo.Name = sFields.Name
		userInfo.Avatar = sFields.Avatar
		userInfo.BackgroundImage = sFields.BackgroundImage
		userInfo.Signature = sFields.Signature

		// 从数据库获取动态字段
		dFields, err := dal.GetDynamicUserFields(s.ctx, toUid)
		if err != nil {
			log.Errorw("get dynamic user fields failed", "err", err)
		} else {
			// 设置动态字段
			userInfo.FollowCount = dFields.FollowCount
			userInfo.FollowerCount = dFields.FollowerCount
			userInfo.TotalFavorited = dFields.TotalFavorited
			userInfo.WorkCount = dFields.WorkCount
			userInfo.FavoriteCount = dFields.FavoriteCount
			return userInfo, nil
		}
	}

	// 如果在 Redis 中没有找到或获取动态字段失败，则从数据库获取所有字段
	log.Debugw("get stable user fields from redis failed, use mysql instead!")
	user, err := dal.GetUserByUserId(s.ctx, toUid)
	if err != nil {
		log.Errorw("get user by user id failed", "err", err)
		return nil, err
	}

	// 设置所有字段
	userInfo.Name = user.Username
	userInfo.Avatar = user.Avatar
	userInfo.BackgroundImage = user.BackgroundImage
	userInfo.Signature = user.Signature
	userInfo.FollowCount = user.FollowCount
	userInfo.FollowerCount = user.FollowerCount
	userInfo.TotalFavorited = user.TotalFavorited
	userInfo.WorkCount = user.WorkCount
	userInfo.FavoriteCount = user.FavoriteCount

	return userInfo, nil
}
