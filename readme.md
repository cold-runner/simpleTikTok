simpleTikTok
├── Dockerfile
├── apiServer
│   ├── biz
│   │   ├── handler
│   │   │   ├── ApiServer
│   │   │   │   ├── relation_service.go
│   │   │   │   ├── social_service.go
│   │   │   │   ├── user_service.go
│   │   │   │   └── video_service.go
│   │   │   ├── ping.go
│   │   │   └── response
│   │   │       ├── relation.go
│   │   │       ├── social.go
│   │   │       ├── user.go
│   │   │       └── video.go
│   │   ├── model
│   │   │   ├── ApiServer
│   │   │   │   └── apiServer.pb.go
│   │   │   └── api
│   │   │       └── api.pb.go
│   │   └── router
│   │       ├── ApiServer
│   │       │   ├── apiServer.go
│   │       │   └── middleware.go
│   │       └── register.go
│   ├── build.sh
│   ├── cmd
│   │   └── init.go
│   ├── main.go
│   ├── router.go
│   ├── router_gen.go
│   ├── rpc
│   │   ├── init.go
│   │   ├── relation.go
│   │   ├── social.go
│   │   ├── user.go
│   │   └── video.go
│   └── script
│       └── bootstrap.sh
├── config
│   ├── config.yaml
│   └── docker 启动.md
├── data
│   ├── images
│   └── videos
├── docker-compose.yml
├── docs
│   ├── error_code.md
│   └── img.png
├── go.mod
├── go.sum
├── idl
│   ├── BaseResp.proto
│   ├── RelationServer.proto
│   ├── SocialServer.proto
│   ├── UserServer.proto
│   ├── VideoServer.proto
│   ├── api.proto
│   ├── apiServer.proto
│   └── 代码生成.md
├── init.sql
├── kitex_gen
│   ├── BaseResponse
│   ├── RelationService
│   ├── SocialService
│   ├── UserService
│   └── VideoService
├── package-lock.json
├── package.json
├── pkg
│   ├── config
│   │   └── init.go
│   ├── db
│   │   ├── init.go
│   │   └── 数据库实例使用说明.md
│   ├── encryption
│   │   └── encrypter.go
│   ├── errno
│   │   ├── code.go
│   │   └── errno.go
│   ├── log
│   │   ├── log.go
│   │   ├── options.go
│   │   └── 日志使用说明.md
│   └── middleware
│       └── jwt.go
├── script
│   └── bootstrap.sh
├── service
│   ├── dtm
│   │   ├── init.go
│   │   ├── main.go
│   │   └── request
│   │       ├── commentRequest.go
│   │       ├── favoriteRequest.go
│   │       └── init.go
│   ├── message
│   │   ├── api
│   │   └── cmd
│   ├── relation
│   │   ├── build.sh
│   │   ├── dal
│   │   │   ├── init.go
│   │   │   ├── model
│   │   │   │   └── relation.go
│   │   │   ├── redis
│   │   │   │   ├── init.go
│   │   │   │   └── relation.go
│   │   │   └── relation.go
│   │   ├── handler.go
│   │   ├── init.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   ├── mq
│   │   │   ├── consumer.go
│   │   │   ├── init.go
│   │   │   └── producer.go
│   │   ├── response
│   │   │   └── resp.go
│   │   ├── rpc
│   │   │   ├── init.go
│   │   │   └── user.go
│   │   └── service
│   │       ├── getFollowList.go
│   │       ├── getFollowerList.go
│   │       ├── queryRelation.go
│   │       ├── queryUserInfosWithRelation.go
│   │       └── relationAction.go
│   ├── social
│   │   ├── build.sh
│   │   ├── dal
│   │   │   ├── comment.go
│   │   │   ├── favorite.go
│   │   │   ├── init.go
│   │   │   └── model
│   │   │       ├── comment.go
│   │   │       └── favorite.go
│   │   ├── handler.go
│   │   ├── init.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   ├── response
│   │   │   └── resp.go
│   │   ├── rpc
│   │   │   └── init.go
│   │   └── service
│   │       ├── commentAction.go
│   │       ├── commentList.go
│   │       ├── favoriteAction.go
│   │       └── favoriteList.go
│   ├── user
│   │   ├── build.sh
│   │   ├── dal
│   │   │   ├── init.go
│   │   │   ├── model
│   │   │   │   ├── dynamicUserFields.go
│   │   │   │   ├── stableUserFields.go
│   │   │   │   ├── user.go
│   │   │   │   └── userInfo.go
│   │   │   ├── redis
│   │   │   │   ├── init.go
│   │   │   │   └── user.go
│   │   │   └── user.go
│   │   ├── handler.go
│   │   ├── init.go
│   │   ├── kitex_info.yaml
│   │   ├── main.go
│   │   ├── mq
│   │   │   ├── consumer.go
│   │   │   ├── init.go
│   │   │   └── producer.go
│   │   ├── response
│   │   │   └── resp.go
│   │   ├── rpc
│   │   │   └── init.go
│   │   ├── service
│   │   │   ├── changeUserFollowCount.go
│   │   │   ├── getUserInfo.go
│   │   │   ├── login.go
│   │   │   ├── mGetUserInfo.go
│   │   │   ├── register.go
│   │   │   └── updateUserInfo.go
│   │   ├── test
│   │   │   └── testRedis.go
│   │   └── user
│   └── video
│       ├── dal
│       │   ├── init.go
│       │   ├── model
│       │   │   └── video.go
│       │   └── video.go
│       ├── handler.go
│       ├── init.go
│       ├── kitex_info.yaml
│       ├── main.go
│       ├── minio
│       │   ├── docker 启动说明.md
│       │   ├── init.go
│       │   └── service.go
│       ├── mq
│       │   ├── consumer.go
│       │   ├── init.go
│       │   ├── producer.go
│       │   ├── updateInfoConsumer.go
│       │   └── updateInfoProducer.go
│       ├── response
│       │   └── resp.go
│       ├── rpc
│       │   └── init.go
│       └── service
│           ├── getVideoInfo.go
│           ├── getVideoListByIds.go
│           ├── getVideoPublishList.go
│           ├── publishAction.go
│           ├── updateVideoInfo.go
│           └── videoFeed.go
├── test
│   └── main.go
└── tmp
└── log
├── common
│   └── log.log
├── dtm
│   └── dtm.log
├── relation
│   └── relation.log
├── social
│   └── social.log
└── user
└── user.log
