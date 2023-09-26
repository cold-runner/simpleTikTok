package redis

import (
	"github.com/cold-runner/simpleTikTok/pkg/log"
	_ "github.com/cold-runner/simpleTikTok/service/user/dal/redis"
	"strconv"
	"time"
)

const (
	// 1小时的缓存过期时间
	cacheExpiration       = time.Hour * 48
	FollowerListThreshold = 1000
)

// CacheRelation 将关注关系缓存到 Redis 中
func CacheRelation(uid, toUid int64) {
	// 将用户ID转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SAdd 在 Redis 中添加关注关系
	if err := redisRelationClient.SAdd(Ctx, uidStr, toUid).Err(); err != nil {
		log.Errorw("Error caching relation", "err", err)
	}

	// 设置过期时间
	if err := redisRelationClient.Expire(Ctx, uidStr,
		cacheExpiration).Err(); err != nil {
		log.Errorw("Error setting expiration for relation", "err", err)
	}

	return
}

// CacheFollowList 将关注列表缓存到 Redis 中
func CacheFollowList(uid int64, toUids []int64) {
	// 将用户ID转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SAdd 在 Redis 中添加关注的用户列表
	for _, id := range toUids {
		if err := redisFollowClient.SAdd(Ctx, uidStr, id).Err(); err != nil {
			log.Errorw("Error caching follow list", "err", err)
		}
	}

	// 设置过期时间
	if err := redisFollowClient.Expire(Ctx, uidStr,
		cacheExpiration).Err(); err != nil {
		log.Errorw("Error setting expiration for follow list", "err", err)
		return
	}

	return
}

// CacheFollowerList 将粉丝列表缓存到 Redis 中
func CacheFollowerList(uid int64, fanIds []int64) {
	if len(fanIds) > FollowerListThreshold {
		return
	}

	for _, fanId := range fanIds {
		if err := redisFollowerClient.SAdd(Ctx, strconv.Itoa(int(uid)), fanId).Err(); err != nil {
			log.Errorw("Error caching follower list", "err", err)
		}
	}

	if err := redisFollowerClient.Expire(Ctx, strconv.Itoa(int(uid)), cacheExpiration).Err(); err != nil {
		log.Errorw("Error setting expiration for follower list", "err", err)
	}
	return
}

// QueryUserFollowing 查询用户是否关注了另一个用户
func QueryUserFollowing(uid, toUid int64) (bool, error) {
	// 将用户ID转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SIsMember 函数来检查 `toUid` 是否存在于该用户的关注列表中
	isMember, err := redisRelationClient.SIsMember(Ctx, uidStr, toUid).Result()
	if err != nil {
		return false, err
	}
	// 缓存命中，刷新过期时间
	if isMember {
		redisRelationClient.Expire(Ctx, uidStr,
			cacheExpiration)
		return true, nil
	}
	return false, nil
}

// GetFollowList 获取用户的关注列表
func GetFollowList(uid int64) ([]int64, error) {
	// 将用户ID转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SMembers 函数获取该用户的完整关注列表
	followList, err := redisFollowClient.SMembers(Ctx, uidStr).Result()
	if err != nil {
		return nil, err
	}

	// 将关注列表从字符串数组转换为 int64 数组
	var result []int64
	for _, idStr := range followList {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Errorw("Error parsing id", "err", err)
			continue
		}
		result = append(result, id)
	}

	return result, nil
}

// GetFollowerList 获取用户的粉丝列表
func GetFollowerList(uid int64) ([]int64, error) {
	// 将用户ID转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SMembers 函数获取该用户的完整粉丝列表
	followerList, err := redisFollowerClient.SMembers(Ctx, uidStr).Result()
	if err != nil {
		return nil, err
	}

	// 将粉丝列表从字符串数组转换为 int64 数组
	var result []int64
	for _, idStr := range followerList {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Errorw("Error parsing id", "err", err)
			continue
		}
		result = append(result, id)
	}

	return result, nil
}

// RemoveRelationCache 从 Redis 中删除关注关系
func RemoveRelationCache(uid, toUid int64) {
	// 将用户ID转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SRem 在 Redis 中删除关注关系
	if err := redisRelationClient.SRem(Ctx, uidStr, toUid).Err(); err != nil {
		log.Errorw("Error removing relation from cache", "err", err)
	}
	return
}

// DeleteFollowListCache 删除某个用户的关注列表缓存
func DeleteFollowListCache(uid int64) {
	uidStr := strconv.Itoa(int(uid))

	if err := redisFollowClient.Del(Ctx, uidStr).Err(); err != nil {
		log.Errorw("Error deleting follow list cache", "err", err)
	}
}

// DeleteFollowerListCache 删除某个用户的粉丝列表缓存
func DeleteFollowerListCache(uid int64) {
	uidStr := strconv.Itoa(int(uid))

	if err := redisFollowerClient.Del(Ctx, uidStr).Err(); err != nil {
		log.Errorw("Error deleting follower list cache", "err", err)
	}
}

// RemoveFromFollowerList 从粉丝列表中移除一个用户
func RemoveFromFollowerList(uid, fanId int64) error {
	uidStr := strconv.Itoa(int(uid))

	if err := redisFollowerClient.SRem(Ctx, uidStr, fanId).Err(); err != nil {
		log.Errorw("Error removing from follower list cache", "err", err)
		return err
	}

	return nil
}

// RemoveFromFollowList 从关注列表中移除一个用户
func RemoveFromFollowList(uid, toUid int64) error {
	uidStr := strconv.Itoa(int(uid))

	if err := redisFollowClient.SRem(Ctx, uidStr, toUid).Err(); err != nil {
		log.Errorw("Error removing from follow list cache", "err", err)
		return err
	}

	return nil
}

// AddToFollowerList 在粉丝列表中添加一个用户
func AddToFollowerList(uid, fanId int64) error {
	// 将用户 ID 转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SAdd 在 Redis 中添加粉丝
	if err := redisFollowerClient.SAdd(Ctx, uidStr, fanId).Err(); err != nil {
		log.Errorw("Error adding to follower list cache", "err", err)
		return err
	}
	if err := redisFollowerClient.Expire(Ctx, uidStr, cacheExpiration).Err(); err != nil {
		log.Errorw("Error setting expiration for follower list", "err", err)
		return err
	}
	return nil
}

// AddToFollowList 在关注列表中添加一个用户
func AddToFollowList(uid, toUid int64) error {
	// 将用户 ID 转换为字符串形式，以便作为 Redis key 使用
	uidStr := strconv.Itoa(int(uid))

	// 使用 SAdd 在 Redis 中添加关注者
	if err := redisFollowClient.SAdd(Ctx, uidStr, toUid).Err(); err != nil {
		log.Errorw("Error adding to follow list cache", "err", err)
		return err
	}

	if err := redisFollowClient.Expire(Ctx, uidStr, cacheExpiration).Err(); err != nil {
		log.Errorw("Error setting expiration for follow list", "err", err)
		return err
	}

	return nil
}
