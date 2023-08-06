package apiServer

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "github.com/cold-runner/simpleTikTok/apiServer/biz/router"
)

// Register registers all routers.
func Register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
