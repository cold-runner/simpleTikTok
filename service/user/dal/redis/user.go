package redis

import (
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"strconv"
	"time"
)

var (
	cacheExpiration = time.Hour * 48
)

//func GetUser(ctx context.Context, uid int64) (*model.User, error) {
//	// 从Redis中尝试获取用户
//	userJson, err := redisUserInfoClient.Get(ctx, fmt.Sprintf("user:%d", uid)).Result()
//	if err == redis.Nil {
//		log.Infow("User %d not found in cache, retrieving from database", uid)
//	} else if err != nil {
//		log.Infow("Error retrieving user %d from cache: %v", uid, err)
//	} else {
//		user := &model.User{}
//		err = json.Unmarshal([]byte(userJson), user)
//		if err != nil {
//			log.Infow("Error unmarshalling user %d from cache: %v", uid, err)
//		} else {
//			log.Infow("User %d found in cache", uid)
//			return user, nil
//		}
//	}
//
//	// 如果Redis中没有缓存，则从数据库中获取
//	user, err := dal.GetUserByUserId(ctx, uid)
//	if err != nil {
//		return nil, err
//	}
//
//	// 将用户信息缓存到Redis
//	if err = cacheUser(ctx, uid, user); err != nil {
//		log.Infow("Error caching user %d: %v", uid, err)
//	}
//
//	return user, nil
//}
//
//func cacheUser(ctx context.Context, uid int64, user *model.User) error {
//	userJson, err := json.Marshal(user)
//	if err != nil {
//		return fmt.Errorf("error marshaling user %d for cache: %v", uid, err)
//	}
//
//	err = redisUserInfoClient.Set(ctx, fmt.Sprintf("user:%d", uid), userJson, 0).Err() // 可以设置过期时间
//	if err != nil {
//		return fmt.Errorf("error caching user %d: %v", uid, err)
//	}
//
//	return nil
//}

func CacheStableUserFields(user *model.User) error {
	uidStr := strconv.Itoa(int(user.ID))
	stableFields := map[string]interface{}{
		"name":             user.Username,
		"avatar":           user.Avatar,
		"background_image": user.BackgroundImage,
		"signature":        user.Signature,
	}
	if err := redisUserInfoClient.HMSet(Ctx, uidStr,
		stableFields).Err(); err != nil {
		log.Errorw("Error caching user %d: %v", user.ID, err)
		return err
	}
	// 设置过期时间
	redisUserInfoClient.Expire(Ctx, uidStr, cacheExpiration)
	return nil
}

// GetStableUserFieldsByUserId 通过用户ID从Redis中获取稳定字段
func GetStableUserFieldsByUserId(uid int64) (*model.StableUserFields, error) {
	if redisUserInfoClient != nil {
		var vals []interface{}
		var err error
		// 将用户ID转换为字符串形式，以便作为 Redis key 使用
		uidStr := strconv.Itoa(int(uid))

		// 初始化StableUserFields结构体
		stableFields := &model.StableUserFields{}

		// 使用 HMGet 从 Redis 中获取稳定字段
		fieldNames := []string{"name", "avatar", "background_image", "signature"}
		vals, err = redisUserInfoClient.HMGet(Ctx, uidStr, fieldNames...).Result()
		if err != nil {
			log.Errorw("Error retrieving stable fields from cache", "uid", uid, "err", err)
			return nil, err
		}

		// 判断是否在缓存中找到了数据
		if len(vals) == 0 || vals[0] == nil {
			log.Infow("No stable fields found in cache for user", "uid", uid)
			return nil, nil
		}

		// 将取得的数据赋值到StableUserFields结构体中
		stableFields.Name = vals[0].(string)
		stableFields.Avatar = vals[1].(string)
		stableFields.BackgroundImage = vals[2].(string)
		stableFields.Signature = vals[3].(string)

		return stableFields, nil
	} else {
		log.Errorw("redis is not connected")
		return nil, errno.ErrRedisNotInitialized
	}

}
