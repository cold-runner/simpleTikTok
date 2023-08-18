package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/db"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
)

func main() {
	config.InitViperConfig()
	log.Init("log-user")
	db, err := db.NewDB("db-user")
	if err != nil {
		log.Fatalw("db failed", "err", err)
	}

	//if err := db.Exec("SELECT 1").Error; err != nil {
	//	log.Fatalw("Database query failed", "err", err)
	//}
	//log.Infow("Selecting 1")
	db.AutoMigrate(&model.User{})
}
