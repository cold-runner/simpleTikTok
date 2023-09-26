package main

import (
	"context"
	"fmt"
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	rd "github.com/cold-runner/simpleTikTok/service/user/dal/redis"
	"github.com/cold-runner/simpleTikTok/service/user/service"
	"github.com/redis/go-redis/v9"
)

func testConnection() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 配置中的地址
		Password: "",               // 配置中的密码
		DB:       0,                // 配置中的数据库号
	})

	// 创建一个上下文
	ctx := context.Background()

	// 尝试ping命令以测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong) // 输出“PONG”表示连接成功
	return nil
}

func main() {
	config.InitViperConfig()
	err := testConnection()
	dal.InitDB()
	rd.InitRedis()
	log.Init("log-default")
	if err != nil {
		fmt.Println("连接失败:", err)
	} else {
		fmt.Println("连接成功")
	}
	s := service.NewGetUserInfoService(context.Background())
	u, err := s.GetUserInfoById(5)
	if err != nil {
		fmt.Println("获取失败:", err)
	} else {
		fmt.Println("获取成功")
		fmt.Printf("用户详情: %+v", u)
	}
}
