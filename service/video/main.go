package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	VideoService "github.com/cold-runner/simpleTikTok/kitex_gen/VideoService/videoservice"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"net"
)

func main() {
	Init()
	// etcd 注册
	endpoints := []string{viper.GetString("etcd.address")}
	r, err := etcd.NewEtcdRegistry(endpoints)
	if err != nil {
		log.Panicw("Failed to create etcd registry", "err", err)
	}

	serviceAddr, err := net.ResolveTCPAddr("tcp",
		viper.GetString("etcd.video-service-address"))

	if err != nil {
		log.Panicw("Failed to resolve service address", "err", err)
	}
	// 服务器设置
	svr := VideoService.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(serviceAddr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 100, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.
			EndpointBasicInfo{ServiceName: viper.GetString("etcd." +
			"video-service-name")}),
	)
	// 启动服务器
	err = svr.Run()
	if err != nil {
		log.Fatalw("Failed to start server", "err", err)
	}
}
