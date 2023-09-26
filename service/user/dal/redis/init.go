package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var redisUserInfoClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	redisUserInfoClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db.userInfo"), // use default DB
	})
}
