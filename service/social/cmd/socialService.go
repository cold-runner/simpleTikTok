package main

import (
	"context"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService/socialservice"
	"github.com/cold-runner/simpleTikTok/service/social"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/service/v1"
	"log"
)

func main() {
	// 初始化数据库
	social.Init()
	// 初始化服务实例 TODO 注入rpc实例
	socialService := v1.NewCommentActionService(dal.MysqlIns, nil)
	// 注册服务实例到路由中
	serviceImpl := social.NewServiceImpl(context.Background(), socialService)
	// 注册路由
	svr := SocialService.NewServer(serviceImpl)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
