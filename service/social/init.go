package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/service/dtm/request"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
)

func InitSocialServer() {
	config.InitViperConfig()
	rpc.InitRPC()
	dal.InitDB()
	request.InitRequestVariables()
}
