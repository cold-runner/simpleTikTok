package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// var redisClient *redis.Client
var redisRelationClient *redis.Client
var redisFollowClient *redis.Client
var redisFollowerClient *redis.Client
var Ctx = context.Background()

type redisConfig struct {
	Addr     string
	Password string
	DB       int
}

func InitRedis() {
	//redisClient = redis.NewClient(&redis.Options{
	//	Addr:     viper.GetString("redis.address"),
	//	Password: viper.GetString("redis.password"),
	//	DB:       viper.GetInt("redis.db.relation"), // use default DB
	//})

	redisRelationClient = redis.NewClient(readRedisConfig("relation"))
	redisFollowClient = redis.NewClient(readRedisConfig("follow"))
	redisFollowerClient = redis.NewClient(readRedisConfig("follower"))
}

func readRedisConfig(configName string) *redis.Options {
	return &redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db." + configName),
	}
}
