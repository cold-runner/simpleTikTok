package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/service/relation/dal"
	"github.com/cold-runner/simpleTikTok/service/relation/dal/redis"
	"github.com/cold-runner/simpleTikTok/service/relation/mq"
	"github.com/cold-runner/simpleTikTok/service/relation/rpc"
)

func InitRelationServer() {
	config.InitViperConfig()
	rpc.InitUserRPCinRelation()
	mq.InitMQ()
	redis.InitRedis()
	dal.InitDB()
}
