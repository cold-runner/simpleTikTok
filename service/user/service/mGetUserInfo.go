package service

import (
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"github.com/cold-runner/simpleTikTok/service/user/dal/redis"
)

func (s *GetUserInfoService) MGetUserInfoById(toUids []int64) ([]*model.UserInfo, error) {
	var u []*model.UserInfo
	for _, toUid := range toUids {
		userInfo, err := s.mGetUserInfoById(toUid)
		if err != nil {
			return nil, errno.ErrUserNotExist
		}
		u = append(u, userInfo)
	}
	log.Debugw("Get user info list success!", "user_info_list", u)
	return u, nil
}

// mGetUserInfoById 通过用户 ID 获取用户信息，但是不包括关注关系
func (s *GetUserInfoService) mGetUserInfoById(toUid int64) (*model.UserInfo, error) {

	userInfo := &model.UserInfo{ID: toUid}

	// 尝试从 Redis 获取稳定字段

	sFields, err := redis.GetStableUserFieldsByUserId(toUid)
	if err != nil {
		log.Errorw("get stable user fields cache failed", "err", err)
	}

	// 如果在 Redis 中找到了稳定字段
	if sFields != nil {
		log.Debugw("get stable user fields from redis success!", "sFields",
			sFields, "user_id", toUid)
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
			log.Debugw("get complete user info success!", "user_info", userInfo)
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
	log.Debugw("get user by user id success!", "user", user)
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
	err = redis.CacheStableUserFields(user)
	if err != nil {
		log.Errorw("cache stable user fields failed", "err", err)
	}
	log.Debugw("Cache stable user fields success!")

	log.Debugw("get complete user info success!", "user_info", userInfo)
	return userInfo, nil
}
