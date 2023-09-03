package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/service/user/dal"
	"github.com/cold-runner/simpleTikTok/service/user/dal/redis"
	"github.com/cold-runner/simpleTikTok/service/user/mq"
	"github.com/cold-runner/simpleTikTok/service/user/rpc"
)

func Init() {
	config.InitViperConfig()
	dal.InitDB()
	mq.InitMQ()
	mq.InitSocialMQ()
	rpc.InitRelationRPCinUser()
	redis.InitRedis()
}
