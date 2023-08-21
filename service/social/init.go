package social

import (
	"github.com/cold-runner/simpleTikTok/pkg/dao/query"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/skylark/pkg/db"
)

func Init() {
	log.Init(&log.Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             "debug",
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	})
	log.Infow("日志初始化成功")

	mySQL, err := db.NewMySQL(&db.Options{
		Host:                  "127.0.0.1",
		Username:              "root",
		Password:              "my-secret",
		Database:              "simple_tiktok",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: 100,
		LogLevel:              2,
		Logger:                nil,
	})
	if err != nil {
		panic(err)
	}
	log.Infow("数据库初始化成功")
	query.SetDefault(mySQL)

	// 实例化dal层接口
	dal.MysqlIns = &dal.Mysql{Query: query.Q}
}
