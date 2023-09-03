package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/dtm/favoriteRequest"
	socialDal "github.com/cold-runner/simpleTikTok/service/social/dal"
	socialRPC "github.com/cold-runner/simpleTikTok/service/social/rpc"
)

func main() {
	config.InitViperConfig()
	favoriteRequest.InitRequestVariables()
	log.Init("log-dtm")
	socialDal.InitDB()
	socialRPC.InitRPC()
	InitDTM()
}
