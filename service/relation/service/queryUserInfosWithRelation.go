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

type UserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoWithRelationService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) GetUserInfoList(req *RelationService.QueryUserInfosWithRelationRequest) ([]*UserService.User, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	log.Debugw("GetUserFollowList", "req", req)
	// 查询用户 id 的关注列表, 先通过 redis 缓存查找
	followIdList, err := redis.GetFollowList(req.GetUid())
	if err != nil {
		log.Errorw("Error querying follow list in redis", "err", err)
	}

	// 没有找到缓存，使用数据库查找，然后加入缓存
	if len(followIdList) == 0 || followIdList == nil {
		log.Debugw("GetUserFollowList cache not found, use database")
		followIdList, err = dal.GetFollowListByUserID(s.ctx, req.GetUid())

		if err != nil {
			log.Errorw("Error querying follow list in database", "err", err)
			return nil, err
		}
		// 用户没有任何关注
		if len(followIdList) == 0 {
			log.Debugw("GetUserFollowList no follow found")
		} else {
			// 将关注列表缓存到 Redis 中
			log.Debugw("cache follow list to redis")
			redis.CacheFollowList(req.GetUid(), followIdList)
		}
	}

	log.Debugw("follow list cache found", "followIdList", followIdList)

	// 并发获取两个内容
	// 1.查询所需要的用户信息：通过 rpc 调用获取
	// 2.所查询用户的是否关注的信息：查询数据库
	var users []*UserService.User
	var followSet map[int64]struct{}
	var followList []int64

	go func() {
		log.Debugw("Calling user rpc to get user info")
		mUserInfoResp, err := rpc.MGetUserInfo(s.ctx,
			&UserService.MGetUserInfoRequest{
				Uids: req.GetIdList(),
			})

		if err != nil {
			log.Errorw("Error calling user service to get user info", "err", err)
		}

		users = mUserInfoResp.Users
		log.Debugw("User info without is_follow", "users",
			users)
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
			// 如果找到了缓存，将关注列表转换为 map
			for _, v := range followList {
				followSet[v] = struct{}{}
			}
		}
		log.Debugw("User follow list from database", "followList", followList)
		log.Debugw("User follow set from database", "followSet", followSet)
		wg.Done()
	}()
	// 等待并发操作完成
	wg.Wait()

	// 补充关注信息, 写入要返回的结果集
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
