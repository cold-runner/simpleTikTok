package cmd

import (
	"github.com/cold-runner/simpleTikTok/apiServer/rpc"
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	mw "github.com/cold-runner/simpleTikTok/pkg/middleware"
)

func Init() {
	config.InitViperConfig()
	log.Init("log-default")
	rpc.InitRpc()
	mw.InitJwt()
}
