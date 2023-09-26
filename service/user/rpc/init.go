package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cold-runner/simpleTikTok/kitex_gen/RelationService/relationservice"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)

var RelationClient relationservice.Client

func InitRelationRPCinUser() {
	etcdAddress := []string{viper.GetString("etcd.address")}
	r, err := etcd.NewEtcdResolver(etcdAddress)
	if err != nil {
		log.Panicw("initUserRPC Error", "err", err)
	}
	log.Infow("Initializing Relation-RPC-Server from etcd server:",
		"etcdAddress",
		etcdAddress)

	c, err := relationservice.NewClient(
		viper.GetString("etcd.relation-service-name"), // 设置用户服务的名称
		client.WithResolver(r),                        // 使用之前创建的ETCD解析器
		client.WithMuxConnection(1),                   // 设置客户端的多路复用连接数为1
		//client.WithMiddleware(mw.CommonMiddleware), // 设置客户端的通用中间件
		//client.WithInstanceMW(mw.ClientMiddleware), // 设置客户端的实例中间件
		//client.WithSuite(tracing.NewClientSuite()), // 使用新的客户端追踪套件
		// 设置客户端的基本信息
		client.WithClientBasicInfo(&rpcinfo.
			EndpointBasicInfo{ServiceName: viper.GetString("etcd.relation-service-name")}),
	)
	RelationClient = c
}
