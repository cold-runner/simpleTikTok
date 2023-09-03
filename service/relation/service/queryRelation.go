package service

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/relation/dal"
	"github.com/cold-runner/simpleTikTok/service/relation/dal/redis"
)

type QueryRelationService struct {
	ctx context.Context
}

func NewQueryRelationService(ctx context.Context) *QueryRelationService {
	return &QueryRelationService{ctx: ctx}
}

// QueryRelation 查询关系
func (s *QueryRelationService) QueryRelation(req *RelationService.
	RelationQueryRequest) (bool, error) {

	// 查看个人信息的时候，不需要判断关系
	if req.Uid == req.ToUserId {
		log.Debugw("user is querying his own info, no need to query relation")
		return true, nil
	}

	isFollow, err := redis.QueryUserFollowing(req.Uid, req.ToUserId)
	if err != nil {
		log.Errorw("Error querying relation in redis", "err", err)
	}
	// 在缓存中找到关系，返回结果
	if isFollow {
		log.Debugw("found relation in redis cache!")
		return true, nil
	}
	log.Debugw("relation not found in redis cache, querying db...")
	following, err := dal.QueryFollowing(s.ctx, req.Uid, req.ToUserId)
	if err != nil {
		log.Errorw("Error querying relation in db", "err", err)
	}
	// 在数据中查询到关系，并且没有缓存，就将关系缓存到 Redis 中
	if following {
		log.Debugw("found relation in db, caching relation to redis!")
		redis.CacheRelation(req.Uid, req.ToUserId)
		return true, nil
	} else {
		return false, nil
	}
}
