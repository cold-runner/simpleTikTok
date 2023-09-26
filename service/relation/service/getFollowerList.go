package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/dal"
	"github.com/cold-runner/simpleTikTok/service/relation/dal/redis"
	"github.com/cold-runner/simpleTikTok/service/relation/rpc"
	"sync"
)

type GetFollowerListService struct {
	ctx context.Context
}

func NewGetFollowerListService(ctx context.Context) *GetFollowerListService {
	return &GetFollowerListService{ctx: ctx}
}

func (s *GetFollowerListService) GetUserFollowerList(req *RelationService.
	RelationFollowerListRequest) ([]*UserService.User, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	log.Debugw("GetUserFollowerList", "req", req)
	followerIdList, err := redis.GetFollowerList(req.GetToUid())
	if err != nil {
		log.Errorw("Error querying follower list in redis", "err", err)
	}
	// 没有找到粉丝列表缓存，使用数据库查找，然后加入缓存
	if len(followerIdList) == 0 || followerIdList == nil {
		log.Debugw("GetUserFollowerList cache not found, use database")
		followerIdList, err := dal.GetFollowerListByUserID(s.ctx, req.GetToUid())

		if err != nil {
			log.Errorw("Error querying follower list in database", "err", err)
			return nil, err
		}
		// 用户没有任何粉丝
		if len(followerIdList) == 0 {
			log.Debugw("GetUserFollowerList no follower found")
			return nil, nil
		}

		// 将粉丝列表缓存到 Redis 中
		log.Debugw("cache follower list to redis")
		redis.CacheFollowerList(req.GetToUid(), followerIdList)
	}
	log.Debugw("follower list cache found", "followerIdList", followerIdList)
	// 并发获取两个内容
	// 1. 查询的用户信息 通过 rpc 调用获取
	// 2. 使用者关注信息 查询数据库
	var users []*UserService.User
	var followSet map[int64]struct{}
	var followList []int64

	go func() {
		log.Debugw("Calling user rpc to get user info")
		mUserInfoResp, err := rpc.MGetUserInfo(s.ctx, &UserService.MGetUserInfoRequest{
			Uids: followerIdList,
		})

		if err != nil {
			log.Errorw("Error calling user rpc to get user info", "err", err)
		}

		users = mUserInfoResp.Users
		log.Debugw("User info without is_follow", "users", users)
		wg.Done()
	}()

	go func() {
		var err error
		// 先尝试从 redis 获取关注列表
		followList, err = redis.GetFollowList(req.GetUid())
		if err != nil {
			log.Errorw("Error querying follow list in redis", "err", err)
		}
		// 如果没有找到缓存，使用数据库查找
		if followList == nil || len(followList) == 0 {
			log.Debugw("Get my follow list cache not found, use database")
			followSet, followList, err = dal.GetUserFollowSetAndList(s.ctx, req.GetUid())
			if err != nil {
				log.Errorw("Error querying follow list in database", "err", err)
			}
			// 将关注列表缓存到 Redis 中
			redis.CacheFollowList(req.GetUid(), followList)
		} else {
			if followSet == nil {
				followSet = make(map[int64]struct{})
			}
			// 如果找到了缓存，将关注列表转换为map
			for _, v := range followList {
				followSet[v] = struct{}{}
			}
		}
		log.Debugw("User follow list from database", "followList", followList)
		log.Debugw("User follow set from database", "followSet", followSet)
		wg.Done()
	}()
	// 等待所有并发操作完成
	wg.Wait()

	// 将用户信息和关注信息合并，写入到返回值中
	userInfos := make([]*UserService.User, 0)
	for _, u := range users {
		_, isFollow := followSet[u.Id]
		u.IsFollow = isFollow
		userInfos = append(userInfos, u)
	}
	if len(users) != len(userInfos) {
		log.Errorw("Error converting user info")
	}
	log.Debugw("User info with follow info", "userInfos", userInfos)
	return userInfos, nil
}
