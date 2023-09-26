package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService/socialservice"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"net"
)

func main() {
	InitSocialServer()
	// etcd 注册
	endpoints := []string{viper.GetString("etcd.address")}
	r, err := etcd.NewEtcdRegistry(endpoints)
	if err != nil {
		log.Panicw("Failed to create etcd registry", "err", err)
	}

	serviceAddr, err := net.ResolveTCPAddr("tcp",
		viper.GetString("etcd.social-service-address"))

	if err != nil {
		log.Panicw("Failed to resolve service address", "err", err)
	}
	// 服务器设置
	svr := SocialService.NewServer(new(SocialServiceImpl),
		server.WithServiceAddr(serviceAddr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 100, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.
			EndpointBasicInfo{ServiceName: viper.GetString("etcd." +
			"social-service-name")}),
	)
	// 启动服务器
	err = svr.Run()
	if err != nil {
		log.Fatalw("Failed to start server", "err", err)
	}
}
