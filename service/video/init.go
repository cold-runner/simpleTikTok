package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/service/video/dal"
	"github.com/cold-runner/simpleTikTok/service/video/minio"
	"github.com/cold-runner/simpleTikTok/service/video/mq"
	"github.com/cold-runner/simpleTikTok/service/video/rpc"
)

func Init() {
	config.InitViperConfig()
	dal.InitDB()
	minio.InitializeMinIO()
	mq.InitMQ()
	rpc.InitRPC()
}
