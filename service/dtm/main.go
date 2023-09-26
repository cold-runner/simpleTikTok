package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/dtm/request"
	socialDal "github.com/cold-runner/simpleTikTok/service/social/dal"
	socialRPC "github.com/cold-runner/simpleTikTok/service/social/rpc"
)

func main() {
	config.InitViperConfig()
	request.InitRequestVariables()
	log.Init("log-dtm")
	socialDal.InitDB()
	socialRPC.InitRPC()
	InitDTM()
}
